"use strict";

const API = "http://localhost:8080";
let field = Array(3).fill().map(Array(3).fill);
const board = document.getElementById("board");

function createBoard() {
  console.log("42134312");
  board.innerHTML = "";
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      const cell = document.createElement("div");
      cell.className = "cell";
      cell.dataset.index = i * 3 + j;
      // cell.addEventListener("click", () => updateGame());
      board.appendChild(cell);
    }
  }
}

createBoard();
