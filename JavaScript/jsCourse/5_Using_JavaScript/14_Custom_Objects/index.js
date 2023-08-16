//5.14

(function () {
    'use strict';
    // Object.create()
    // create a new object with a specified prototype.It allows you to create an object
    // that inherits properties and methods from a parent object or prototype
    const personPrototype = {
        greet() {
            console.log(`Hello, my name is ${this.name}.`);
        }
    };
    const person = Object.create(personPrototype);
    person.name = 'John';
    person.greet(); // Output: "Hello, my name is John."

}());

(function () {
    'use strict';
    // Object.defineProperty()
    // define a new property directly on an object or modify an existing property,
    // allowing you to specify various attributes for the property.

    // value: The value assigned to the property('John' in this case).
    // writable: A boolean indicating whether the property can be changed or reassigned(false in this case).
    // enumerable: A boolean indicating whether the property will be listed during enumeration(true in this case).
    // configurable: A boolean indicating whether the property can be redefined or deleted(true in this case).

    const person = {};
    Object.defineProperty(person, 'name', {
        value: 'John',
        writable: false,
        enumerable: true,
        configurable: true
    });
    console.log(person.name); // Output: 'John'
    person.name = 'Jane'; // This assignment will be ignored (strict mode) or throw an error (non-strict mode)
    console.log(person.name); // Output: 'John'

}());

(function () {
    'use strict';
    // Object.defineProperties() to define multiple properties on an object at once
    const person = {};
    Object.defineProperties(person, {
        name: {
            value: 'John',
            writable: false,
            enumerable: true,
            configurable: true
        },
        age: {
            value: 30,
            writable: true,
            enumerable: true,
            configurable: true
        }
    });
    console.log(person.name); // Output: 'John'
    console.log(person.age); // Output: 30
    person.name = 'Jane'; // This assignment will be ignored (strict mode) or throw an error (non-strict mode)
    person.age = 31; // The age property can be changed since it is writable
    console.log(person.name); // Output: 'John' (unchanged)
    console.log(person.age); // Output: 31 (changed)
}());