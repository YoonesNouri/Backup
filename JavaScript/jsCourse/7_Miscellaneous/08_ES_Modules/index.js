// index.js

// Importing from myModule.js
import { variable, greet } from './myModule.js';
import defaultItem from './myModule.js';

// Output to the console
console.log(variable); // output: 42
console.log(greet('John')); // output: Hello, John!
console.log(defaultItem); // output: This is the default item
