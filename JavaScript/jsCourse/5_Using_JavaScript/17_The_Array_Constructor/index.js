//5.17

(function () {
    'use strict';
    // Array.isArray(): to determine whether a given value is an array or not?
    const array = [1, 2, 3];
    const notArray = 'This is not an array';
    console.log(Array.isArray(array)); // Output: true
    console.log(Array.isArray(notArray)); // Output: false
}());

(function () {
    'use strict';
    // Array.from(): creates a new shallow copy of an array - like or iterable object.
    const arrayLike = { length: 3, 0: 'a', 1: 'b', 2: 'c' };
    const iterable = 'hello';
    const newArray1 = Array.from(arrayLike);
    const newArray2 = Array.from(iterable);
    console.log(newArray1); // Output: ['a', 'b', 'c']
    console.log(newArray2); // Output: ['h', 'e', 'l', 'l', 'o']
}());

(function () {
    'use strict';
    // arguments object is a special object available within the scope of a function.
    // It contains an array - like structure 
    // that holds the arguments passed to the function when it is called.
    function sum() {
        const argsArray = Array.from(arguments);
        const total = argsArray.reduce((acc, curr) => acc + curr, 0);
        return total;
    }
    console.log(sum(1, 2, 3)); // Output: 6
    console.log(sum(4, 5, 6, 7)); // Output: 22
}());

(function () {
    'use strict';
    // numArr array is mapped to its corresponding value in the mapObj object.
    let mapObj = {
        1: 'one',
        2: 'two',
        3: 'three'
    };
    let numArr = [1, 2, 3];
    let mappedArray = Array.from(numArr, function (item) {
        return this[item];
    }, mapObj);
    console.log(mappedArray); // output: ['one', 'two', 'three']
}());

