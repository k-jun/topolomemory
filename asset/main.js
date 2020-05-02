const socket = io("ws://localhost:8000");


const current_status = document.getElementById("status")
const header = document.getElementById("header")
const contence = document.getElementById("contence")
const footer = document.getElementById("footer")

let isInRoom = false

// send
function IndexScreen() {
  header.innerHTML = `
    <button id="room-create">Room Create</button>
  `
  document.getElementById("room-create").addEventListener("click", () => {
    socket.emit("room/create")
  })
}

// receive
socket.on("room/index", () => {

})

// index screen
// room/index
// room/create
// room/join








// game screen
// room/leave


IndexScreen()



socket.on("status", (payload) => {
  console.log("current status", payload)
  current_status.innerHTML = payload

})
