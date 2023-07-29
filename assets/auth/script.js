$(document).ready(() => {
    $("#loginForm")[0].addEventListener("submit", (event) => {
        event.preventDefault()

        const inputData = {
            username: $("#username").val(),
            password: $("#password").val()
        }

        fetch("api/auth/"+event.submitter.id, 
        {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(inputData)
        })
        .then(response => response.text())
        .then(data => {
            $("#message").text(data)
            if(data == "success"){
                window.location.replace("/home")
            }
        })
        .catch(error => {
            console.error(error)
        })
    })
})