//6.3

(function () {
    'use strict';
    // Global Scope: Variables declared outside any function, or at the top - level scope,
    // have global scope.These variables can be accessed from anywhere in your code,
    // including inside functions.
    const globalVariable = 'I am a global variable';
    function globalFunction() {
        console.log(globalVariable);
    }
    globalFunction(); // Output: "I am a global variable"
}());

(function () {
    'use strict';
    // Local Scope: Variables declared inside a function have local scope.
    // They are accessible only within the function or block in which they are defined.
    function localFunction() {
        const localVariable = 'I am a local variable';
        console.log(localVariable);
    }
    localFunction(); // Output: "I am a local variable"
    // console.log(localVariable); // Error: "localVariable is not defined"
}());

(function () {
    'use strict';
    // Block scope refers to the visibility and accessibility of variables within 
    // a specific block of code, which is typically defined by a pair of curly braces { }.
    // Variables declared using let and const keywords are block-scoped, meaning they are
    // limited to the nearest enclosing block.
    function example() {
        if (true) {
            let blockVariable = 'I am a block variable';
            const blockConstant = 'I am a block constant';
            console.log(blockVariable); // Output: "I am a block variable"
            console.log(blockConstant); // Output: "I am a block constant"
        }
        console.log(blockVariable); // Error: "blockVariable is not defined"
        console.log(blockConstant); // Error: "blockConstant is not defined"
    }
    example();
}());