
(function () {
    'use strict';
    const link = document.getElementById("link");
    const myForm = document.getElementById("myForm");
    function linkClickHandler(event) {
        event.preventDefault(); // Prevents the default navigation behavior of the link
        console.log("Link clicked, but default navigation was prevented.");
    }
    function formSubmitHandler(event) {
        event.preventDefault(); // Prevents the default form submission behavior
        console.log("Form submitted, but default form submission was prevented.");
    }
    link.addEventListener("click", linkClickHandler);
    myForm.addEventListener("submit", formSubmitHandler);
}());