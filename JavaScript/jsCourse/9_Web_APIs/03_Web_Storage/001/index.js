//9.3

(function () {
    'use strict';
    // index.js
    const textInput = document.getElementById('textInput');
    const saveButton = document.getElementById('saveButton');
    const output = document.getElementById('output');

    // Check if Local Storage is supported by the browser
    if (typeof (Storage) !== "undefined") {
        // Retrieve the saved text from Local Storage if it exists
        if (localStorage.savedText) {
            output.textContent = localStorage.savedText;
        }

        // Save the entered text to Local Storage when the button is clicked
        saveButton.addEventListener('click', () => {
            const textToSave = textInput.value;
            localStorage.savedText = textToSave;
            output.textContent = textToSave;
        });
    } else {
        // Local Storage is not supported
        alert("Sorry, your browser does not support Web Storage.");
    }

}())

