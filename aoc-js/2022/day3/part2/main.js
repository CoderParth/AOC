const fs = require("fs");
const readline = require("readline");

// --- Part Two ---
// As you finish identifying the misplaced items, the Elves come to you
// with another issue.
//
// For safety, the Elves are divided into groups of three. Every Elf carries
// a badge that identifies their group. For efficiency, within each group of
// three Elves, the badge is the only item type carried by all three Elves.
// That is, if a group's badge is item type B, then all three Elves will have
// item type B somewhere in their rucksack, and at most two of the Elves will
// be carrying any other item type.
//
// The problem is that someone forgot to put this year's updated authenticity
// sticker on the badges. All of the badges need to be pulled out of the
// rucksacks so the new authenticity stickers can be attached.
//
// Additionally, nobody wrote down which item type corresponds to each
// group's badges. The only way to tell which item type is the right one is
// by finding the one item type that is common between all three Elves in each
// group.
//
// Every set of three lines in your list corresponds to a single group, but each
// group can have a different badge item type. So, in the above example, the
// first group's rucksacks are the first three lines:
//
// vJrwpWtwJgWrhcsFMMfFFhFp
// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
// PmmdzqPrVvPwwTWBwg
// And the second group's rucksacks are the next three lines:
//
// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
// ttgJtRGJQctTZtZT
// CrZsJsPPZsGzwwsLwLmpwMDw
// In the first group, the only item type that appears in all three rucksacks is
// lowercase r; this must be their badges. In the second group, their badge item
// type must be Z.
//
// Priorities for these items must still be found to organize the sticker attachment
// efforts: here, they are 18 (r) for the first group and 52 (Z) for the second group.
// The sum of these is 70.
//
// Find the item type that corresponds to the badges of each three-Elf group.
// What is the sum of the priorities of those item types?

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

function findCommonItem(groupOfThree) {
  const l = groupOfThree[0].length,
    m = groupOfThree[1].length,
    n = groupOfThree[2].length;

  for (let i = 0; i < l; i++) {
    for (let j = 0; j < m; j++) {
      for (let k = 0; k < n; k++) {
        if (
          groupOfThree[0][i] == groupOfThree[1][j] &&
          groupOfThree[0][i] == groupOfThree[2][k]
        ) {
          return groupOfThree[0][i];
        }
      }
    }
  }
}

function main() {
  const fileScanner = createFileScanner();
  let prioritySum = 0;
  let groupOfThree = [];
  fileScanner.on("line", (line) => {
    if (groupOfThree.length < 3) {
      groupOfThree.push(line);
    }
    if (groupOfThree.length == 3) {
      const item = findCommonItem(groupOfThree);
      prioritySum += itemPriorities[item];
      groupOfThree = [];
    }
  });
  fileScanner.on("close", () => {
    console.log(`Total Sum of the priorities: ${prioritySum}`);
  });
}

main();
