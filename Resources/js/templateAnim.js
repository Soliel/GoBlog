const baseNavBarHighlightPos = "105vw"

$(document).ready(function(){
    registerNavBarAnimation();
});

function registerNavBarAnimation() {
    $('.navBar li').mouseover(mouseOverNavBar);
    $('.navBar li').mouseleave(mouseLeaveNavBar);
}

function mouseOverNavBar(event) {
    let leftPos = event.target.offsetLeft;
    animateNavHighlight(leftPos);
}

function mouseLeaveNavBar() {
    animateNavHighlight(baseNavBarHighlightPos);
}

function animateNavHighlight(animateTo) {
    let navHighlight = $('#navHighlight');
    navHighlight.stop();
    navHighlight.animate({
        'left': animateTo
    }, 150, "linear");
}
