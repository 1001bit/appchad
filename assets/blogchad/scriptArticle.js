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
    
    // light button
    if($("#up").hasClass("up")){
        return
    }
    $("#up").toggleClass("up", true);

    // show text
    upvotes += 1
    $("#up").html("upvote ("+ upvotes + ")");

    // unlight opposite button
    if(!$("#down").hasClass("down")){
        return
    }
    $("#down").toggleClass("down", false);

    // show text
    downvotes -= 1
    $("#down").html("downvote ("+ downvotes + ")");

    vote("up")
});

$("#down").click(function (e) { 
    e.preventDefault();

    // light button
    if($("#down").hasClass("down")){
        return
    }
    $("#down").toggleClass("down", true);

    // show text
    downvotes += 1
    $("#down").html("downvote ("+ downvotes + ")");

    // unlight opposite button
    if(!$("#up").hasClass("up")){
        return
    }
    $("#up").toggleClass("up", false);

    // show text
    upvotes -= 1
    $("#up").html("upvote ("+ upvotes + ")");

    vote("down")
});