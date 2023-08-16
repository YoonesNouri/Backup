//5.11

(function () {
    'use strict';
    // Object.prototype.toString.call() is used to determine the internal[[Class]] property of an object
    console.log(Object.prototype.toString.call({}));          // Output: [object Object]
    console.log(Object.prototype.toString.call(123));         // Output: [object Number]
    console.log(Object.prototype.toString.call("Hello"));     // Output: [object String]
    console.log(Object.prototype.toString.call(true));        // Output: [object Boolean]
    console.log(Object.prototype.toString.call(null));        // Output: [object Null]
    console.log(Object.prototype.toString.call(undefined));   // Output: [object Undefined]
    console.log(Object.prototype.toString.call(function () { })); // Output: [object Function]
}());

(function () {
    'use strict';
    // new object instance is created based on the constructor function or class definition
    function Person(name, age) {
        this.name = name;
        this.age = age;
    }
    const person1 = new Person('John', 30);
    console.log(person1); // Output: Person { name: 'John', age: 30 }
    const person2 = new Person('Jane', 25);
    console.log(person2); // Output: Person { name: 'Jane', age: 25 }
}());

(function () {
    'use strict';
// you can create object instances using classes, which were introduced in ECMAScript 2015(ES6):
    class Person {
        constructor(name, age) {
            this.name = name;
            this.age = age;
        }
    }
    const person1 = new Person('John', 30);
    console.log(person1); // Output: Person { name: 'John', age: 30 }
    const person2 = new Person('Jane', 25);
    console.log(person2); // Output: Person { name: 'Jane', age: 25 }
}());

(function () {
    'use strict';
    // propertyIsEnumerable(): checks if a specific property exists on an object and is enumerable
    // (i.e., it can be iterated over using a for...in loop).
    const obj = { a: 1, b: 2 };
    console.log(obj.propertyIsEnumerable('a')); // Output: true
    console.log(obj.propertyIsEnumerable('b')); // Output: true
    console.log(obj.propertyIsEnumerable('c')); // Output: false
}());

(function () {
    'use strict';
    // (i.e., it can be iterated over using a for...in loop).
const obj = { a: 1, b: 2 };
for (let prop in obj) {
    if (obj.propertyIsEnumerable(prop)) {
        console.log(prop + ': ' + obj[prop]);
    }
    }
    // Output: 
    // a: 1
    // b: 2
}());