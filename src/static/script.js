"use strict";

const API = "http://localhost:8080";

const dataBIS = {
  field: [],
  uuid: "",
};

function initDataBIS(dataBIS) {
  for (let i = 0; i < 3; i++) {
    dataBIS.field[i] = Array(3).fill("");
  }
  console.log(dataBIS);
}
initDataBIS(dataBIS);

function updateDataBIS(dataBIS, field, uuid) {
  dataBIS.field = field;
  dataBIS.uuid = uuid;
}

const statusBarBIS = {
  side: undefined,
  whosMoving: undefined,
  status: undefined,
};

const fieldDOM = document.getElementById("board");
function initBoard() {
  console.log("init board");
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      const cell = document.createElement("div");
      cell.className = "cell";
      // × ✖ ○ 𐤏
      cell.dataset.index = i * 3 + j + 1;
      cell.addEventListener("click", () => updateGame());
      fieldDOM.appendChild(cell);
    }
  }
}

const statusBarDOM = document.getElementById("status_bar");
// init
function initStatusBarDOM(dataBIS, statusBarBIS) {
  console.log("init status bar DOM");
  const sideDOM = document.createElement("div");
  sideDOM.className = "side";
  sideDOM.textContent = `Your side is ${statusBarBIS.side}`;
  statusBarDOM.appendChild(sideDOM);

  const whosMovingDOM = document.createElement("div");
  whosMovingDOM.className = "whos_moving";
  whosMovingDOM.textContent = `Who's moving?: ${statusBarBIS.whosMoving}`;
  statusBarDOM.appendChild(whosMovingDOM);

  const statusDOM = document.createElement("div");
  statusDOM.className = "status";
  statusDOM.textContent = `Status: ${statusBarBIS.status}`;
  statusBarDOM.appendChild(statusDOM);

  const idDOM = document.createElement("div");
  idDOM.className = "id";
  idDOM.textContent = `Game ID: ${dataBIS.uuid}`;
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
      fieldDOM.children[i * 3 + j].textContent = decodeSide(
        dataBIS.field[i][j],
      );
    }
  }
}

// new game
async function newGame(dataBIS) {
  const response = await fetch(`${API}/game/create_game`, { method: "POST" });
  const data = await response.json();
  console.log(data.field);
  console.log(data.uuid);

  console.log(dataBIS);
  console.log(dataBIS.field);
  console.log(dataBIS.uuid);
  updateDataBIS(dataBIS, data.field, data.uuid);

  statusBarBIS.side = decodeSide(data.player_side);

  statusBarDOM.querySelector(".id").textContent = `Game ID: ${dataBIS.uuid}`;
  statusBarDOM.querySelector(".side").textContent =
    `Your side is ${statusBarBIS.side}`;
  updateField();
}

// update game
async function updateGame() {
  whosMovingDOM.textContent = "Who's moving: Computer";

  const response = await fetch(`${API}/game/${UUIDBIS}`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(dataBIS),
  });

  const data = await response.json();
  console.log(data);
  whosMovingDOM.textContent = "Who's moving: Me";
}

initBoard();
initStatusBarDOM(dataBIS, statusBarBIS);
