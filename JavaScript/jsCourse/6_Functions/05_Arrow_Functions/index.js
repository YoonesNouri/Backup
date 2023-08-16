//6.5

(function () {
    'use strict';
    // 1.	Arrow Function Syntax:
    const functionName = (parameter1, parameter2) => {
        // Function body
    };
}());

(function () {
    'use strict';
    // Concise Syntax:
    // If the function body consists of a single expression, you can omit the curly braces { }
    // and the return keyword.The result of the expression is automatically returned.
    const double = (num) => num * 2;
}());

(function () {
    'use strict';
    // Lexical this Binding:
    // Arrow functions do not have their own this value.Instead, they inherit this from the surrounding context(lexical scope).
    // This behavior is different from regular functions, which have their own this value determined by how they are called.
    // Arrow functions are commonly used in situations where the this context needs to be preserved, such as in event handlers or when defining methods within objects.
    const person = {
        name: 'John',
        sayHello: function () {
            setTimeout(() => {
                console.log(`Hello, ${this.name}!`);  // 'this' refers to the outer 'person' object
            }, 1000);
        }
    };
    person.sayHello();  // Output: "Hello, John!" after 1 second
}());

(function () {
    'use strict';
    // No arguments Binding:
    // In this example, the arrow function used in the reduce method does not have
    // its own arguments object.Instead, we access the arguments of the 
    // outer function sum to gather the numbers passed as arguments.
    function sum() {
        const numbers = Array.from(arguments);
        return numbers.reduce((total, num) => total + num, 0);
    }
    console.log(sum(1, 2, 3, 4));  // Output: 10
}());

(function () {
    'use strict';
    // Implicit Return:
    // In this example, the arrow function implicitly returns the result of
    // multiplying a and b without using the return keyword explicitly.
    // The expression(a * b) is evaluated and automatically returned.
    const multiply = (a, b) => (a * b);
    console.log(multiply(5, 6));  // Output: 30
}());

(function () {
    'use strict';
    let fn = () => console.log('hi');
    fn(); // output: hi
}());

(function () {
    'use strict';
    // Binding of this:
    // In this example, the arrow function inside the setInterval callback has
    // access to the count property of the obj object through the lexical this
    // binding.It increments the count property every second and logs the updated value.
    const obj = {
        count: 0,
        increment: function () {
            setInterval(() => {
                this.count++;  // 'this' refers to the 'obj' object
                console.log(this.count);
            }, 1000);
        }
    };
    obj.increment();  // Output: Incremented count every second
}());