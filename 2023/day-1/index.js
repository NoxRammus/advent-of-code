import { ascend, head, insert, last, sum, sortWith, prop } from "ramda";
import input from "./input.json";

const regexes = [
  { pattern: /one/g, value: "1" },
  { pattern: /two/g, value: "2" },
  { pattern: /three/g, value: "3" },
  { pattern: /four/g, value: "4" },
  { pattern: /five/g, value: "5" },
  { pattern: /six/g, value: "6" },
  { pattern: /seven/g, value: "7" },
  { pattern: /eight/g, value: "8" },
  { pattern: /nine/g, value: "9" },
];

const isNumber = (entry) => {
  if (Number.isNaN(Number(entry))) {
    return false;
  }

  return true;
};

const numbers = input.map((entry) => {
  const updateArray = [];

  const entryUpdates = regexes.reduce((updateArray, regex) => {
    const { pattern, value } = regex;

    let match;

    while ((match = pattern.exec(entry)) !== null) {
      const insertionIndex = match.index;

      updateArray.push({ index: insertionIndex, value });
    }

    return updateArray;
  }, updateArray);

  const sortedUpdates = sortWith([ascend(prop("index"))], entryUpdates);

  const adjustedUpdates = sortedUpdates.map((update, index) => {
    return { ...update, index: update.index + index };
  });

  const correctedEntry = adjustedUpdates.reduce((updatedEntry, update) => {
    const { index, value } = update;

    return insert(index, value, updatedEntry).join("");
  }, entry);

  //  console.log(sortedUpdates);

  //console.log("============================ ENTRY ================= \n", entry);
  //console.log(
  //  "====================== CORRECTED ENTRY ============= \n",
  //  correctedEntry,
  //);

  const entryArray = Array.from(correctedEntry);
  const numericalEntries = entryArray.filter(isNumber);

  return numericalEntries;
});

const concatedNumbers = numbers.map((numberArray) =>
  Number(`${head(numberArray)}${last(numberArray)}`)
);

const result = sum(concatedNumbers);

console.log(result);
