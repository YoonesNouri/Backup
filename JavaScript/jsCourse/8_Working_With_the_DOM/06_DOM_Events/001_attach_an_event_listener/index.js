//8.6

(function () {
    'use strict';
    // JavaScript code to attach an event listener to a button
    const button = document.getElementById("myButton");
    function handleClick() {
        alert("Button clicked!");
    }
    // Attach the click event listener to the button
    button.addEventListener("click", handleClick);
}());

(function () {
    'use strict';
    // JavaScript code to implement event delegation
    const myList = document.getElementById("myList");
    function handleItemClick(event) {
        const clickedItem = event.target;
        if (clickedItem.tagName === "LI") {
            console.log("Clicked item: ", clickedItem.textContent);
        }
    }
    myList.addEventListener("click", handleItemClick);
}());