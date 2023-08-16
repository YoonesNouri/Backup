//6.7

(function () {
    'use strict';
    function outer() {
        const outerVar = 'I am from the outer function';

        function inner() {
            console.log(outerVar);  // Accessing outerVar from the enclosing scope
        }

        return inner;
    }
    const closureFn = outer();  // Assigning the inner function to closureFn
    closureFn();  // Output: "I am from the outer function"
}());

// In this example, the inner function forms a closure over the outerVar variable.
// Even after the outer function has finished executing, closureFn still has access
// to outerVar and can log its value when invoked.