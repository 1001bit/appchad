const button = $("#edit-form")[0]

$("#edit-btn").click(function (e) { 
    e.preventDefault();
    button.style.display = button.style.display == "block" ?
    button.style.display = "none" : button.style.display = "block"
});