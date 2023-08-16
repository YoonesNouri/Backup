//2.11
'use strict';

//*********************************************

//1- normal function (reusable)
function greet(par1, par2) {
    console.log(`${par1} ${par2}`);
}

greet('I am par1', 'I am par2'); // Output: I am par1 I am par2


//arguments object is a built-in object that returns arguments in an array when passed in a function and invoked.
function function1() {
    console.log(arguments);
}
function1('I am arg1', 'I am arg2'); // Output: Arguments(2) ['I am arg1', 'I am arg2', callee: (...), Symbol(Symbol.iterator): ƒ]

//Accessing the elements of the object "argument" in an array
function function2() {
    console.log(`Hello, ${arguments[0]} and ${arguments[1]}!`);
}
function2('John', 'Jane'); // Output: Hello, John and Jane!


//*********************************************

//2- parameterless function
function func1() {
    console.log('func1 invoked')
};
//call / invoke
func1();

//*********************************************

//3- anonymous function
const addNumbers = function (a, b) {
    return a + b;
};
// Calling the anonymous function
const result = addNumbers(5, 10);
console.log(result); // Output: 15

//*********************************************

//4- void function
function greet1(name) {
    console.log(`Hello, ${name}!`);
}
// Calling the void function
greet1('John');

//*********************************************

//5- normal function
function greet2(name) {
    console.log(`Hello, ${name}!`);
}
// Calling
greet2('world'); // Output: Hello, world!

//*********************************************

//6- Function Expression فاکشن را به یک متغیر اختصاص دادن
const greet3 = function identifier(name) {
    console.log(`Hello, ${name}!`);
};
//Calling
greet3('nouri'); // Output: Hello, nouri!

//*********************************************

//7- Arrow Function . "=>" replaced with "function".
const greet4 = (name) => {
    console.log(`Hello, ${name}!`);
};
//Calling
greet4('yoon'); // Output: Hello, yoon!

//*********************************************

//8- Immediately Invoked Function Expression (IIFE)
(function () {
    console.log('IIFE executed!');
})();
//Calling is automatic . // Output: IIFE executed!

//*********************************************

//9- Functions can be used as properties of objects.?
const myObject = {
    myFunction: function () {
        console.log("Hello from myFunction!");
    }
};
myObject.myFunction(); // Output: Hello from myFunction!

//*********************************************

//10- Function Constructor
const myFunction = new Function('console.log("Dynamically created function");');
myFunction(); // Output: Dynamically created function

//*********************************************

//11- eval() Function : evaluates a string of code as JavaScript code.
eval('function myFunction() { console.log("Dynamically created function"); }');
myFunction(); // Output: Dynamically created function
