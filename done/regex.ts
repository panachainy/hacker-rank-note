let input = readLine();
const replaceChar = "_________";
const regex = /hackerrank/g;

const found = input.match(regex);
console.log(`Number of matches : ${found.length}`);

//============

let input = readLine();
const regex = /^(\w{3}).(\w{3}).(\w{3}).(\w{3})$/g;

const found = input.match(regex);

if (found) {
  console.log(true);
} else {
  console.log(false);
}
