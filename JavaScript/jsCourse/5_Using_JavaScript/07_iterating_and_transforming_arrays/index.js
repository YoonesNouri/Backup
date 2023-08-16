//5.7

(function () {
    'use strict';
    // forEach(callbackFn, thisArg): Executes a provided function once for each array element.
    const array = [1, 2, 3, 4, 5];
    array.forEach((element) => {
        console.log(element); // Output: 1, 2, 3, 4, 5
    });
}());

(function () {
    'use strict';
    // some(): checks if there is at least one element that satisfies a specific condition
    const array = [1, 2, 3, 4, 5];
    const hasEvenNumber = array.some((element) => element % 2 === 0);
    console.log(hasEvenNumber); // Output: true
}());

(function () {
    'use strict';
    // every(callbackFn, thisArg): Tests whether all elements in the array pass the provided testing function.
    const array = [1, 2, 3, 4, 5];
    const allEvenNumbers = array.every((element) => element % 2 === 0);
    console.log(allEvenNumbers); // Output: false

}());

(function () {
    'use strict';
    // reduce(callbackFn, initialValue): Applies a function against an accumulator and each element in the array, resulting in a single value.
    const array = [1, 2, 3, 4, 5];
    const sum = array.reduce((accumulator, element) => accumulator + element, 0);
    console.log(sum); // Output: 15
}());

(function () {
    'use strict';
    // map(callbackFn, thisArg): Creates a new array with the results of calling a provided function on every element in the array.
    const array = [1, 2, 3, 4, 5];
    const doubledArray = array.map((element) => element * 2);
    console.log(doubledArray); // Output: [2, 4, 6, 8, 10]
}());