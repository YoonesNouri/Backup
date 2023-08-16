//10.8

(function () {
    'use strict';
    document.addEventListener('DOMContentLoaded', () => {
        // Ensure the DOM is ready before executing the JavaScript code

        const button = document.getElementById('loadButton');

        // Verify if the button element was found
        if (button) {
            button.addEventListener('click', async () => {
                try {
                    // Dynamically load the 'utils' module when the button is clicked
                    const utils = await import('./utils.js');
                    utils.logMessage('Module loaded dynamically!');
                } catch (error) {
                    console.error('Error loading module:', error);
                }
            });
        } else {
            console.error("Button element with ID 'loadButton' not found.");
        }
    });
}());


