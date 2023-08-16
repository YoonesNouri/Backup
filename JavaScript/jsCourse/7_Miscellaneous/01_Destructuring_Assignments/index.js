//7.1

(function () {
    'use strict';
    // Array Destructuring:
    const numbers = [1, 2, 3];
    const [a, b, c] = numbers;
    console.log(a);  // Output: 1
    console.log(b);  // Output: 2
    console.log(c);  // Output: 3
}());

(function () {
    'use strict';
    // Object Destructuring:
    const person = { name: 'John', age: 30 };
    const { name, age } = person;
    console.log(name);  // Output: "John"
    console.log(age);   // Output: 30
}());

(function () {
    'use strict';
    // Default Values and Renaming:
    const person = { name: 'John' };
    const { name: personName, age = 30 } = person;
    console.log(personName);  // Output: "John"
    console.log(age);         // Output: 30
}());

(function () {
    'use strict';
// Nested Destructuring:
    const data = { name: 'John', details: { age: 30, city: 'New York' } };
    const { name, details: { age, city } } = data;
    console.log(name);  // Output: "John"
    console.log(age);   // Output: 30
    console.log(city);  // Output: "New York"
}());

(function () {
    'use strict';


}());

