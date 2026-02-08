"use strict";

const API = "http://localhost:8080";
let fieldBIS = [];
for (let i = 0; i < 3; i++) {
  fieldBIS[i] = Array(3).fill("");
}

let UUIDBIS = "-1";
const statusBarDOM = document.getElementById("status_bar");

let sideBIS = "unknown";
const fieldDOM = document.getElementById("board");

function initBoard() {
  console.log("init board");
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      const cell = document.createElement("div");
      cell.className = "cell";
      // × ✖ ○ 𐤏
      // cell.textContent = "×";
      // if ((i * 3 + j) % 2 == 0) {
      //   cell.textContent = "○";
      // }
      cell.dataset.index = i * 3 + j + 1;
      cell.addEventListener("click", () => updateGame());
      fieldDOM.appendChild(cell);
    }
  }
}

function initStatusBar() {
  console.log("init status bar");
  const sideDOM = document.createElement("div");
  sideDOM.className = "side";
  sideDOM.textContent = `Your side is ${sideBIS}`;
  statusBarDOM.appendChild(sideDOM);

  const idDOM = document.createElement("div");
  idDOM.className = "id";
  idDOM.textContent = `Game ID: ${UUIDBIS}`;
  statusBarDOM.appendChild(idDOM);
}

function decodeSide(side) {
  if (side == 1) {
    return "×";
  } else if (side == 2) {
    return "○";
  }
  return "";
}

function updateField() {
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      fieldDOM.children[i * 3 + j].textContent = decodeSide(fieldBIS[i][j]);
    }
  }
}

async function newGame() {
  const response = await fetch(`${API}/game/create_game`, { method: "POST" });
  const data = await response.json();

  UUIDBIS = data.uuid;
  fieldBIS = data.field;
  console.log(fieldBIS);
  sideBIS = decodeSide(data.player_side);
  updateDataBIS(fieldBIS, UUIDBIS);

  statusBarDOM.querySelector(".id").textContent = `Game ID: ${UUIDBIS}`;
  statusBarDOM.querySelector(".side").textContent = `Your side is ${sideBIS}`;
  updateField();
}

const dataBIS = {
  field: fieldBIS,
  uuid: UUIDBIS,
};

function updateDataBIS(fieldBIS, UUIDBIS) {
  dataBIS.field = fieldBIS;
  dataBIS.uuid = UUIDBIS;
}

async function updateGame() {
  const response = await fetch(`${API}/game/${UUIDBIS}`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(dataBIS),
  });

  const data = await response.json();
  console.log(data);
}

initBoard();
initStatusBar();
