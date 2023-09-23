// Notifications here
var socket = new WebSocket("ws://localhost:8000/ws")

socket.onopen = (event) => {
    console.log("connected")
}