const chatBox = $("#chat");
const typebox = $("#typebox")

const urlParams = new URLSearchParams(window.location.search);
const messageId = urlParams.get("id")

// fetch api and add all messages to page on starup
$(document).ready(() => {
    fetch(`/api/chatchad`, {
        method: "GET"
    })
    .then(response => response.json())
    .then(data => {
        if (data.length > 0){
            addNewMessages(data)
        }
    }).then(() => {
        if (messageId == null){
            chatBox.scrollTop(chatBox[0].scrollHeight);
        } else {
            distance = $(`#${messageId}`)[0].offsetTop - chatBox[0].offsetTop
            console.log($(`#${messageId}`)[0].offsetTop, chatBox[0].offsetTop)
            chatBox.scrollTop(distance)
            $(`#${messageId}`).addClass("highlight")
        }
    })
    .catch(error => {
        console.log(error)
    })
})

// add messages to page from data
function addNewMessages(data){
    const fragment = document.createDocumentFragment()

    // make readable messages from data
    for(let i = 0; i < data.length; i++){
        fragment.append(makeNewMessage(data[i])[0]);
    }
    chatBox.append(fragment)
}

// when server sent a message - add the chat message to a wall
socket.onmessage = (event) => {
    data = JSON.parse(event.data)
    if (data.type != "chat") {
        return
    }

    const doScroll = chatBox.scrollTop() + chatBox.innerHeight() >= chatBox[0].scrollHeight-1;
    message = makeNewMessage(data)
    chatBox.append(message)
    // automatically scroll down
    if(doScroll){
        chatBox.scrollTop(chatBox[0].scrollHeight);
    }
}

// create one new message from data
function makeNewMessage(data){
    const message = $("<div></div>").addClass("message").attr("id", data['id']);
    const date = $("<pre></pre>").text(data['date']);
    const user = $("<a></a>").text(`${data['username']}:`).attr("href", "/chad/"+data["userID"]);

    var pattern = /\[img\]([^\]]+?)\[\/img\]/g;
    const text = $("<pre></pre>").html(data['text'].replace(pattern, '<img alt="[img][/img]" src="$1"></img>'))
    
    message.append(date);
    message.append(user)
    message.append(text);
    return message
}

// post message to chat
function chatPost(msgText){
    socket.send(JSON.stringify({type: "chat", text: msgText}))
}

// submit button
$(".send").click(async () => {
    if(typebox.val().length == 0){
        return
    }

    const text = typebox.val()
    typebox.val("")
    
    chatPost(text)
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