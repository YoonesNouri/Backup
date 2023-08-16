//5.8

(function () {
    'use strict';
    // includes() method checks if the specified element is present in the array and returns a boolean value indicating its presence.
    const array = [1, 2, 3, 4, 5];
    const isIncluded = array.includes(3);
    console.log(isIncluded); // Output: true
}());

(function () {
    'use strict';
    // find() method searches for an element in the array based on the provided condition specified in the callback function. It returns the first element that satisfies the condition or undefined if no element is found.
    const array = [
        { id: 1, name: 'Alice' },
        { id: 2, name: 'Bob' },
        { id: 3, name: 'Charlie' }
    ];
    const result = array.find(item => item.id === 2);
    console.log(result); // Output: { id: 2, name: 'Bob' }
}());

(function () {
    'use strict';
    // flat() method to flatten nested arrays into a single - level array.
    const nestedArray = [1, [2, [3, [4, [5]]]]];
    const flattenedArray = nestedArray.flat(2); // flattens the nested arrays 2 times.
    console.log(flattenedArray);
    // Output: [1, 2, 3, [4, [5]]]]]
}());

(function () {
    'use strict';
    const nestedArray = [1, [2, [3, [4, [5]]]]];
    const flattenedArray = nestedArray.flat(Infinity); // flattens the nested arrays completely.
    console.log(flattenedArray);
    // Output: [1, 2, 3, 4, 5]
}());