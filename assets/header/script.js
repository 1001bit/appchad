const socket = new WebSocket(`ws://${location.hostname}:8000/ws`)
const notifBox = $("#notif-box")

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
            notifications(data)
        default:
            break;
    }
}

$(document).ready(() => {
    notifications(0)
})

// on notifications button click
$(".notif-icon").click(() => {
    switch(notifBox[0].style.display){
        case "none":
            notifBox[0].style.display = "block"
            $(".notif-icon").toggleClass("new", false);
            break;
        default:
            notifBox[0].style.display = "none"
            break
    }
})

// make new notification element
function makeNewNotification(data){
    const notif = $("<div></div>").addClass("notif")

    const date = $("<pre></pre>").text(data['date'])
    notif.append(date)

    // what type of notification
    switch(data["nType"]){
        case "chatReply":
            const messageData = data["messageData"]
            const username = $("<a></a>").text(`${messageData['username']} replied:`).attr("href", `/chatchad?id=${messageData["id"]}`)
            const text = $("<pre></pre>").text(messageData['text'])

            notif.append(username)
            notif.append(text)
            break;
        default:
            break;
    }

    return notif
}

// add messages to box from localstorage
function notifications(data){
    let storedNotifications = JSON.parse(localStorage.getItem("notifications"))

    // add new notification to localstorage
    if (data != 0){
        if(!storedNotifications) storedNotifications = []
        storedNotifications.push(data)
        if(storedNotifications.length > 50) storedNotifications.slice(1)
        localStorage.setItem("notifications", JSON.stringify(storedNotifications))
        // change color of button if new messages
        if (notifBox[0].style.display == "none"){
            $(".notif-icon").toggleClass("new", true);
        }
    }

    if (!storedNotifications) return

    // add notifications to box
    const fragment = document.createDocumentFragment()
    // make readable notifications from data
    for(let i = storedNotifications.length-1; i >= 0; i--){
        fragment.append(makeNewNotification(storedNotifications[i])[0]);
    }
    notifBox.empty()
    notifBox.append(fragment)
}