
(function () {
    'use strict';
    const parent = document.getElementById("parent");
    const innerButton = document.getElementById("innerButton");
    function parentClickHandler() {
        console.log("Parent clicked");
    }
    function innerButtonClickHandler(event) {
        console.log("Inner button clicked");
        event.stopPropagation(); // Prevents the click event from propagating to the parent
    }
    parent.addEventListener("click", parentClickHandler);
    innerButton.addEventListener("click", innerButtonClickHandler);
}());