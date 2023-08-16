//6.9

(function () {
    'use strict';
    // Spreading Arrays
    const arr1 = [1, 2, 3];
    const arr2 = [...arr1, 4, 5, 6];
    console.log(arr2);  // Output: [1, 2, 3, 4, 5, 6]
}());

(function () {
    'use strict';
    // Spreading Strings:
    const str = 'Hello';
    const chars = [...str];
    console.log(chars);  // Output: ['H', 'e', 'l', 'l', 'o']
}());

(function () {
    'use strict';
    // Spreading Objects:
    const obj1 = { x: 1, y: 2 };
    const obj2 = { ...obj1, z: 3 };
    console.log(obj2);  // Output: { x: 1, y: 2, z: 3 }
}());

(function () {
    'use strict';
    // Spreading Function Arguments:
    function sum(a, b, c) {
        return a + b + c;
    }
    const numbers = [1, 2, 3];
    console.log(sum(...numbers));  // Output: 6
}());

(function () {
    'use strict';
    // Combining Spreads:
    const arr1 = [1, 2];
    const arr2 = [3, 4];
    const combined = [...arr1, ...arr2, 5, ...[6, 7]];
    console.log(combined);  // Output: [1, 2, 3, 4, 5, 6, 7]
}());

(function () {
    'use strict';


}());

(function () {
    'use strict';


}());

