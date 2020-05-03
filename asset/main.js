const socket = io("ws://localhost:8000");


const current_status = document.getElementById("status")
const header = document.getElementById("header")
const contence = document.getElementById("contence")
const footer = document.getElementById("footer")
const gameStartButton = document.getElementById("game-start")
const gameDrawButton = document.getElementById("game-draw")
const gameUserName = document.getElementById("username")
const gameUserNameInput = document.getElementById("username-input")
  // document.getElementById('name-input').addEventListener('change', function(e) {
  //   e.preventDefault();
  //   socket.emit('name-update', e.target.value)
  // })

let aId = 0
let bId = 0

socket.on("game/status", (data) => {
  console.log("game/status")
  json = JSON.parse(data)
  console.log(json)
  html = ""
  footer.innerHTML = json.Score
  if (json.Field == null) {
    contence.innerHTML = html
    return

  }
  for (let i = 0; i < json.Field.length; i++) {
    html += `<div class="card" cardId="${json.Field[i].Id}">${json.Field[i].Svg}</div> `
  }
  contence.innerHTML = html
  AddEventListenerToCards()
  aId = 0
  bId = 0
})


socket.on("game/win", () => {
  alert("you got point!!")
})
socket.on("game/lose", () => {
  alert("you could not get point...")
})

socket.on("username/change", (payload) => {
  gameUserName.innerHTML = payload
})

function AddEventListenerToCards() {
  cards = document.getElementsByClassName("card")
  for (let i = 0; i < cards.length; i++) {
    cards[i].addEventListener("click", () => {
      cardId = cards[i].getAttribute("cardId")
      if (aId == 0 && bId == 0) {
        aId = cardId
        console.log("aId", aId, "bId", bId)
      }
      if (aId != cardId && aId != 0 && bId == 0) {
        bId = cardId
        console.log("aId", aId, "bId", bId)
        socket.emit("game/hit", `${aId},${bId}`)
      }
    })
  }
}

function init() {
  gameStartButton.addEventListener("click", (e) => {
    e.preventDefault()
    console.log("game/start")
    socket.emit("game/start", "")
  })
  gameDrawButton.addEventListener("click", (e) => {
    e.preventDefault()
    console.log("game/draw")
    socket.emit("game/draw", "")
  })
  gameUserNameInput.addEventListener("change", (e) => {
    e.preventDefault()
    console.log("username/change", e.target.value)
    socket.emit("username/change", e.target.value)
  })
}




init()


// game screen
// room/leave

