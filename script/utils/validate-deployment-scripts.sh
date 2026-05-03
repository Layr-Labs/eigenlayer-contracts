#!/usr/bin/env bash
#
# validate-deployment-scripts.sh
#
# Local equivalent of .github/workflows/validate-deployment-scripts.yml
# Runs `zeus test` for every release script across all environments.
#
# Usage:
#   ./script/utils/validate-deployment-scripts.sh                  # run all envs
#   ./script/utils/validate-deployment-scripts.sh mainnet base     # run specific envs
#
# RPC Configuration (pick one):
#   1. Export environment variables before running:
#        export RPC_MAINNET="https://..."
#        export RPC_SEPOLIA="https://..."
#        ...
#   2. Create a .env.rpc file in the repo root (gitignored):
#        RPC_MAINNET=https://...
#        RPC_SEPOLIA=https://...
#   3. Do nothing — the script will use free public RPCs as defaults.
#
# Environment variables:
#   SKIP_BUILD=1        Skip forge build step
#   SKIP_FMT=1          Skip forge fmt check
#   VERBOSE=1           Show full zeus output (default: summary only)
#   FAIL_FAST=1         Stop on first failure (default: run all, report at end)
#
set -euo pipefail

# ── Repo root ──────────────────────────────────────────────────────────────────
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
cd "$REPO_ROOT"

# ── Colors ─────────────────────────────────────────────────────────────────────
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
BOLD='\033[1m'
NC='\033[0m' # No Color

# ── Helpers ────────────────────────────────────────────────────────────────────
info()  { echo -e "${BLUE}ℹ${NC}  $*"; }
ok()    { echo -e "${GREEN}✔${NC}  $*"; }
warn()  { echo -e "${YELLOW}⚠${NC}  $*"; }
fail()  { echo -e "${RED}✖${NC}  $*"; }
header(){ echo -e "\n${BOLD}═══ $* ═══${NC}"; }

# ── Load .env.rpc if present ───────────────────────────────────────────────────
if [[ -f "$REPO_ROOT/.env.rpc" ]]; then
    info "Loading RPC URLs from .env.rpc"
    set -a
    # shellcheck disable=SC1091
    source "$REPO_ROOT/.env.rpc"
    set +a
fi

# ── Default public RPCs ───────────────────────────────────────────────────────
: "${RPC_MAINNET:=https://ethereum-rpc.publicnode.com}"
: "${RPC_SEPOLIA:=https://ethereum-sepolia-rpc.publicnode.com}"
: "${RPC_HOODI:=https://rpc.hoodi.ethpandaops.io}"
: "${RPC_BASE_SEPOLIA:=https://base-sepolia-rpc.publicnode.com}"
: "${RPC_BASE:=https://base-rpc.publicnode.com}"

# ── Environment matrix (same as CI) ───────────────────────────────────────────
ALL_ENVS=(mainnet testnet-sepolia testnet-hoodi testnet-base-sepolia base preprod-hoodi)

# Map env name → RPC URL
declare -A RPC_MAP=(
    [mainnet]="$RPC_MAINNET"
    [testnet-sepolia]="$RPC_SEPOLIA"
    [testnet-hoodi]="$RPC_HOODI"
    [testnet-base-sepolia]="$RPC_BASE_SEPOLIA"
    [base]="$RPC_BASE"
    [preprod-hoodi]="$RPC_HOODI"
)

# ── Determine which envs to run ───────────────────────────────────────────────
if [[ $# -gt 0 ]]; then
    ENVS=("$@")
    # Validate user-supplied envs
    for env in "${ENVS[@]}"; do
        if [[ -z "${RPC_MAP[$env]+x}" ]]; then
            fail "Unknown environment: '${env}'"
            echo "    Valid environments: ${ALL_ENVS[*]}"
            exit 1
        fi
    done
else
    ENVS=("${ALL_ENVS[@]}")
fi

# ── Pre-flight: check tools ───────────────────────────────────────────────────
header "Pre-flight checks"

if ! command -v zeus &>/dev/null; then
    fail "zeus not found. Install with: npm install -g @layr-labs/zeus --ignore-scripts"
    exit 1
fi
ok "zeus  $(zeus --version 2>/dev/null || echo '(version unknown)')"

if ! command -v forge &>/dev/null; then
    fail "forge not found. Install Foundry: https://book.getfoundry.sh/getting-started/installation"
    exit 1
fi
ok "forge $(forge --version 2>/dev/null | head -1)"

# ── RPC connectivity check ────────────────────────────────────────────────────
header "RPC connectivity"
for env in "${ENVS[@]}"; do
    rpc="${RPC_MAP[$env]}"
    if [[ -z "$rpc" ]]; then
        fail "$env — no RPC URL configured"
        echo "    Set RPC_MAINNET / RPC_SEPOLIA / RPC_HOODI / RPC_BASE_SEPOLIA / RPC_BASE"
        echo "    via environment variable or .env.rpc file."
        exit 1
    fi
    # Quick JSON-RPC ping (eth_chainId)
    if curl -sf -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}' \
        --max-time 10 "$rpc" >/dev/null 2>&1; then
        ok "$env → $rpc"
    else
        warn "$env → $rpc  (unreachable — tests may fail)"
    fi
done

# ── Forge build ────────────────────────────────────────────────────────────────
if [[ "${SKIP_BUILD:-}" != "1" ]]; then
    header "Forge build"
    forge build --sizes
    ok "Build succeeded"
else
    info "Skipping forge build (SKIP_BUILD=1)"
fi

# ── Forge fmt ──────────────────────────────────────────────────────────────────
if [[ "${SKIP_FMT:-}" != "1" ]]; then
    header "Forge fmt check"
    if forge fmt --check && FOUNDRY_PROFILE=test forge fmt --check; then
        ok "Formatting OK"
    else
        warn "Formatting issues detected (non-blocking)"
    fi
else
    info "Skipping forge fmt (SKIP_FMT=1)"
fi

# ── Discover release scripts (same find as CI) ────────────────────────────────
header "Discovering release scripts"
RELEASE_FILES=$(find script/releases -type f -name "*.sol" \
    ! -name "Env.sol" \
    ! -name "CrosschainDeployLib.sol" \
    ! -name "TestUtils.sol" \
    ! -name "CoreContractsDeployer.sol" \
    ! -name "CoreUpgradeQueueBuilder.sol" \
    2>/dev/null | sort || echo "")

if [[ -z "$RELEASE_FILES" ]]; then
    warn "No release scripts found under script/releases/"
    exit 0
fi

FILE_COUNT=$(echo "$RELEASE_FILES" | wc -l | tr -d ' ')
info "Found ${FILE_COUNT} release scripts:"
echo "$RELEASE_FILES" | while read -r f; do echo "    $f"; done

# ── Run tests ──────────────────────────────────────────────────────────────────
TOTAL=0
PASSED=0
FAILED=0
SKIPPED=0
FAILURES=()

for env in "${ENVS[@]}"; do
    rpc="${RPC_MAP[$env]}"
    header "Environment: ${env}"
    info "RPC: ${rpc}"

    for file in $RELEASE_FILES; do
        TOTAL=$((TOTAL + 1))
        short_file="${file#script/releases/}"
        echo -ne "  ${BLUE}▶${NC} [${env}] ${short_file} ... "

        LOGFILE=$(mktemp)
        EXIT_CODE=0
        zeus test --env "$env" --rpcUrl "$rpc" "$file" >"$LOGFILE" 2>&1 || EXIT_CODE=$?

        if [[ $EXIT_CODE -eq 0 ]]; then
            # Check if zeus reported test failures despite exit 0
            if grep -q "test failed" "$LOGFILE" 2>/dev/null; then
                FAILED=$((FAILED + 1))
                FAILURES+=("[${env}] ${short_file}")
                echo -e "${RED}FAIL${NC}"
                if [[ "${VERBOSE:-}" == "1" ]]; then
                    cat "$LOGFILE"
                else
                    grep -E "(FAIL|fail|Error|error|revert)" "$LOGFILE" | head -5
                fi
            elif grep -q "test succeeded" "$LOGFILE" 2>/dev/null; then
                PASSED=$((PASSED + 1))
                echo -e "${GREEN}PASS${NC}"
            else
                # No clear pass/fail — might be a no-op (e.g. base env skipping core scripts)
                SKIPPED=$((SKIPPED + 1))
                echo -e "${YELLOW}SKIP${NC} (no test output)"
            fi
        else
            FAILED=$((FAILED + 1))
            FAILURES+=("[${env}] ${short_file}")
            echo -e "${RED}FAIL${NC} (exit ${EXIT_CODE})"
            if [[ "${VERBOSE:-}" == "1" ]]; then
                cat "$LOGFILE"
            else
                tail -20 "$LOGFILE"
            fi
        fi

        rm -f "$LOGFILE"

        if [[ "${FAIL_FAST:-}" == "1" && $FAILED -gt 0 ]]; then
            fail "Stopping early (FAIL_FAST=1)"
            break 2
        fi
    done
done

# ── Summary ────────────────────────────────────────────────────────────────────
header "Summary"
echo -e "  ${BOLD}Total:${NC}   ${TOTAL}"
echo -e "  ${GREEN}Passed:${NC}  ${PASSED}"
echo -e "  ${YELLOW}Skipped:${NC} ${SKIPPED}"
echo -e "  ${RED}Failed:${NC}  ${FAILED}"

if [[ ${#FAILURES[@]} -gt 0 ]]; then
    echo ""
    fail "Failed tests:"
    for f in "${FAILURES[@]}"; do
        echo -e "    ${RED}•${NC} $f"
    done
    echo ""
    echo -e "  Re-run a single failure with:"
    echo -e "    ${BOLD}zeus test --env <env> --rpcUrl <rpc> <file>${NC}"
    echo ""
    echo -e "  Or re-run with verbose output:"
    echo -e "    ${BOLD}VERBOSE=1 $0 <env>${NC}"
    exit 1
fi

echo ""
ok "All tests passed!"
