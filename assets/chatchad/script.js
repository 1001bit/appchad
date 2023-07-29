const updateTime = 3000;
const interval = setInterval(updateChat, updateTime);
let lastMessageId = 0

$(document).ready(() => {
    updateChat();
})

function postMessage(msgText){
    fetch("/api/chatchad", {
        method: "POST",

        body: JSON.stringify({
            text: msgText
        }),

        headers: { 
            "Content-type": "application/json"
        }
    })
}

// fetch api and add messages to page
function updateChat(){
    try {
        lastMessageId = $(".message").last().attr("id")
    } catch {
        lastMessageId = 0
    }

    fetch("api/chatchad?id="+lastMessageId)
    .then(response => response.json())
    .then(data => {
        addNewMessages(data)
    })
}

// add messages to page from data
function addNewMessages(data){
    const chatBox = $(".chat").last();
    const doScroll = chatBox.scrollTop() + chatBox.innerHeight() >= chatBox[0].scrollHeight;
    for(let i = 0; i < data.length; i++){
        let message = $("<div></div>").addClass("message").attr("id", data[i]['id']);
        let date = $("<pre></pre>").text(data[i]['date']);
        let text = $("<pre></pre>").html(`${data[i]['user']}: ${data[i]['text']}`);
        message.append(date);
        message.append(text);
        chatBox.append(message);
    }
    // automatically scroll down
    if(doScroll){
        chatBox.scrollTop(chatBox[0].scrollHeight);
    }
}

// submit button
$(".send").click(() => {
    const typebox = $(".typebox")
    if(typebox.val().length == 0){
        return
    }

    postMessage(typebox.val())
    typebox.val("")
    updateChat()
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