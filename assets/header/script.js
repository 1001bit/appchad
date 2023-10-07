// Notifications here
const socket = new WebSocket(`ws://${location.hostname}:8000/ws`)

socket.onopen = (event) => {
    console.log("connected")
}

// when server sent a message - add the chat message to a wall
socket.onmessage = (event) => {
    const data = JSON.parse(event.data)
    switch (data.type) {
        case "chat":
            newMessage(data)
            break;
        case "notification":
            console.log(data)
        default:
            break;
    }
}

$(document).ready(() => {
    $(".notif-icon").click(() => {
        const notificationBox = $(".notif-box")[0]
        notificationBox.style.display = notificationBox.style.display == "block" ?
        notificationBox.style.display = "none" : notificationBox.style.display = "block"
    })
})