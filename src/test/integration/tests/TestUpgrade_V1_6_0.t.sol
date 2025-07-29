// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

import "src/test/integration/IntegrationChecks.t.sol";
import "src/contracts/token/Eigen.sol";

interface IOldPod {
    function GENESIS_TIME() external view returns (uint64);
}

interface IOldALM {
    function owner() external view returns (address);
}

/**
 * Queue txns:
 * - DISTRO: https://etherscan.io/tx/0x8e6f1701abc942d468a5cea427ae16069b5ee6341407e2f3ac64e18e01e06756
 * - MOOCOW: https://etherscan.io/tx/0x6e219b1ad78aa0e988ba1cde6e0bfc415edc3d1d8733dc495c8c7e79f77414f8
 */
contract UpgradeTest is IntegrationCheckUtils {

    using Strings for *;
    using StdStyle for *;

    // Ops Multisig
    address constant opsMS = 0xBE1685C81aA44FF9FB319dD389addd9374383e90;
    // Protocol Council Multisig
    address constant pCMS = 0x461854d84Ee845F905e0eCf6C288DDEEb4A9533F;
    // Executor Multisig
    address constant execMS = 0x369e6F597e22EaB55fFb173C6d9cD234BD699111;

    TimelockController constant tcl = TimelockController(payable(0xC06Fd4F821eaC1fF1ae8067b36342899b57BAa2d));
    ProxyAdmin constant proxyAdmin = ProxyAdmin(0x8b9566AdA63B64d1E1dcF1418b43fd1433b72444);

    address constant beigen_addr = 0x83E9115d334D248Ce39a6f36144aEaB5b3456e75;
    address constant eigen_addr = 0xec53bF9167f50cDEB3Ae105f56099aaaB9061F83;
    Eigen constant eigen = Eigen(eigen_addr);

    bytes calldataToExecutor_DISTRO = hex"6A76120200000000000000000000000040A2ACCBD92BCA938B02010E17A5B8929B49130D000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000C800000000000000000000000000000000000000000000000000000000000000B048D80FF0A00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000AA2008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC400000000000000000000000039053D51B77DC0D36036FC1FCC8CB819DF8EF37A0000000000000000000000006EED6C2802DF347E05884857CDDB2D3E96D12F89008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC4000000000000000000000000948A420B8CC1D6BFD0B6087C2E7C344A2CD0BC39000000000000000000000000C97602648FA52F92B4EE2B0E5A54BD15B6CB0345008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC4000000000000000000000000858646372CC42E1A627FCE94AA7A7033E7CF075A00000000000000000000000046AEFD30415BE99E20169EE7046F65784B46D123008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC400000000000000000000000091E677B07F7AF907EC9A428AAFA9FC14A0D3A338000000000000000000000000E48D7CAEC1790B293667E4BB2DE1E00536F2BABD008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC4000000000000000000000000ACB55C530ACDB2849E6D4F36992CD8C9D50ED8F7000000000000000000000000530FDB7ADF7D489DF49C27E3D3512C0DD64886BE000ED6703C298D28AE0878D1B28E88CA87F9662FE9000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000243659CFE6000000000000000000000000D4D1746142642DB4C1AB17B03B9C58BAAC913E5B008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC400000000000000000000000093C4B944D05DFE6DF7645A86CD2206016C51564D00000000000000000000000062F7226FB9D615590EADB539713B250FB2FDF4E0008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC40000000000000000000000001BEE69B7DFFFA4E2D53C2A2DF135C388AD25DCD200000000000000000000000062F7226FB9D615590EADB539713B250FB2FDF4E0008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC400000000000000000000000054945180DB7943C0ED0FEE7EDAB2BD24620256BC00000000000000000000000062F7226FB9D615590EADB539713B250FB2FDF4E0008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC40000000000000000000000009D7ED45EE2E8FC5482FA2428F15C971E6369011D00000000000000000000000062F7226FB9D615590EADB539713B250FB2FDF4E0008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC400000000000000000000000013760F50A9D7377E4F20CB8CF9E4C26586C658FF00000000000000000000000062F7226FB9D615590EADB539713B250FB2FDF4E0008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC4000000000000000000000000A4C637E0F704745D182E4D38CAB7E7485321D05900000000000000000000000062F7226FB9D615590EADB539713B250FB2FDF4E0008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC400000000000000000000000057BA429517C3473B6D34CA9ACD56C0E735B94C0200000000000000000000000062F7226FB9D615590EADB539713B250FB2FDF4E0008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC40000000000000000000000000FE4F44BEE93503346A3AC9EE5A26B130A5796D600000000000000000000000062F7226FB9D615590EADB539713B250FB2FDF4E0008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC40000000000000000000000007CA911E83DABF90C90DD3DE5411A10F1A611218400000000000000000000000062F7226FB9D615590EADB539713B250FB2FDF4E0008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC40000000000000000000000008CA7A5D6F3ACD3A7A8BC468A8CD0FB14B6BD28B600000000000000000000000062F7226FB9D615590EADB539713B250FB2FDF4E0008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC4000000000000000000000000AE60D8180437B5C34BB956822AC271097258447300000000000000000000000062F7226FB9D615590EADB539713B250FB2FDF4E0008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC4000000000000000000000000298AFB19A105D59E74658C4C334FF360BADE6DD200000000000000000000000062F7226FB9D615590EADB539713B250FB2FDF4E0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041000000000000000000000000C06FD4F821EAC1FF1AE8067B36342899B57BAA2D00000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000";
    bytes calldataToExecutor_MOOCOW = hex"6A76120200000000000000000000000040A2ACCBD92BCA938B02010E17A5B8929B49130D0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002E000000000000000000000000000000000000000000000000000000000000001648D80FF0A00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000112005A2A4F2F3C18F09179B6703E63D9EDD165909073000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000243659CFE6000000000000000000000000CB27A4819A64FBA93ABD4D480E4466AEC0503745008B9566ADA63B64D1E1DCF1418B43FD1433B724440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004499A88EC4000000000000000000000000EC53BF9167F50CDEB3AE105F56099AAAB9061F830000000000000000000000002C4A81E257381F87F5A5C4BD525116466D972E500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041000000000000000000000000C06FD4F821EAC1FF1AE8067B36342899B57BAA2D00000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000";

    // Timestamp of block where transaction was queued
    // (Fri, Jul 11 @ ~12.30 EST)
    uint queueTimestamp = 1752251303;
    // 10 Days
    uint queueDelay = 864000;

    // Used to make calls to the EIP-4788 oracle work when testing EigenPod methods
    uint startTimestamp;
    uint executeTimestamp = queueTimestamp + queueDelay;

    // Sample mainnet pod with one active validator and the given validator pubkey
    address constant pod_addr = 0xA6f93249580EC3F08016cD3d4154AADD70aC3C96;
    EigenPod constant pod = EigenPod(payable(pod_addr));
    bytes constant validatorPubkey = hex"9067f3fb35c5fc6aa9e05f102519c22050c447acaf6b88aeb16d92fc14e30891f4c2bc21ee95ef5779224f5ceef3a168";

    struct Proxies {
        AllocationManager alm;
        DelegationManager dm;
        StrategyManager sm;
        EigenPodManager em;
        EigenStrategy eigenStrategy;
        StrategyBase factoryStrat;
        StrategyBaseTVLLimits instanceStrat;
    }

    Proxies p;

    // Set up "current impl address list"
    function _setup() internal {
        p = Proxies({
            alm: AllocationManager(0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39),
            dm: DelegationManager(0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A),
            sm: StrategyManager(0x858646372CC42E1A627fcE94aa7A7033e7CF075A),
            em: EigenPodManager(0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338),
            eigenStrategy: EigenStrategy(0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7),
            factoryStrat: StrategyBase(0x7079A4277eAF578cbe9682ac7BC3EfFF8635ebBf),          // EtherFi strategy
            instanceStrat: StrategyBaseTVLLimits(0x93c4b944D05dfe6df7645A86cd2206016c51564D) // stETH strategy
        });
    }

    function test_TEMP() public {
        _setup();

        cheats.createSelectFork(cheats.rpcUrl("mainnet"));
        // cheats.createSelectFork(cheats.rpcUrl("mainnet"), block.number - 1);

        console.log("");
        _logBanner("Using current block/time, reading mainnet state...".yellow().bold());
        startTimestamp = block.timestamp;
        console.log("timestamp info:".dim());
        console.log(" - current timestamp: %d", startTimestamp);
        console.log(" - execute timestamp: %d", executeTimestamp);
        console.log("");

        // _checkRoles();
        bool upgradeDone = _checkIsPending();

        // Warp forward if we haven't hit execution time yet
        if (block.timestamp < executeTimestamp) {
            require(!upgradeDone, "Expected upgrade to be incomplete, since it can't be completed yet.".red().bold());

            _logBanner("Warping to execute timestamp".yellow().bold());
            cheats.warp(executeTimestamp);

            _checkIsPending();
        }
        
        if (!upgradeDone) {
            _logBanner("Execution time has been reached; upgrade not active. Checking new functions.".yellow().bold());

            // Try calling methods before upgrade, then run upgrade.
            _tryCallFunctions();
            _executeUpgrade();

            // Timelock should say the upgrade is completed
            upgradeDone = _checkIsPending();
            require(upgradeDone, "Expected TimelockController to mark operations as complete post-upgrade.".red().bold());
        }

        // Try calling new methods now that the upgrade is complete
        _logBanner("Upgrade is complete. Calling new functions.".green().bold());
        _tryCallFunctions();
    }

    function _checkRoles() internal {
        console.log("_checkRoles()...".dim());

        bytes32 PROPOSER_ROLE = tcl.PROPOSER_ROLE();
        bytes32 CANCELLER_ROLE = tcl.CANCELLER_ROLE();
        bytes32 EXECUTOR_ROLE = tcl.EXECUTOR_ROLE();
        
        // opsMS has PROPOSER/CANCELLER
        require(tcl.hasRole(PROPOSER_ROLE, opsMS), "opsMS should have proposer role");
        require(tcl.hasRole(CANCELLER_ROLE, opsMS), "opsMS should have canceller role");
        require(!tcl.hasRole(EXECUTOR_ROLE, opsMS), "opsMS should NOT have executor role");

        // protocolCouncilMS has CANCELLER/EXECUTOR
        require(tcl.hasRole(PROPOSER_ROLE, pCMS), "pCMS should have proposer role");
        require(!tcl.hasRole(CANCELLER_ROLE, pCMS), "pCMS should NOT have canceller role");
        require(tcl.hasRole(EXECUTOR_ROLE, pCMS), "pCMS should have executor role");

        console.log("_checkRoles success\n".green());
    }

    function _checkIsPending() internal returns (bool upgradeDone) {
        console.log("checking timelock controller state...".dim());

        bytes32 DISTRO_OPERATION_ID = tcl.hashOperation({
            target: execMS,
            value: 0,
            data: calldataToExecutor_DISTRO,
            predecessor: bytes32(0),
            salt: bytes32(0)
        });

        bytes32 MOOCOW_OPERATION_ID = tcl.hashOperation({
            target: execMS,
            value: 0,
            data: calldataToExecutor_MOOCOW,
            predecessor: bytes32(0),
            salt: bytes32(0)
        });

        __logTruthy(" - DISTRO is pending", tcl.isOperationPending(DISTRO_OPERATION_ID));
        __logTruthy(" - MOOCOW is pending", tcl.isOperationPending(MOOCOW_OPERATION_ID));

        __logTruthy(" - DISTRO is ready", tcl.isOperationReady(DISTRO_OPERATION_ID));
        __logTruthy(" - MOOCOW is ready", tcl.isOperationReady(MOOCOW_OPERATION_ID));

        bool distroDone = tcl.isOperationDone(DISTRO_OPERATION_ID);
        bool moocowDone = tcl.isOperationDone(MOOCOW_OPERATION_ID);

        __logTruthy(" - DISTRO is done", distroDone);
        __logTruthy(" - MOOCOW is done", moocowDone);
        require(distroDone == moocowDone, "expected both upgrades to be in same state");

        console.log("");
        __logTruthy("current timestamp >= execution time", block.timestamp >= executeTimestamp);
        __logTruthy("upgrade complete", distroDone && moocowDone);
        console.log("");

        return (distroDone && moocowDone);
    }

    function _tryCallFunctions() internal {
        // Since we want to be able to call this method more than once, we snapshot state here
        // and revert state changes at the end of the method.
        uint id = cheats.snapshotState();

        _tryCallVersions();

        {
            // Call new DISTRO allocationManager methods
            console.log("%s %s %s", "Checking".cyan(), "DISTRO (alm)".magenta(), "functions...".cyan());

            // Using Singularity AVS and its permissioned caller (pulled from Etherscan)
            address avsCaller = address(0x639ebD317364962Ce8F785D3d523118841C2b81F);
            address avs = address(0x128e65461Ee2D794fc7F1A8c732dd9A275A8618c);   
            cheats.startPrank(avsCaller);

            {
                // 1. alm.owner(): REMOVED after upgrade
                address almOwner;
                bool methodExists;
                try IOldALM(address(p.alm)).owner() returns (address ao) {
                    methodExists = true;
                    almOwner = ao;
                } catch (bytes memory) {
                    methodExists = false;
                }
                __logTruMoo(" - alm.owner() method removed", !methodExists);
                if (methodExists) console.log(" -- (current owner: %s)".dim(), almOwner);
            }

            CreateSetParams[] memory params = new CreateSetParams[](1);
            params[0].operatorSetId = 420;
            params[0].strategies = new IStrategy[](1);
            params[0].strategies[0] = p.instanceStrat; // stETH
            address[] memory recipients = new address[](1);
            recipients[0] = address(0x93C4b944D05dfE6DF7645A86Cd2206016C51564e); // some random empty address

            OperatorSet memory opSet = OperatorSet(avs, 420);

            {
                // 2. alm.createRedistributingOperatorSets(): ADDED after upgrade
                bool methodExists;
                try p.alm.createRedistributingOperatorSets(avs, params, recipients) {
                    methodExists = true;
                } catch (bytes memory err) {
                    methodExists = false;
                }

                __logTruMoo(" - alm.createRedistributingOpSets() method added", methodExists);
            }

            {
                // 3. alm redistribution getters: ADDED after upgrade
                bool methodExists;
                address recipient;
                try p.alm.getRedistributionRecipient(opSet) returns (address r) {
                    methodExists = true;
                    recipient = r;
                } catch (bytes memory err) {
                    methodExists = false;
                }

                __logTruMoo(" - alm.getRedistributionRecipient() method added", methodExists);
                if (methodExists) __logTruthy(" -- recipient is expected", recipient == recipients[0]);
            }

            {
                // 4. alm.isRedistributingOperatorSet(): ADDED after upgrade
                bool methodExists;
                bool isRedistributing;
                try p.alm.isRedistributingOperatorSet(opSet) returns (bool b) {
                    methodExists = true;
                    isRedistributing = b;
                } catch (bytes memory err) {
                    methodExists = false;
                }

                __logTruMoo(" - alm.isRedistributingOperatorSet() method added", methodExists);
                if (methodExists) __logTruthy(" -- status is expected", isRedistributing);
            }

            {
                // 5. alm.getSlashCount(): ADDED after upgrade
                bool methodExists;
                try p.alm.getSlashCount(opSet) returns (uint) {
                    methodExists = true;
                } catch (bytes memory err) {
                    methodExists = false;
                }

                __logTruMoo(" - alm.getSlashCount() method added", methodExists);
            }

            {
                // 5. alm.isOperatorRedistributable(): ADDED after upgrade
                bool methodExists;
                try p.alm.isOperatorRedistributable(pod_addr) returns (bool) {
                    methodExists = true;
                } catch (bytes memory err) {
                    methodExists = false;
                }

                __logTruMoo(" - alm.isOperatorRedistributable() method added", methodExists);
            }

            {
                // 7. alm.eigenStrategy(): ADDED after upgrade
                bool methodExists;
                IStrategy es;
                try p.alm.eigenStrategy() returns (IStrategy s) {
                    methodExists = true;
                    es = s;
                } catch (bytes memory err) {
                    methodExists = false;
                    
                }

                __logTruMoo(" - alm.eigenStrategy() method added", methodExists);
                if (methodExists) __logTruthy(" -- returned expected address", IStrategy(p.eigenStrategy) == es);
            }

            // alm.slashOperator returns (uint, uint[])

            cheats.stopPrank();
            console.log("");
        }

        _tryCallMOOCOW();
        _tryCallEigen();

        cheats.revertToState(id);
    }

    function _tryCallVersions() internal {
        console.log("%s %s %s", "Checking".cyan(), "version()".magenta(), "methods...".cyan());
        
        {
            // 1. alm.version should read 1.5.0 in DISTRO; 1.3.0 before
            string memory _version = p.alm.version();
            bool is1_5 = _strEq(_version, "1.5.0");
            __logTruMoo(" - alm.version is 1.5.0", is1_5);
            if (!is1_5) console.log(" -- (current version: %s)".dim(), _version);
        }

        {
            // 2. dm.version should read 1.5.0 in DISTRO; 1.3.0 before
            string memory _version = p.dm.version();
            bool is1_5 = _strEq(_version, "1.5.0");
            __logTruMoo(" - dm.version is 1.5.0", is1_5);
            if (!is1_5) console.log(" -- (current version: %s)".dim(), _version);
        }

        {
            // 3. sm.version should read 1.5.0 in DISTRO; 1.3.0 before
            string memory _version = p.sm.version();
            bool is1_5 = _strEq(_version, "1.5.0");
            __logTruMoo(" - sm.version is 1.5.0", is1_5);
            if (!is1_5) console.log(" -- (current version: %s)".dim(), _version);
        }

        {
            // 4. factoryStrat.version should read 1.5.0 in DISTRO; 1.3.0 before
            string memory _version = p.factoryStrat.version();
            bool is1_5 = _strEq(_version, "1.5.0");
            __logTruMoo(" - factoryStrat.version is 1.5.0", is1_5);
            if (!is1_5) console.log(" -- (current version: %s)".dim(), _version);
        }

        {
            // 5. instanceStrat.version should read 1.5.0 in DISTRO; 1.3.0 before
            string memory _version = p.instanceStrat.version();
            bool is1_5 = _strEq(_version, "1.5.0");
            __logTruMoo(" - instanceStrat.version is 1.5.0", is1_5);
            if (!is1_5) console.log(" -- (current version: %s)".dim(), _version);
        }

        console.log("");
    }

    function _tryCallMOOCOW() internal {
        {
            // Call new EigenPod MOOCOW methods
            console.log("%s %s %s", "Checking".cyan(), "MOOCOW (pod)".magenta(), "functions...".cyan());

            address podOwner = pod.podOwner();
            cheats.startPrank(podOwner);

            // Basic info
            console.log(" - Using pod:        %s".dim(), pod_addr);
            console.log(" - podOwner():       %s".dim(), podOwner);
            console.log(" - proofSubmitter(): %s".dim(), pod.proofSubmitter());
            console.log(" - activeValidatorCount(): %s".dim(), pod.activeValidatorCount());
            __logTruthy(" - checkpoint active", pod.currentCheckpointTimestamp() != 0);
            console.log("");

            {
                // 1. Version string should read 1.6.0 in MOOCOW; 1.4.1 before
                string memory _version = pod.version();
                bool is1_6 = _strEq(_version, "1.6.0");
                __logTruMoo(" - version is 1.6.0", is1_6);
                if (!is1_6) console.log(" -- (current version: %s)".dim(), _version);
            }

            {
                // 2. In MOOCOW, when a checkpoint is finalized, we should see the currentCheckpoint update
                //    to the final version, with proofsRemaining == 0. Previously, the currentCheckpoint
                //    would not update (so proofsRemaining would be nonzero). This also means that in pods
                //    with activeValidatorCount == 0, starting a checkpoint would finalize it immediately
                //    without a currentCheckpoint update of any kind.
                //
                // - We test this here on our mainnet pod, which does not currently have a checkpoint active.
                //   - First, we manually store 0 into activeValidatorCount.slot
                //   - Then, we read currentCheckpoint()
                //   - Next, we call startCheckpoint()
                //   - Finally, we read currentCheckpoint() again and compare
                //
                // - In MOOCOW, the two read checkpoints will be different. Pre-MOOCOW, they will be the same.

                // activeValidatorCount.slot: 57
                cheats.store(pod_addr, bytes32(uint(57)), 0);

                Checkpoint memory prevCheckpoint = pod.currentCheckpoint();
                {
                    // startCheckpoint calls the EIP-4788 oracle, which needs a valid timestamp to fetch a block root
                    // Here, we warp to a valid timestamp for this one method call, then set it back to its prev value
                    uint ts = block.timestamp;
                    cheats.warp(startTimestamp);
                    pod.startCheckpoint(false);
                    cheats.warp(ts);
                } 
                Checkpoint memory newCheckpoint = pod.currentCheckpoint();
                require(pod.currentCheckpointTimestamp() == 0, "expected checkpoint to be finalized".red().bold());

                __logTruMoo(
                    " - currentCheckpoint updates on completion",
                    keccak256(abi.encode(prevCheckpoint)) != keccak256(abi.encode(newCheckpoint))
                );
            }

            {
                // 3. In MOOCOW, we removed the method "GENESIS_TIME"
                uint64 genesisTime;
                bool methodExists;
                try IOldPod(pod_addr).GENESIS_TIME() returns (uint64 gt) {
                    methodExists = true;
                    genesisTime = gt;
                } catch (bytes memory) {
                    methodExists = false;
                }
                __logTruMoo(" - GENESIS_TIME method removed", !methodExists);
                if (methodExists) console.log(" -- (GENESIS_TIME: %d)".dim(), genesisTime);
            }

            // Interlude: create consolidation/withdrawal requests
            ConsolidationRequest[] memory conReqs = new ConsolidationRequest[](1);
            conReqs[0] = ConsolidationRequest({
                srcPubkey: validatorPubkey,
                targetPubkey: validatorPubkey
            });

            WithdrawalRequest[] memory witReqs = new WithdrawalRequest[](1);
            witReqs[0] = WithdrawalRequest({
                pubkey: validatorPubkey,
                amountGwei: 0
            });

            {
                // 4. In MOOCOW, we should be able to call requestConsolidation and getConsolidationRequestFee
                bool conMethodsExist;
                bool feeRequestSucceeds;
                bool conRequestSucceeds;
                uint conFee;
                try pod.getConsolidationRequestFee() returns (uint fee) {
                    feeRequestSucceeds = true;
                    conFee = fee;
                } catch (bytes memory) {
                    feeRequestSucceeds = false;
                }

                try pod.requestConsolidation{ value: conFee }(conReqs) {
                    conRequestSucceeds = true;
                } catch (bytes memory) {
                    conRequestSucceeds = false;
                }

                // There should be no world where these disagree with each other
                require(feeRequestSucceeds == conRequestSucceeds, "consolidation method existence mismatch".red().bold());
                conMethodsExist = feeRequestSucceeds && conRequestSucceeds;

                __logTruMoo(" - can call consolidation methods", conMethodsExist);
                if (conMethodsExist) console.log(" -- (consolidation request fee: %d)".dim(), conFee);
            }

            {
                // 5. In MOOCOW, we should be able to call requestWithdrawal and getWithdrawalRequestFee
                bool witMethodsExist;
                bool feeRequestSucceeds;
                bool witRequestSucceeds;
                uint witFee;
                try pod.getWithdrawalRequestFee() returns (uint fee) {
                    feeRequestSucceeds = true;
                    witFee = fee;
                } catch (bytes memory) {
                    feeRequestSucceeds = false;
                }

                try pod.requestWithdrawal{ value: witFee }(witReqs) {
                    witRequestSucceeds = true;
                } catch (bytes memory) {
                    witRequestSucceeds = false;
                }

                // There should be no world where these disagree with each other
                require(feeRequestSucceeds == witRequestSucceeds, "withdrawal method existence mismatch".red().bold());
                witMethodsExist = feeRequestSucceeds && witRequestSucceeds;

                __logTruMoo(" - can call withdrawal methods", witMethodsExist);
                if (witMethodsExist) console.log(" -- (withdrawal request fee: %d)".dim(), witFee);
            }

            cheats.stopPrank();
            console.log("");
        }
    }

    function _tryCallEigen() internal {
        {
            // Call EIGEN methods
            console.log("%s %s %s", "Checking".cyan(), "MOOCOW (eigen token)".magenta(), "functions...".cyan());

            address padmin = proxyAdmin.getProxyAdmin(ITransparentUpgradeableProxy(eigen_addr));
            
            // Basic info
            console.log(" - EIGEN proxy:  %s".dim(), eigen_addr);
            __logTruthy(" - proxy admin is correct", padmin == address(proxyAdmin));
            __logTruthy(" - name is Eigen", _strEq(eigen.name(), "Eigen"));
            __logTruthy(" - symbol is EIGEN", _strEq(eigen.symbol(), "EIGEN"));
            __logTruthy(" - totalSupply is expected value", eigen.totalSupply() == 1736730273473851030769230821);
            __logTruthy(" - bEIGEN() is bEIGEN proxy", address(eigen.bEIGEN()) == beigen_addr);

            console.log("");

            {
                // 1. Version string should read 1.6.0 in MOOCOW; should not be queryable before
                bool versionExists;
                string memory _version;
                try eigen.version() returns (string memory v) {
                    versionExists = true;
                    _version = v;
                } catch (bytes memory err) {
                    versionExists = false;
                }

                __logTruMoo(" - version exists and is 1.6.0", versionExists && _strEq(_version, "1.6.0"));
            }

            console.log("");
        }
    }

    function _executeUpgrade() internal {
        _logBanner("Executing upgrade as protocol council.".yellow().bold());
        
        cheats.startPrank(pCMS);

        tcl.execute({
            target: execMS,
            value: 0,
            payload: calldataToExecutor_DISTRO,
            predecessor: bytes32(0),
            salt: bytes32(0)
        });

        tcl.execute({
            target: execMS,
            value: 0,
            payload: calldataToExecutor_MOOCOW,
            predecessor: bytes32(0),
            salt: bytes32(0)
        });

        cheats.stopPrank();
    }

    function _logBanner(string memory str) internal {
        console.log("=================".dim());
        console.log(str);
        console.log("=================".dim());
        console.log("");
    }

    uint PAD_LEN = 50;

    function __logTruthy(string memory str, bool truthy) internal {
        str = _padStr(str, PAD_LEN);
        if (truthy) {
            console.log("%s %s".dim(), str, "TRUE".green());
        } else {
            console.log("%s %s".dim(), str, "FALSE".red());
        }
    }

    // __logTruthy, but logs "MOOCOW ACTIVE" vs "MOOCOW INACTIVE"
    function __logTruMoo(string memory str, bool truthy) internal {
        str = _padStr(str, PAD_LEN);
        if (truthy) {
            console.log("%s %s %s".dim(), str, "TRUE".green(), "(MOOCOW ACTIVE)".green().italic());
        } else {
            console.log("%s %s %s".dim(), str, "FALSE".red(), "(MOOCOW INACTIVE)".red().italic());
        }
    }

    // Pad a string's length with spaces. Revert if str is already > toLength
    // Used to align output on the console so it's pretty
    function _padStr(string memory str, uint toLength) internal pure returns (string memory) {
        require(bytes(str).length <= toLength, "input string length");
        string memory empty = new string(toLength - bytes(str).length);

        for (uint i = 0; i < bytes(empty).length; i++) {
            bytes(empty)[i] = 0x20; // hex representation of " "
        }

        return string.concat(str, empty);
    }

    function _strEq(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }
}