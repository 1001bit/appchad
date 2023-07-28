const updateTime = 3000;
const interval = setInterval(updateChat, updateTime);
let lastMessageId = 0

function postMessage(){
    // TODO
}

// fetch api and add messages to page
function updateChat(){
    lastMessageId = $(".message").last().attr("id")

    fetch("chatchad/chat?id="+lastMessageId)
    .then(response => response.json())
    .then(data => {
        addNewMessages(data)
    })
}

// fetch database and add new messages to box
function addNewMessages(data){
    let chatBox = $(".chat").last();
    let doScroll = chatBox.scrollTop() + chatBox.innerHeight() >= chatBox[0].scrollHeight;
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

/////////////////////
// submit button
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
  
  