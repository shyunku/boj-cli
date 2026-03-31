# BOJ CLI — Idea

A command-line tool for Baekjoon Online Judge (BOJ) that lets you fetch problems and test solutions without leaving the terminal.

## Commands

### `boj show <problem_id>`
Display a problem by its ID.
- Fetches from BOJ if not cached locally; uses cache on subsequent calls.
- `-tc` / `--testcases`: also print example input/output test cases.

### `boj test <problem_id> <file> -e|--environment <target>`
Run a solution file against the problem's example test cases.
- `-e` / `--environment`: target runtime (e.g. `nodejs`, `python`, `cpp`). Defaults to `nodejs`.
- Compares actual output against expected output for each test case and reports pass/fail.

## Language

**Go**

Chosen for:
- Fast development and iteration
- `cobra` for subcommand structure (`boj show`, `boj test`, ...)
- `goquery` for scraping BOJ problem pages (jQuery-like API)
- Trivial cross-compilation and single binary distribution
- Simple subprocess management for running user solution files
