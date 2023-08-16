
(function () {
    'use strict';
    // Mouse Events
    const button = document.getElementById("myButton");
    button.addEventListener("mouseover", () => console.log("Mouse over"));
    button.addEventListener("mousedown", () => console.log("Mouse down"));
    button.addEventListener("mouseup", () => console.log("Mouse up"));

    // Keyboard Events
    const input = document.getElementById("myInput");
    input.addEventListener("keydown", () => console.log("Key down"));
    input.addEventListener("keyup", () => console.log("Key up"));

    // Form Events
    const form = document.getElementById("myForm");
    form.addEventListener("submit", (event) => {
        event.preventDefault();
        console.log("Form submitted");
    });

    // Focus Events
    input.addEventListener("focus", () => console.log("Input focused"));
    input.addEventListener("blur", () => console.log("Input blurred"));

    // Window Events
    window.addEventListener("load", () => console.log("Page loaded"));
    window.addEventListener("resize", () => console.log("Window resized"));
    window.addEventListener("scroll", () => console.log("Page scrolled"));

}());