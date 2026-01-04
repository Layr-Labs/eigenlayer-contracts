---
name: audit-extractor
description: Extract findings from PDF audit reports and convert them to a markdown checklist. Use when the user asks to process an audit report, extract audit findings, or create an audit action items list.
allowed-tools: Read, Write, Bash(python3:*, source:*)
---

# Audit Extractor

Extract findings from PDF security audit reports and convert them to a structured markdown checklist for tracking remediation.

## Overview

This skill processes PDF audit reports (from firms like Certora, Sigma Prime, Cantina, etc.) and extracts:
- Finding ID (e.g., C-01, H-01, M-01, L-01, I-01)
- Title/Description
- Severity (Critical, High, Medium, Low, Informational)
- Status (Fixed, Partially fixed, Awaiting Fix, Acknowledged, Disregarded)
- Impact and Likelihood (when available)

## Prerequisites

The extraction requires Python with the `pypdf` library. Set up a virtual environment:

```bash
python3 -m venv /tmp/pdf_venv
source /tmp/pdf_venv/bin/activate
pip install pypdf
```

## Extraction Script

Use this Python script to extract findings from an audit PDF:

```python
import pypdf
import re

def extract_audit_findings(pdf_path):
    """Extract findings from an audit PDF."""
    reader = pypdf.PdfReader(pdf_path)
    
    findings = []
    
    # Try to find the summary table (usually around page 6)
    for page_idx in range(min(10, len(reader.pages))):
        page = reader.pages[page_idx]
        text = page.extract_text()
        if not text:
            continue
        
        # Look for table with ID, Title, Severity, Status
        normalized = re.sub(r'\s+', ' ', text)
        
        # Pattern for summary table entries
        pattern = r'([CHMLIP]-\d+)\s+(.+?)\s+(Critical|High|Medium|Low|Informational)\s+(Fixed\.?|Partially ﬁxed\.?|Partially fixed\.?|Acknowledged(?:\s+Wont\s+Fix)?\.?|Awaiting Fix\.?|Disregarded\.?)'
        
        matches = list(re.finditer(pattern, normalized, re.IGNORECASE))
        
        if matches:
            for match in matches:
                finding_id = match.group(1)
                title = match.group(2).strip().rstrip('.')
                severity = match.group(3)
                status = match.group(4).strip().rstrip('.')
                
                findings.append({
                    'id': finding_id,
                    'title': title,
                    'severity': severity,
                    'status': status
                })
            break  # Found the table, stop looking
    
    # Extract impact/likelihood from detailed sections
    full_text = ""
    for p in reader.pages:
        t = p.extract_text()
        if t:
            full_text += t + "\n"
    
    full_normalized = re.sub(r'\s+', ' ', full_text)
    
    detail_pattern = r'([CHMLIP]-\d+)\s+.+?Severity:\s*(\w+)\s+Impact:\s*(\w+)\s+Likelihood:\s*(\w+)'
    
    detail_map = {}
    for match in re.finditer(detail_pattern, full_normalized, re.IGNORECASE):
        finding_id = match.group(1)
        detail_map[finding_id] = {
            'impact': match.group(3),
            'likelihood': match.group(4)
        }
    
    for f in findings:
        if f['id'] in detail_map:
            f['impact'] = detail_map[f['id']]['impact']
            f['likelihood'] = detail_map[f['id']]['likelihood']
    
    # Sort by severity
    severity_order = {'Critical': 0, 'High': 1, 'Medium': 2, 'Low': 3, 'Informational': 4}
    findings.sort(key=lambda x: (severity_order.get(x['severity'], 5), x['id']))
    
    return findings


def findings_to_markdown(findings, audit_name):
    """Convert findings to markdown checklist."""
    md = f"# Audit Findings: {audit_name}\n\n"
    md += f"**Total findings:** {len(findings)}\n\n"
    
    # Status summary table
    status_counts = {}
    for f in findings:
        stat = f['status']
        status_counts[stat] = status_counts.get(stat, 0) + 1
    
    md += "## Summary\n\n"
    md += "| Status | Count |\n|--------|-------|\n"
    for stat, count in sorted(status_counts.items()):
        md += f"| {stat} | {count} |\n"
    md += "\n"
    
    # Action items section (items needing attention)
    action_needed = [f for f in findings if 'awaiting' in f.get('status', '').lower() or 'partially' in f.get('status', '').lower()]
    if action_needed:
        md += f"## ⚠️ Action Required ({len(action_needed)} items)\n\n"
        for f in action_needed:
            md += f"- [ ] **{f['id']}** ({f['severity']}): {f['title']}\n"
        md += "\n"
    
    # All findings grouped by severity
    md += "## All Findings\n\n"
    for severity in ['Critical', 'High', 'Medium', 'Low', 'Informational']:
        sev_findings = [f for f in findings if f['severity'] == severity]
        if not sev_findings:
            continue
        
        md += f"### {severity} ({len(sev_findings)})\n\n"
        for f in sev_findings:
            # Status emoji
            status_emoji = {
                'Fixed': '✅',
                'Partially ﬁxed': '⚠️',
                'Partially fixed': '⚠️',
                'Acknowledged': '📝',
                'Acknowledged Wont Fix': '🚫',
                'Awaiting Fix': '🔄',
                'Disregarded': '❌'
            }.get(f['status'], '❓')
            
            # Checkbox - checked if resolved
            resolved = f['status'] in ['Fixed', 'Disregarded', 'Acknowledged Wont Fix', 'Acknowledged']
            checkbox = "[x]" if resolved else "[ ]"
            
            md += f"- {checkbox} **{f['id']}**: {f['title']}\n"
            md += f"  - {status_emoji} {f['status']}"
            if f.get('impact'):
                md += f" | Impact: {f['impact']} | Likelihood: {f['likelihood']}"
            md += "\n\n"
    
    return md


# Usage
if __name__ == "__main__":
    import sys
    pdf_path = sys.argv[1] if len(sys.argv) > 1 else "audit.pdf"
    audit_name = sys.argv[2] if len(sys.argv) > 2 else "Audit Report"
    
    findings = extract_audit_findings(pdf_path)
    print(findings_to_markdown(findings, audit_name))
```

## Usage

### Step 1: Set up the environment

```bash
python3 -m venv /tmp/pdf_venv
source /tmp/pdf_venv/bin/activate
pip install pypdf
```

### Step 2: Run extraction

```bash
source /tmp/pdf_venv/bin/activate
python3 extract_audit.py "audits/V1.0.0 (Slashing) - Certora - Feb 2025.pdf" "V1.0.0 Slashing - Certora"
```

### Step 3: Save output

```bash
source /tmp/pdf_venv/bin/activate
python3 extract_audit.py "audits/report.pdf" "Audit Name" > audit_findings.md
```

## Output Format

The script generates markdown with:

1. **Summary table** - Count of findings by status
2. **Action Required section** - Unchecked items needing attention (Awaiting Fix, Partially fixed)
3. **All Findings by Severity** - Grouped checklist with:
   - Checkbox (checked if resolved)
   - Finding ID and title
   - Status emoji and text
   - Impact/Likelihood when available

### Status Emojis

| Status | Emoji | Checkbox |
|--------|-------|----------|
| Fixed | ✅ | [x] |
| Partially fixed | ⚠️ | [ ] |
| Awaiting Fix | 🔄 | [ ] |
| Acknowledged | 📝 | [x] |
| Acknowledged Wont Fix | 🚫 | [x] |
| Disregarded | ❌ | [x] |

## Example Output

```markdown
# Audit Findings: V1.0.0 (Slashing) - Certora

**Total findings:** 9

## Summary

| Status | Count |
|--------|-------|
| Acknowledged Wont Fix | 2 |
| Awaiting Fix | 2 |
| Fixed | 4 |
| Partially ﬁxed | 1 |

## ⚠️ Action Required (3 items)

- [ ] **H-01** (High): The protocol is susceptible to reentrancy attacks
- [ ] **L-03** (Low): OperatorSet decode should be unique
- [ ] **L-04** (Low): burnShares should only work for valid strategies

## All Findings

### Critical (1)

- [x] **C-01**: Underflow when calculating new encumbered magnitude leads to over-risk
  - ✅ Fixed | Impact: High | Likelihood: High

### High (1)

- [ ] **H-01**: The protocol is susceptible to reentrancy attacks
  - ⚠️ Partially ﬁxed | Impact: High | Likelihood: Very Low
```

## Supported Audit Formats

This extraction works with audit reports that follow common formats:
- **Certora** - Security Assessment Reports
- **Sigma Prime** - Security Assessments
- **Cantina** - Audit Reports

The script looks for:
1. A summary table with columns: ID, Title, Severity, Status
2. Detailed sections with: Severity, Impact, Likelihood metadata

## Audit Files Location

Audit reports are stored in: `audits/`

Example files:
- `audits/V1.0.0 (Slashing) - Certora - Feb 2025.pdf`
- `audits/M2 Mainnet - Cantina - Apr 2024.pdf`
- `audits/Rewards v2 - SigmaPrime - Dec 2024.pdf`

## Troubleshooting

### No findings extracted
- The PDF may have a different structure. Check if the summary table exists.
- Try adjusting the `page_idx` range if the table is on a different page.

### Garbled titles
- PDF text extraction can sometimes merge lines. The script tries to handle this but may need adjustment for specific report formats.

### Missing impact/likelihood
- Not all reports include this metadata. The script will show what's available.
