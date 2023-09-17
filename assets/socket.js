var socket = new WebSocket("ws://localhost:8000/ws")

socket.onopen = (event) => {
    socket.send("socket opened")
    console.log("connected")
}