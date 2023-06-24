#!python3

import argparse
import os
import re
from functools import reduce
from itertools import chain
from pathlib import Path
from sys import stderr
from typing import Any, Dict, Optional
from typing import Tuple

METHOD_NAME_RE = r'(?P<name>[a-zA-Z_][a-zA-Z_0-9.]*)'
sinvoke_re = re.compile(rf'(\bsinvoke\s+{METHOD_NAME_RE})\s*\(')
invoke_re = re.compile(rf'(\binvoke\s+{METHOD_NAME_RE})\s*\(')
# Some funny stuff here in order to properly parse 'if(...foo().selector)'
selector_re = re.compile(rf'[^:]\b{METHOD_NAME_RE}(?=(\s*\([^)]*\)\.selector))')
# Find the start of the methods block. Ignore methods block that may be all on one
# line - this confuses the rest of the script and is not worth it.
methods_block_re = re.compile(r'methods\s*{\s*(?://.*)?$', re.MULTILINE)
method_line_re = re.compile(rf'\s*(?:/\*.*\*/)?\s*(?P<func>function)?\s*(?P<name_w_paren>{METHOD_NAME_RE}\s*\()')
rule_prefix_re = re.compile(rf'^{METHOD_NAME_RE}+\s*(\([^)]*\))?\s*{{?\s*$')
double_sig = re.compile(rf'sig:{METHOD_NAME_RE}.sig:')

def fixup_sinvoke(line: str) -> Tuple[str, int]:
    matches = sinvoke_re.findall(line)
    # stat[] += len(matches)
    return (reduce(lambda l, m: l.replace(m[0], f'{m[1]}'),
                   matches,
                   line),
            len(matches))


def fixup_invoke(line: str) -> Tuple[str, int]:
    matches = invoke_re.findall(line)
    return (reduce(lambda l, m: l.replace(m[0], f'{m[1]}@withrevert'),
                   matches,
                   line),
            len(matches))


def fixup_selector_sig(line: str) -> Tuple[str, int]:
    matches = selector_re.findall(line)
    matches = list(filter(lambda m: m[0] not in ('if', 'require', 'assert'), matches))
    l, n = (reduce(lambda l, m: l.replace(m[0] + m[1], f'sig:{m[0] + m[1]}'),
                   matches,
                   line),
            len(matches))
    # fix double sig
    double_sig_matches = double_sig.findall(l)
    for match in double_sig_matches:
        l = l.replace(f'sig:{match}.sig:', f'sig:{match}.')
    return l, n


def fixup_rule_prefix(line: str) -> Tuple[str, int]:
    matches = rule_prefix_re.match(line)
    if matches and matches['name'] != "methods":
        return 'rule ' + line, 1
    return line, 0


def fixup_static_assert(line: str) -> Tuple[str, int]:
    if line.lstrip().startswith('static_assert '):
        return (line.replace('static_assert', 'assert', 1), 1)
    return (line, 0)


def fixup_static_require(line: str) -> Tuple[str, int]:
    if line.lstrip().startswith('static_require '):
        return (line.replace('static_require', 'require', 1), 1)
    return (line, 0)


def find_invoke_whole(line: str) -> bool:
    return 'invoke_whole(' in line


def methods_block_add_semicolon(line: str, next_line: Optional[str]) -> Tuple[str, int]:
    l, *comment = line.split('//', 1)
    l = l.rstrip()
    do_nothing = (line, 0)
    if not l.lstrip():
        # an empty line
        return do_nothing

    if any(l.endswith(s) for s in (';', '=>', '(', '{', '/*', '/**', '*/', ',')):
        # this line doesn't need a semicolon
        return do_nothing

    if any(w in l.split() for w in ('if', 'else')):
        # this is a branching line, skip it
        return do_nothing

    if next_line is not None and (next_line.lstrip().startswith("=>") or next_line.lstrip().startswith(')')):
        # the method's summary is defined in the next line, don't append a ;
        # also if we have a ) in the next line it means we broke lines for the parameters
        return do_nothing

    return l + ';' + (f' //{comment[0]}' if comment else ''), 1


def methods_block_prepend_function(line: str) -> Tuple[str, int]:
    m = method_line_re.match(line)
    if m is not None and m['func'] is None:
        return line.replace(m['name_w_paren'], f'function {m["name_w_paren"]}', 1), 1
    return line, 0


def methods_block_add_external_visibility_no_summary(line: str) -> Tuple[str, int]:
    m = method_line_re.match(line)
    if m is not None and ('=>' not in line or '=> DISPATCHER' in line):
        replacables = [' returns ', ' returns(', ' envfree', ' =>', ';']
        for r in replacables:
            if r in line:
                if all(f' {vis}{r2}' not in line for vis in ('internal', 'external') for r2 in replacables):
                    line = line.replace(r, f' external{r}')
                    return line, 1
                return line, 0
        else:
            print(f"Unable to add 'external' modifier to {line}")
    return line, 0


def methods_block_summary_should_have_wildcard(line: str) -> Tuple[str, int]:
    m = method_line_re.match(line)
    if m is not None and '=>' in line and '.' not in line and ' internal ' not in line:
        line = line.replace("function ", "function _.")
        return line, 1

    return line, 0


def append_semicolons_to_directives(line: str) -> Tuple[str, int]:
    if line.lstrip().startswith(('pragma', 'import', 'using', 'use ')) and not line.rstrip().endswith((';', '{')):
        line = line.rstrip() + ';' + os.linesep
        return line, 1
    return line, 0


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--files', metavar='FILE', nargs='+', help='list of files to change', type=Path,
                        default=list())
    parser.add_argument('-d', '--dirs', metavar='DIR', nargs='+', help='the dir to search', type=Path, default=list())
    parser.add_argument('-r', '--recursive', action='store_true', help='if set dirs will be searched recursively')
    args = parser.parse_args()

    if not args.dirs and not args.files:
        print('No files/dirs specified', file=stderr)
        return 1
    if non_existent := list(filter(lambda f: not (f.exists() and f.is_file()), args.files)):
        print(f'Cannot find files {non_existent}', file=stderr)
        return 1
    if non_existent := list(filter(lambda d: not (d.exists() and d.is_dir()), args.dirs)):
        print(f'Cannot find dirs {non_existent}', file=stderr)
        return 1

    spec_files: chain = chain()
    for d in args.dirs:
        dir_files_unfiltered = (f for f in d.glob(f'{"**/" if args.recursive else ""}*')
                                if f.suffix in ('.cvl', '.spec'))
        dir_files = filter(
            lambda fname: not any(p.startswith('.certora_config') for p in fname.parts),
            dir_files_unfiltered)
        spec_files = chain(spec_files, dir_files)
    spec_files = chain(spec_files, args.files)

    sinvoke_str = 'sinvoke foo(...) -> foo(...)'
    invoke_str = 'invoke foo(...) -> foo@withrevert(...)'
    static_assert_str = 'static_assert -> assert'
    static_require_str = 'static_require -> require'
    invoke_whole_str = "lines with 'invoke_whole'"
    sig_selector_str = "selectors prepended with 'sig:'"
    rule_prefix_str = "rule declarations prepended with 'rule'"
    semicolon_in_directives_str = "'pragma', 'import', 'using' or 'use' directives appended with ';'"
    semicolon_in_methods_str = "lines in 'methods' block appended with ';'"
    prepend_function_for_methods_str = "lines in 'methods' block prepended with 'function'"
    add_external_visibility_for_non_summary_str = "declarations in 'methods' block with 'external' visibility added"
    summary_should_have_wildcard_str = "declarations in 'methods' block with wildcard added"

    stats: Dict[Any, Any] = {}
    for fname in spec_files:
        print(f"processing {fname}")
        stats[fname] = {}
        stats[fname][sinvoke_str] = 0
        stats[fname][invoke_str] = 0
        stats[fname][static_assert_str] = 0
        stats[fname][static_require_str] = 0
        stats[fname][sig_selector_str] = 0
        stats[fname][rule_prefix_str] = 0
        stats[fname][invoke_whole_str] = []
        stats[fname][semicolon_in_directives_str] = 0
        stats[fname][semicolon_in_methods_str] = 0
        stats[fname][prepend_function_for_methods_str] = 0
        stats[fname][add_external_visibility_for_non_summary_str] = 0
        stats[fname][summary_should_have_wildcard_str] = 0

        flines = open(fname).readlines()
        for i, l in enumerate(flines):
            l, num = fixup_sinvoke(l)
            stats[fname][sinvoke_str] += num

            l, num = fixup_invoke(l)
            stats[fname][invoke_str] += num

            l, num = fixup_static_assert(l)
            stats[fname][static_assert_str] += num

            l, num = fixup_static_require(l)
            stats[fname][static_require_str] += num

            l, num = append_semicolons_to_directives(l)
            stats[fname][semicolon_in_directives_str] += num

            l, num = fixup_selector_sig(l)
            stats[fname][sig_selector_str] += num
            flines[i] = l

            l, num = fixup_rule_prefix(l)
            stats[fname][rule_prefix_str] += num
            flines[i] = l

        with open(fname, 'w') as f:
            f.writelines(flines)

        # methods block
        contents = open(fname).read()
        methods_block_declaration = methods_block_re.search(contents)
        if not methods_block_declaration:
            print(f"{fname} has no methods block, skipping...")
            continue
        try:
            prev, methods_block = methods_block_re.split(contents)
        except ValueError:
            print(f"{fname}: Failed to find methods block - more than one 'methods' block found")
            continue

        try:
            methods_block, rest = re.split(r'^}', methods_block, maxsplit=1, flags=re.MULTILINE)
        except ValueError:
            print(f"{fname}: Failed to find methods block - couldn't find block end")
            continue

        if methods_block.count("{") != methods_block.count("}"):
            print(f"{fname}: Failed to find methods block - something went wrong with finding the end of the block")
            continue

        # add semicolons in the methods block where it's easy to do so
        methods_block = methods_block.split("\n")
        for i, l in enumerate(methods_block):
            next_line = methods_block[i + 1] if i + 1 < len(methods_block) else None
            l, n = methods_block_add_semicolon(l, next_line)
            stats[fname][semicolon_in_methods_str] += n
            l, n = methods_block_prepend_function(l)
            stats[fname][prepend_function_for_methods_str] += n
            l, n = methods_block_add_external_visibility_no_summary(l)
            stats[fname][add_external_visibility_for_non_summary_str] += n
            l, n = methods_block_summary_should_have_wildcard(l)
            stats[fname][summary_should_have_wildcard_str] += n
            methods_block[i] = l
        methods_block = "\n".join(methods_block)

        with open(fname, 'w') as f:
            f.write(prev + methods_block_declaration.group(0) + methods_block + '}' + rest)

    print('Change Statistics\n-----------------')
    for fname, stat in stats.items():
        if all(not n for n in stat.values()):
            continue
        print(f'{fname}:')
        for s, n in stat.items():
            if not n:
                continue
            print(f'\t{s}:  {n}')

    print("The following files have instances of 'invoke_whole' which need to be removed manually")
    for fname in stats:
        inv_whole_list = stats[fname][invoke_whole_str]
        if not inv_whole_list:
            continue
        print(f"\t{fname} (line{'s' if len(inv_whole_list) > 1 else ''} {', '.join(str(n) for n in inv_whole_list)})")
    return 0


if __name__ == "__main__":
    exit(main())