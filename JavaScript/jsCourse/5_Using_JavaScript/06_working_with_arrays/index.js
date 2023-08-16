//5.6

(function () {
    'use strict';

    // Property:

    // length: Returns the number of elements in the array.
    const array1 = [1, 2, 3, 4, 5];
    console.log(array1.length); // Output: 5

    console.log('***********************************');
    // constructor: Returns a reference to the array's constructor function.
    const array2 = [1, 2, 3];
    console.log(array2.constructor); // Output: Array()

    console.log('***********************************');
    // prototype: Allows you to add properties and methods to all array objects.
    Array.prototype.customMethod = function () {
        console.log("Custom method");
    };
    const array3 = [1, 2, 3];
    array3.customMethod(); // Output: "Custom method"

    console.log('***********************************');
    // Symbol.iterator: Returns the array iterator function, allowing iteration over the array's elements using a for...of loop or the spread operator.
    const array4 = [1, 2, 3];
    const iterator = array4[Symbol.iterator]();
    console.log(iterator.next()); // Output: { value: 1, done: false }

    console.log('***********************************');
    // Symbol.species: Specifies a function used to create derived arrays.It is typically used when subclassing arrays.
    class CustomArray extends Array {
        static get [Symbol.species]() {
            return Array;
        }
    }
    const customArray = new CustomArray(1, 2, 3);
    const mappedArray = customArray.map(num => num * 2);
    console.log(mappedArray instanceof CustomArray); // Output: false
    console.log(mappedArray instanceof Array); // Output: true
}());


//             <<< method >>>


(function () {
    'use strict';

    // Mutating Methods:

    // •	push(element1, element2, ...elementN): Adds one or more elements to the end of an array and returns the new length.
    // •	pop(): Removes the last element from an array and returns that element.
    // •	shift(): Removes the first element from an array and returns that element.
    // •	unshift(element1, element2, ...elementN): Adds one or more elements to the beginning of an array and returns the new length.
    // •	splice(start, deleteCount, item1, item2, ...itemN): Changes the contents of an array by removing, replacing, or adding elements.
    // •	reverse(): Reverses the order of the elements in an array.
    // •	sort(compareFunction): Sorts the elements of an array in place according to the provided compare function or the default sort order.

    const array = [1, 2, 3];

    array.push(4, 5);
    console.log(array); // Output: [1, 2, 3, 4, 5]

    const removedElement = array.pop();
    console.log(removedElement); // Output: 5
    console.log(array); // Output: [1, 2, 3, 4]

    const shiftedElement = array.shift();
    console.log(shiftedElement); // Output: 1
    console.log(array); // Output: [2, 3, 4]

    array.unshift(0);
    console.log(array); // Output: [0, 2, 3, 4]

    array.splice(1, 1, 'a', 'b');
    console.log(array); // Output: [0, 'a', 'b', 3, 4]

    array.reverse();
    console.log(array); // Output: [4, 3, 'b', 'a', 0]

    array.sort();
    console.log(array); // Output: [0, 3, 4, 'a', 'b']

    console.log([10000, 20, 3].sort()); // Output: [10000, 20, 3]
    // compare function
    console.log([10000, 20, 3].sort((a, b) => a - b)); // Output: [3, 20, 10000]

}());

(function () {
    'use strict';

    // Accessing Methods:

    // •	concat(array1, array2, ...arrayN): Returns a new array that combines the elements of the original array with other arrays or values.
    // •	slice(start, end): Returns a shallow copy of a portion of an array into a new array.
    // •	join(separator): Joins all elements of an array into a string, separated by the specified separator.
    // •	toString(): Returns a string representing the elements of an array.

    const array1 = [1, 2, 3];
    const array2 = [4, 5];

    const newArray = array1.concat(array2);
    console.log(newArray); // Output: [1, 2, 3, 4, 5]

    const slicedArray = newArray.slice(2, 4);
    console.log(slicedArray); // Output: [3, 4]

    const joinedString = newArray.join('-');
    console.log(joinedString); // Output: "1-2-3-4-5"

    const arrayToString = newArray.toString();
    console.log(arrayToString); // Output: "1,2,3,4,5"

}());

(function () {
    'use strict';

    // Iteration Methods:

    // •	forEach(callbackFn, thisArg): Executes a provided function once for each array element.
    // •	map(callbackFn, thisArg): Creates a new array with the results of calling a provided function on every element in the array.
    // •	filter(callbackFn, thisArg): Creates a new array with all elements that pass the test implemented by the provided function.
    // •	reduce(callbackFn, initialValue): Applies a function against an accumulator and each element in the array, resulting in a single value.
    // •	find(callbackFn, thisArg): Returns the first element in the array that satisfies the provided testing function.
    // •	some(callbackFn, thisArg): Tests whether at least one element in the array passes the provided testing function.
    // •	every(callbackFn, thisArg): Tests whether all elements in the array pass the provided testing function.

    const array = [1, 2, 3, 4, 5];

    array.forEach((element) => {
        console.log(element); // Output: 1, 2, 3, 4, 5
    });

    const doubledArray = array.map((element) => element * 2);
    console.log(doubledArray); // Output: [2, 4, 6, 8, 10]

    const evenNumbers = array.filter((element) => element % 2 === 0);
    console.log(evenNumbers); // Output: [2, 4]

    const sum = array.reduce((accumulator, element) => accumulator + element, 0);
    console.log(sum); // Output: 15

    const foundElement = array.find((element) => element > 3);
    console.log(foundElement); // Output: 4

    const hasEvenNumber = array.some((element) => element % 2 === 0);
    console.log(hasEvenNumber); // Output: true

    const allEvenNumbers = array.every((element) => element % 2 === 0);
    console.log(allEvenNumbers); // Output: false

}());

(function () {
    'use strict';

    // To sort an array of strings containing both uppercase and lowercase letters
    const strings = ["C", "a", "L", "dog", "boy"];
    strings.sort((a, b) => a.toLowerCase().localeCompare(b.toLowerCase()));
    console.log(strings);
    // Output:  ['a', 'boy', 'C', 'dog', 'L']

    // indexOf() 
    const array = [10, 20, 30, 40, 30, 50];
    console.log(array.indexOf(30)); //searchs form index 0
    // Output: 2
    console.log(array.indexOf(30, 3)); //searchs form index 3
    // Output: 4


}());
