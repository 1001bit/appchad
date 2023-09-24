// Notifications here
const socket = new WebSocket(`ws://${location.hostname}:8000/ws`)

socket.onopen = (event) => {
    console.log("connected")
}