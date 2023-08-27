const upvotes = $("#up").val();
const downvotes = $("#down").val();

$("#up").click(function (e) { 
    e.preventDefault();
    
    $("#up").toggleClass("up", true);
    $("#down").toggleClass("down", false);

    $("#up").html("upvote ("+ (parseInt(upvotes) + 1) + ")");
    $("#down").html("downvote ("+ (parseInt(downvotes)) + ")");
});

$("#down").click(function (e) { 
    e.preventDefault();
    
    $("#up").toggleClass("up", false);
    $("#down").toggleClass("down", true);

    $("#up").html("upvote ("+ (parseInt(upvotes)) + ")");
    $("#down").html("downvote ("+ (parseInt(downvotes) + 1) + ")");
});