//5.12

(function () {
    'use strict';
    // Object.values(): extracts the values of an object's enumerable properties
    // and return them as an array.
    const obj = { a: 1, b: 2, c: 3 };
    const valuesArray = Object.values(obj);
    console.log(valuesArray); // Output: [1, 2, 3]
}());

(function () {
    'use strict';
    // Object.entries()
    // extracts the key - value pairs of an object's enumerable properties
    // and return them as an array of arrays.
    const obj = { a: 1, b: 2, c: 3 };
    const entriesArray = Object.entries(obj);
    console.log(entriesArray);
    // Output: [ [ 'a', 1 ], [ 'b', 2 ], [ 'c', 3 ] ]
}());

(function () {
    'use strict';
    // Object.fromEntries() to create an object from an array of key - value pairs.
    const entriesArray = [['a', 1], ['b', 2], ['c', 3]];
    const obj = Object.fromEntries(entriesArray);
    console.log(obj);
    // Output: { a: 1, b: 2, c: 3 }
}());

(function () {
    'use strict';
    // Object.getOwnPropertyDescriptor() and Object.getOwnPropertyDescriptors()
    // are JavaScript methods used to retrieve property descriptors of an object's own properties.
    // detailed information about a property, including attributes
    // like value, writable, enumerable, and configurable.
    const obj = { a: 1 };
    const descriptor = Object.getOwnPropertyDescriptor(obj, 'a');
    console.log(descriptor);
    // Output: { value: 1, writable: true, enumerable: true, configurable: true }
}());

(function () {
    'use strict';
    // Object.keys() method returns an array containing the keys of the object.
    const obj = { a: 1, b: 2, c: 3 };
    const keys = Object.keys(obj);
    console.log(keys);
    // Output: ['a', 'b', 'c']
}());
