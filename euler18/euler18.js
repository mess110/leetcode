const fs = require("fs");
const readline = require("readline");

let array = [];

const inputStream = fs.createReadStream('67.txt');
var lineReader = readline.createInterface({
  input: inputStream,
  terminal: false,
});
lineReader.on("line", function (line) {
    array.push(line.split(" ").map(e => Number(e)))
});

lineReader.on('close', function () {
    solve(array);
});

// let solve = (array) => {
//     while (array.length != 1) {
//         array = calcRow(array)
//     }
//     console.log(array)
// }

// let calcRow = (a) => {
//     let lastRow = a.pop()
//     let secondLastRow = a.pop()
//     let newRow = []
//     for (let i = 0; i < secondLastRow.length; i++) {
//         let p1 = secondLastRow[i] + lastRow[i];
//         let p2 = secondLastRow[i] + lastRow[i+1];
//         newRow.push(p1 > p2 ? p1 : p2)
//     }
//     a.push(newRow)
//     return a
// }

let solve = (array) => {
    console.log(calcRow(array))
}

let calcRow = (a) => {
    let lastRow = a.pop()
    let secondLastRow = a.pop()
    let newRow = []
    for (let i = 0; i < secondLastRow.length; i++) {
        let p1 = secondLastRow[i] + lastRow[i];
        let p2 = secondLastRow[i] + lastRow[i+1];
        newRow.push(p1 > p2 ? p1 : p2)
    }
    a.push(newRow)

    if (a.length != 1) {
        return calcRow(a)
    } else {
        return a
    }
}

