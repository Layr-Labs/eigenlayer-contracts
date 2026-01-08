---
name: audit-extractor
description: Extract findings from PDF audit reports and convert them to a markdown checklist. Use when the user asks to process an audit report, extract audit findings, or create an audit action items list. (project)
allowed-tools: Read, Write
---

# Audit Extractor

Extract findings from PDF security audit reports and convert them to a structured markdown checklist for tracking remediation.

## Overview

This skill processes PDF audit reports (from firms like Certora, Sigma Prime, Cantina, etc.) and extracts:
- Finding ID (e.g., C-01, H-01, M-01, L-01, I-01)
- Title/Description
- Severity (Critical, High, Medium, Low, Informational)
- Status (Fixed, Partially fixed, Awaiting Fix, Acknowledged, Disregarded, Pending)
- Impact and Likelihood (when available)
- Detailed descriptions and recommendations

## How to Extract Findings

### Step 1: Read the PDF directly

Use the Read tool to read the PDF file. Claude can natively read and understand PDF documents:

```
Read the PDF file at: <pdf_path>
```

### Step 2: Parse the audit content

After reading the PDF, manually extract findings by looking for:

1. **Summary/Findings Table** - Usually contains:
   - Finding IDs (patterns like C-01, H-01, M-01, L-01, I-01)
   - Titles
   - Severity levels
   - Status

2. **Detailed Finding Sections** - Usually organized by PR or category, containing:
   - Full descriptions
   - Code snippets
   - Exploit scenarios
   - Recommendations
   - Customer response / Fix review status

3. **Severity Information** - Look for:
   - Impact (High, Medium, Low)
   - Likelihood (High, Medium, Low)

### Step 3: Generate the markdown output

Create a markdown file with the following structure:

```markdown
# Audit Findings: [Audit Name]

**Auditor:** [Auditor Name]
**Date:** [Date]
**Status:** [Draft/Final]
**Total findings:** [Count]

## Summary

| Severity | Discovered | Confirmed | Fixed |
|----------|------------|-----------|-------|
| Critical | X | | |
| High | X | | |
| Medium | X | | |
| Low | X | | |
| Informational | X | | |
| **Total** | **X** | | |

## Action Required (X items pending)

- [ ] **M-01** (Medium): [Title]
- [ ] **L-01** (Low): [Title]

---

## All Findings

### ([PR/Category Name])

#### [ID]. [Title]
| Severity | Impact | Likelihood | Status |
|----------|--------|------------|--------|
| [Severity] | [Impact] | [Likelihood] | [Status] |

**Files:** `[filename.sol]`

**Description:** [Full description]

**Recommendation:** [Recommendation text]

---
```

### Step 4: Save the output

Use the Write tool to save the markdown file:
- Save to `audits/[Audit-Name]-Findings.md`

## Output Format Guidelines

### Severity Levels
- **Critical** - Direct loss of funds or complete protocol compromise
- **High** - Significant impact on protocol security or functionality
- **Medium** - Moderate impact, potential for exploitation under certain conditions
- **Low** - Minor issues, best practice violations
- **Informational** - Code quality, gas optimizations, documentation

### Status Values
- **Fixed** - Issue has been resolved
- **Partially fixed** - Issue has been partially addressed
- **Awaiting Fix** - Issue acknowledged, fix pending
- **Acknowledged** - Issue acknowledged, may not be fixed
- **Pending** - Awaiting customer response
- **Disregarded** - Issue will not be fixed (by design)

### Checklist Rules
- Use `[ ]` for unresolved items (Pending, Awaiting Fix, Partially fixed)
- Use `[x]` for resolved items (Fixed, Acknowledged, Disregarded)

## Audit Files Location

Audit reports are typically stored in: `audits/`

## Example Extraction

Given a PDF with findings like:

```
M-01 MigrateSlashers may assign incompatible slasher | Medium | Pending
L-01 Instant slasher setting leaves stale field | Low | Pending
I-01 getSlasher is implemented twice | Informational | Pending
```

Generate:

```markdown
# Audit Findings: EigenLayer - Slashing UX Improvements

**Auditor:** Certora
**Date:** December 2025
**Total findings:** 3

## Summary

| Severity | Count |
|----------|-------|
| Medium | 1 |
| Low | 1 |
| Informational | 1 |

## Action Required (2 items)

- [ ] **M-01** (Medium): MigrateSlashers may assign incompatible slasher
- [ ] **L-01** (Low): Instant slasher setting leaves stale field

## All Findings

### Medium (1)

#### M-01. MigrateSlashers may assign an incompatible slasher address
| Severity | Impact | Likelihood | Status |
|----------|--------|------------|--------|
| Medium | High | Low | Pending |

**Description:** [Full description from PDF]

**Recommendation:** [Recommendation from PDF]

---

### Low (1)

#### L-01. Instant slasher setting leaves stale slasher field in storage
| Severity | Impact | Likelihood | Status |
|----------|--------|------------|--------|
| Low | Low | Low | Pending |

**Description:** [Full description from PDF]

**Recommendation:** [Recommendation from PDF]

---

### Informational (1)

#### I-01. getSlasher is implemented twice
| Severity | Status |
|----------|--------|
| Informational | Pending |

**Description:** [Full description from PDF]

**Recommendation:** [Recommendation from PDF]
```

## Troubleshooting

### PDF won't read
- Ensure the file path is correct and the file exists
- Check if the PDF is password protected (not supported)

### Missing findings
- Some PDFs have findings spread across multiple sections
- Check the Table of Contents for finding locations
- Look for "Detailed Findings" or similar sections

### Inconsistent formatting
- Different audit firms use different formats
- Adapt the markdown structure to match the source format
- Preserve all relevant information even if structure differs
