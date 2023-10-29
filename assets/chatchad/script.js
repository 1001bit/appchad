const chatBox = $("#chat");
const typebox = $("#typebox")

const urlParams = new URLSearchParams(window.location.search);
const messageId = urlParams.get("id")

// add many messages on startup
function addNewMessages(data){
    const fragment = document.createDocumentFragment()

    // make readable messages from data
    for(let i = 0; i < data.length; i++){
        fragment.append(makeNewMessage(data[i])[0]);
    }
    chatBox.append(fragment)

    // scroll
    if ($(`#${messageId}`)[0]){
        distance = $(`#${messageId}`)[0].offsetTop - chatBox[0].offsetTop
        chatBox.scrollTop(distance)
        $(`#${messageId}`).addClass("highlight")
    } else {
        chatBox.scrollTop(chatBox[0].scrollHeight);
    }
}

// add new message from websocket
function addNewMessage(data){
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
    const message = $("<div></div>").addClass("message").attr("id", data['id'])

    const link = $("<a></a>").text(`message id: ${data['id']}`).attr("href", "/chatchad?id="+data["id"]).addClass("link")
    const date = $("<pre></pre>").text(data['date'])
    const user = $("<a></a>").text(`${data['username']}:`).attr("href", "/chad/"+data["userID"]).addClass("author-name").val(data["userID"])
    const reply = $("<button>reply</button>")
    reply.on("click", () => {
        typebox.val(`@${data["userID"]}, ` + typebox.val())
    })
    
    const imgPattern = /\[img\]([^\]]+?)\[\/img\]/g;
    const textContent = data['text'].replace(imgPattern, '<img alt="[img][/img]" src="$1"></img>')
    const text = $("<pre></pre>").html(textContent)
    
    message.append(link)
    message.append(date)
    message.append(user)
    message.append(text)
    message.append(reply)
    return message
}

// post message to chat
function chatPost(msgText){
    socket.send(JSON.stringify({type: "chat", text: msgText}))
}

// submit button
$("#send").click(async () => {
    const text = typebox.val().trim()
    if(!text) return
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