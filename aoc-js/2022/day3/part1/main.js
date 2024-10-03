const fs = require("fs");
const readline = require("readline");

const itemPriorities = {
  a: 1,
  b: 2,
  c: 3,
  d: 4,
  e: 5,
  f: 6,
  g: 7,
  h: 8,
  i: 9,
  j: 10,
  k: 11,
  l: 12,
  m: 13,
  n: 14,
  o: 15,
  p: 16,
  q: 17,
  r: 18,
  s: 19,
  t: 20,
  u: 21,
  v: 22,
  w: 23,
  x: 24,
  y: 25,
  z: 26,
  A: 27,
  B: 28,
  C: 29,
  D: 30,
  E: 31,
  F: 32,
  G: 33,
  H: 34,
  I: 35,
  J: 36,
  K: 37,
  L: 38,
  M: 39,
  N: 40,
  O: 41,
  P: 42,
  Q: 43,
  R: 44,
  S: 45,
  T: 46,
  U: 47,
  V: 48,
  W: 49,
  X: 50,
  Y: 51,
  Z: 52,
};

function createFileScanner() {
  return readline.createInterface({
    input: fs.createReadStream("input.txt"),
  });
}

function findCommonItem(line) {
  const n = line.length;
  const firstCompart = line.slice(0, n / 2);
  const secondCompart = line.slice(n / 2, n);

  for (let i = 0; i < firstCompart.length; i++) {
    for (let j = 0; j < secondCompart.length; j++) {
      if (firstCompart[i] == secondCompart[j]) {
        return firstCompart[i];
      }
    }
  }
}

function main() {
  const fileScanner = createFileScanner();
  let prioritySum = 0;
  fileScanner.on("line", (line) => {
    const item = findCommonItem(line);
    prioritySum += itemPriorities[item];
  });
  fileScanner.on("close", () => {
    console.log(`Total Sum of the priorities: ${prioritySum}`);
  });
}

main();
