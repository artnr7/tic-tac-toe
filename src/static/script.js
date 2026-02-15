"use strict";

const API = "http://localhost:8080";

const dataBIS = {
  field: [],
  uuid: undefined,
  player_status: undefined,
  player_side: undefined,
  whos_moving: undefined,
};

function initDataBIS(dataBIS) {
  for (let i = 0; i < 3; i++) {
    dataBIS.field[i] = Array(3).fill("");
  }
}

function updateDataBIS(dataBIS, data) {
  dataBIS.field = data.field;
  dataBIS.uuid = data.uuid;

  if (data.status == 0) {
    dataBIS.player_status = 2;
  } else if (data.status == 2) {
    dataBIS.player_status = 0;
  } else {
    dataBIS.player_status = data.status;
  }

  dataBIS.player_side = data.player_side;
}

const newGameBtn = document.getElementById("new_game_btn");
newGameBtn.onclick = () => newGame(dataBIS);

const fieldDOM = document.getElementById("board");
function initBoard() {
  console.log("init board");
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      const cell = document.createElement("div");
      cell.className = "cell";
      // × ✖ ○ 𐤏
      cell.dataset.index = i * 3 + j;
      cell.addEventListener("click", () => updateGame(i, j));
      fieldDOM.appendChild(cell);
    }
  }
}

const statusBarDOM = document.getElementById("status_bar");
const sideDOM = document.getElementById("side");
const whosMovingDOM = document.getElementById("whos_moving");
const statusDOM = document.getElementById("status");
const uuidDOM = document.getElementById("uuid");
// init
function initStatusBarDOM(dataBIS) {
  console.log("init status bar DOM");
  sideDOM.textContent = `Your side is ${dataBIS.player_side}`;
  whosMovingDOM.textContent = `Who's moving?: ${dataBIS.whos_moving}`;
  statusDOM.textContent = `Status: ${dataBIS.player_status}`;
  uuidDOM.textContent = `Game ID: ${dataBIS.uuid}`;
}

function decodeSide(side) {
  switch (side) {
    case 1:
      return "×";
    case 2:
      return "○";
    default:
      return "";
  }
}

function decodeStatus(status) {
  switch (status) {
    case 0:
      return "Defeat";
    case 1:
      return "Draw";
    case 2:
      return "Victory";
    case 3:
      return "Active";
  }
}

function decodeWhosMoving(whos_playing) {
  switch (whos_playing) {
    case 0:
      return "Computer";
    case 1:
      return "Me";
  }
}

function updateField(dataBIS) {
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      fieldDOM.children[i * 3 + j].textContent = decodeSide(
        dataBIS.field[i][j],
      );
    }
  }
}

function updateDOM(dataBIS) {
  sideDOM.textContent = `Your side: ` + decodeSide(dataBIS.player_side);
  whosMovingDOM.textContent =
    `Who's moving?: ` + decodeWhosMoving(dataBIS.whos_moving);
  statusDOM.textContent = "Status: " + decodeStatus(dataBIS.player_status);
  uuidDOM.textContent = uuidDOM.textContent = `Game ID: ${dataBIS.uuid}`;
  updateField(dataBIS);
}
const WhosMovingState = Object.freeze({ Computer: 0, Me: 1 });

// new game
async function newGame(dataBIS) {
  const response = await fetch(`${API}/game/create_game`, { method: "POST" });
  const data = await response.json();
  updateDataBIS(dataBIS, data);
  dataBIS.whos_moving = WhosMovingState.Me;
  updateDOM(dataBIS);
}

const request = {
  field: dataBIS.field,
  uuid: dataBIS.uuid,
};

// update game
async function updateGame(i, j) {
  if (
    dataBIS.uuid == undefined ||
    dataBIS.field[i][j] != 0 ||
    dataBIS.whos_moving != WhosMovingState.Me ||
    dataBIS.player_status != 3
  ) {
    return;
  }

  // business
  dataBIS.field[i][j] = dataBIS.player_side;
  dataBIS.whos_moving = WhosMovingState.Computer;
  updateDOM(dataBIS);

  // request
  request.field = dataBIS.field;
  request.uuid = dataBIS.uuid;

  const response = await fetch(`${API}/game/${request.uuid}`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(request),
  });
  const pr = response.json();
  pr.then(() => {
    dataBIS.whos_moving = WhosMovingState.Me;
  });

  const data = await pr;
  console.log(data.status);

  updateDataBIS(dataBIS, data);
  updateDOM(dataBIS);
}

initDataBIS(dataBIS);
initBoard();
initStatusBarDOM(dataBIS);
