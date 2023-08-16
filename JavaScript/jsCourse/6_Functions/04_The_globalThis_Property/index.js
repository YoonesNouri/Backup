//6.4

(function () {
    'use strict';
    console.log(globalThis === window); // Output: true (in a browser environment)
    console.log(globalThis === global); // Output: true (in a Node.js environment)
    globalThis.myVariable = 'Hello, globalThis!';
    console.log(window.myVariable); // Output: "Hello, globalThis!" (in a browser environment)
    console.log(global.myVariable); // Output: "Hello, globalThis!" (in a Node.js environment)
}());

(function () {
    'use strict';
    // self keyword refers to the global object,
    // which is different depending on the execution context.
    // In a web browser environment, self refers to the global object window.
    self.addEventListener('message', function (event) {
        const message = event.data;
        console.log(`Received message: ${message}`);
        // Perform other operations
    });
}());

(function () {
    'use strict';
    // The scrollX property represents the current horizontal scroll position of the window or
    // frame.It is available in browser environments
    // and provides the number of pixels the document is currently scrolled horizontally.
    if (typeof globalThis.scrollX === 'number') {
        console.log('scrollX', globalThis.scrollX);
    }
}());