
(function () {
    'use strict';
    const myButton = document.getElementById("myButton");

    function clickHandler() {
        console.log("Button clicked!");
    }

    // Adding the event listener to the button
    myButton.addEventListener("click", clickHandler);

    // After a certain delay, remove the event listener
    setTimeout(() => {
        myButton.removeEventListener("click", clickHandler);
        console.log("Event listener removed.");
    }, 3000); // Remove the event listener after 3 seconds

}());