let upvotes = parseInt($("#up").val());
let downvotes = parseInt($("#down").val());
const articleID = $(".article").attr("id");

function vote(rate){
    fetch("/api/blogchad/vote", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({articleID: articleID, vote: rate}),
    })
    .catch(error => {
        console.log(error)
    })
}

$("#up").click(function (e) { 
    e.preventDefault();
    
    vote("up")
    upvotes += 1

    // unlight opposite button
    if($("#down").hasClass("down") && !$("#up").hasClass("up")){
        downvotes -= 1
    }

    $("#up").toggleClass("up", true);
    $("#down").toggleClass("down", false);

    $("#up").html("upvote ("+ upvotes + ")");
    $("#down").html("downvote ("+ downvotes + ")");
});

$("#down").click(function (e) { 
    e.preventDefault();

    vote("down")
    downvotes += 1

    // unlight opposite button
    if(!$("#down").hasClass("down") && $("#up").hasClass("up")){
        upvotes -= 1
    }

    $("#up").toggleClass("up", false);
    $("#down").toggleClass("down", true);

    $("#up").html("upvote ("+ upvotes + ")");
    $("#down").html("downvote ("+ downvotes + ")");
});