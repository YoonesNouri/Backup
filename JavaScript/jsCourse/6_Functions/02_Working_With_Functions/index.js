//6.2

(function () {
    'use strict';
    // length: Returns the number of named parameters defined in the function's declaration.
    function testFunction(arg1, arg2) {
        return arg1 + arg2;
    }
    console.log(testFunction.length); // output: 2 
    // because the function declaration specifies two named parameters, arg1 and arg2.
}());

(function () {
    'use strict';
    // call(): Calls a function with a specified this value and individual arguments provided as separate arguments.
    // با استفاده از متد کال، میتونیم مشخص کنیم آبجکت «دیس» به چی اشاره میکنه
    let person = { nam: 'yoones' };
    function greet(arg) {
        console.log(`I am ${arg} , Hello, ${this.nam}!`);
    }
    greet.call(person, 'mahdiar'); // output: I am mahdiar , Hello, yoones!
    greet.apply(person, ['mahdiar']) // output: I am mahdiar , Hello, yoones!
    // By calling greet.call(person, 'mahdiar')
    // the greet function is invoked with this set to the person object.با ست کردن کال روی آبجکت پرسون ، مشخص کردیم که دیس به چی اشاره میکنه 
    // The argument 'mahdiar' is passed as the arg parameter of the greet function و همچنین با کال ، آرگیومنت مهدیار رو با پارامتر آرگ ، صریحا به فانکشن تعریف کردیم
}());

(function () {
    'use strict';
    // apply(): Calls a function with a specified this value and an array or array - like object of arguments.
    function sum(a, b, c) {
        return a + b + c;
    }
    const numbers = [1, 2, 3];
    const result = sum.apply(null, numbers);
}());

(function () {
    'use strict';
    // bind(): Creates a new function that, when called, has a specified this value and optional arguments pre - filled.
    function greet(name) {
        console.log(`Hello, ${name}!`);
    }
    const greetJohn = greet.bind(null, 'John');
    greetJohn();
}());

(function () {
    'use strict';
    // toString(): Returns a string representing the source code of the function.
    function greet(name) {
        console.log(`Hello, ${name}!`);
    }
    const greetString = greet.toString();
    console.log(greetString);
}());

(function () {
    'use strict';
    // bind(): Creates a new function that, when called, has its this value bound to the provided object.
    const person = {
        name: 'John',
        greet() {
            console.log(`Hello, ${this.name}!`);
        }
    };
    const greetFunction = person.greet.bind(person);
    greetFunction();
}());

(function () {
    'use strict';
    // arguments: An array - like object that contains the arguments passed to a function.
    function sum() {
        let total = 0;
        for (let i = 0; i < arguments.length; i++) {
            total += arguments[i];
        }
        return total;
    }
    console.log(sum(1, 2, 3)); // Output: 6
}());

(function () {
    'use strict';

}());

(function () {
    'use strict';

}());

