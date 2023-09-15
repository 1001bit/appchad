const updateTime = 3000;
let interval = window.setInterval(chatGet, updateTime);
let lastMessageID = 0

const chatBox = $("#chat");
const typebox = $("#typebox")

$(document).ready(() => {
    chatGet().then(() => {
        chatBox.scrollTop(chatBox[0].scrollHeight);
    })
})

// fetch api and add messages to page
function chatGet(){
    return fetch(`/api/chatchad?id=${lastMessageID}`, {
        method: "GET"
    })
    .then(response => response.json())
    .then(data => {
        if (data.length > 0){
            lastMessageID = data[data.length - 1].id;
            addNewMessages(data)
        }
    })
    .catch(error => {
        console.log(error)
    })
}

// post message to chat
async function chatPost(msgText){
    return fetch("/api/chatchad", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({text: msgText}),
    })
    .catch(error => {
        console.log(error)
    })
}

// add messages to page from data
function addNewMessages(data){
    const doScroll = chatBox.scrollTop() + chatBox.innerHeight() >= chatBox[0].scrollHeight-1;
    const fragment = document.createDocumentFragment()

    for(let i = 0; i < data.length; i++){
        const message = $("<div></div>").addClass("message").attr("id", data[i]['id']);
        const date = $("<pre></pre>").text(data[i]['date']);
        const user = $("<a></a>").text(`${data[i]['username']}:`).attr("href", "/chad/"+data[i]["userid"]);
        const text = $("<pre></pre>").html(data[i]['text'])
        message.append(date);
        message.append(user)
        message.append(text);
        fragment.append(message[0]);
    }
    chatBox.append(fragment)

    // automatically scroll down
    if(doScroll){
        chatBox.scrollTop(chatBox[0].scrollHeight);
    }
}

// submit button
$(".send").click(async () => {
    if(typebox.val().length == 0){
        return
    }

    const text = typebox.val()
    typebox.val("")

    window.clearInterval(interval)
    interval = window.setInterval(chatGet, updateTime);
    
    chatPost(text).then(() => {
        chatGet()
    })
})

//////////////////
// STYLES
function getCursorPosition(element, event) {
        const rect = element.getBoundingClientRect();
        const centerX = rect.left + rect.width / 2;
        const centerY = rect.top + rect.height / 2;
        const x = event.clientX - centerX;
        const y = centerY - event.clientY;
        return { x, y };
    }
    
    const buttons = document.querySelectorAll("button");
    [...buttons].map((button) => {
        button.addEventListener("pointermove", (event) => {
        const { x, y } = getCursorPosition(event.target, event);
        button.style.setProperty("--coord-x", x);
        button.style.setProperty("--coord-y", y);
        });
        button.addEventListener("pointerleave", (event) => {
        button.style.setProperty("--coord-x", 0);
        button.style.setProperty("--coord-y", 0);
        });
});