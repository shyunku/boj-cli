# boj-cli

A command-line tool for [Baekjoon Online Judge (BOJ)](https://www.acmicpc.net) that lets you fetch problems, create solution files, and test your code without leaving the terminal.

## Installation

```bash
git clone https://github.com/shyunku/boj-cli.git
cd boj-cli
./install.sh
```

Requires Go 1.22+.

## Commands

### `boj show <problem_id>`
Display a problem by its ID. Fetches from BOJ on first run, then uses local cache.

```bash
boj show 1000
boj show 1000 -t        # also show example test cases
```

### `boj code <problem_id> -l <language>`
Create a solution file in the current directory. The problem statement is included as a comment header.

```bash
boj code 1000 -l nodejs
boj code 1000 -l python
boj code 1000 -l cpp
```

Generated file example (`1000.js`):
```js
// ============================================================
// Problem 1000. A+B
// ============================================================
//
// [Problem]
// ...
// [Test Case 1]
// Input:
//   1 2
// Output:
//   3
//
// ============================================================

const readline = require('readline');
...
function answer(lines) {
  // TODO: solve problem
}
```

### `boj test <problem_id> <file>`
Run a solution file against the problem's example test cases.

```bash
boj test 1000 1000.js
boj test 1000 1000.js -l nodejs   # default is nodejs
```

Output:
```
[PASS] Test case 1
[FAIL] Test case 2
  Expected: "3"
  Got:      "4"

1/2 passed
```

### `boj resume <problem_id>`
Find existing solution files for a problem. If multiple files exist, shows a selection prompt.

```bash
boj resume 1000
```

### `boj langs`
List all supported languages.

```bash
boj langs
```

```
  nodejs     .js
  python     .py
  cpp        .cpp
```

## Caching

Problems are cached at `~/.cache/boj/<problem_id>.json` with a 24-hour TTL. Cached problems load instantly and are available offline. If the cache is stale and the network is unavailable, the stale cache is used with a warning.
