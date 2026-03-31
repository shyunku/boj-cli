// ============================================================
// Problem 1000. A+B
// ============================================================
//
// [Problem]
// 두 정수 A와 B를 입력받은 다음, A+B를 출력하는 프로그램을 작성하시오.
//
// [Input]
// 첫째 줄에 A와 B가 주어진다. (0 < A, B < 10)
//
// [Output]
// 첫째 줄에 A+B를 출력한다.
//
// [Test Case 1]
// Input:
//   1 2
// Output:
//   3
//
// ============================================================

const readline = require('readline');

const rl = readline.createInterface({ input: process.stdin });
const lines = [];
rl.on('line', line => lines.push(line.trim()));
rl.on('close', () => {
  console.log(answer(lines));
});

function answer(lines) {
  // TODO: solve problem
}
