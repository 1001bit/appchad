const updateTime = 4000;
let interval = window.setInterval(chatGet, updateTime);
let lastMessageId = 0

$(document).ready(chatGet)

// fetch api and add messages to page
async function chatGet(){
    try {
        lastMessageId = $(".message").last().attr("id")
    } catch {
        lastMessageId = 0
    }

    return fetch("api/chatchad?id="+lastMessageId, {
        method: "GET"
    })
    .then(response => response.json())
    .then(data => {
        addNewMessages(data)
    })
    .catch(error => {
        console.log(error)
    })
}

// post message to chat
async function chatPost(msgText){
    // try {
    //     lastMessageId = $(".message").last().attr("id")
    // } catch {
    //     lastMessageId = 0
    // }

    return fetch("api/chatchad?text="+msgText, {
        method: "POST"
    })
    // .then(response => response.json())
    // .then(data => {
    //     addNewMessages(data)
    // })
    .catch(error => {
        console.log(error)
    })
}

// add messages to page from data
function addNewMessages(data){
    const chatBox = $("#chat");
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
$(".send").click(async () => {
    const typebox = $(".typebox").last()
    if(typebox.val().length == 0){
        return
    }

    const text = typebox.val()
    typebox.val("")

    window.clearInterval(interval)
    interval = window.setInterval(chatGet, updateTime);

    let start = Date.now()

    let postPromise = chatPost(text)
    postPromise.then(() => {
        chatGet()
        console.log(Date.now() - start)
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