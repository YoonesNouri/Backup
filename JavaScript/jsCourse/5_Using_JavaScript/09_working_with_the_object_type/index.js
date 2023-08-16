//5.9

(function () {
    'use strict';

    const person = {
        name: 'John'
    };
    const prototypeOfPerson = Object.getPrototypeOf(person);
    console.log(prototypeOfPerson === Object.prototype); // Output: true
    Object.prototype.greet = function () {
        console.log(`Hello, my name is ${this.name}.`);
    };
    person.greet(); // Output: Hello, my name is John.

}());

(function () {
    'use strict';
    Array.prototype.customProperty = 'Custom Value';
    const array = [1, 2, 3];
    console.log(array.customProperty); // Output: 'Custom Value'
}());

(function () {
    'use strict';
    // creating a copy of a frozen object and modifying the copy:
    const obj = { name: 'John', age: 30 };
    const frozenObj = Object.freeze(obj); // Freeze the original object
    const newObj = { ...frozenObj }; // Create a shallow copy of the frozen object
    newObj.name = 'Jane'; // Modify the property of the new object
    console.log(obj);       // Output: { name: 'John', age: 30 }
    console.log(newObj);    // Output: { name: 'Jane', age: 30 }
}());

(function () {
    'use strict';
    // Object.isFrozen(): checks whether the object is frozen or not.
    const obj = { name: 'John', age: 30 };
    console.log(Object.isFrozen(obj)); // Output: false
    Object.freeze(obj);
    console.log(Object.isFrozen(obj)); // Output: true
}());

(function () {
    'use strict';
    // Object.seal() method is used to seal an object,
    //     which means it prevents new properties from being
    //     added and marks existing properties as non - configurable.
    //     However, it allows modifying the values of existing properties.
    const obj = { name: 'John', age: 30 };
    console.log(Object.isSealed(obj)); // Output: false
    Object.seal(obj);
    console.log(Object.isSealed(obj)); // Output: true
    obj.age = 31; // Modifying the value of an existing property is allowed
    obj.city = 'New York'; // Adding a new property is not allowed
    console.log(obj); // Output: { name: 'John', age: 31 }
}());

(function () {
    'use strict';
    // Object.preventExtensions(): it cannot be extended with new properties.However,
    // it allows modifying or deleting existing properties.
    const obj = { name: 'John', age: 30 };
    console.log(Object.isExtensible(obj)); // Output: true
    Object.preventExtensions(obj);
    console.log(Object.isExtensible(obj)); // Output: false
    obj.age = 31; // Modifying the value of an existing property is allowed
    obj.city = 'New York'; // Adding a new property is not allowed
    console.log(obj); // Output: { name: 'John', age: 31 }
}());

//این اسکوپ رو آخر گذاشتم تا خروجی بقیه حالت ها در کنسول دیده بشه.
(function () {
    'use strict';
    // Object.freeze() method in JavaScript is used to freeze an object, making it immutable.
    // When an object is frozen, its properties cannot be added, modified, or deleted.
    // Any attempt to modify a frozen object will result in an error or silently fail in strict mode.
    const obj = { name: 'John', age: 30 };
    Object.freeze(obj);
    obj.name = 'Jane'; // Attempt to modify a property => Uncaught TypeError: Cannot assign to read only property 'name' of object '#<Object>'
    delete obj.age;   // Attempt to delete a property
    console.log(obj);
    // Output: { name: 'John', age: 30 }
}());
