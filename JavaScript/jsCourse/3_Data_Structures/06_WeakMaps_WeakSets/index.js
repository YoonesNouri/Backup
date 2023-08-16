//3.5

// WeakMaps (like Maps) are like Objects.
// WeakMap Example:
(function () {
    'use strict';

    let myWeakMap = new WeakMap();

    let obj1 = {};
    let obj2 = {};

    myWeakMap.set(obj1, 'value 1').set(obj2, 'value 2');

    console.log(myWeakMap); // Output: WeakMap {{…} => 'value 1', {…} => 'value 2'}
    console.log(myWeakMap.get(obj1)); // Output: value 1

    obj1 = null; // Removing the strong reference to obj1

    console.log(myWeakMap.get(obj1)); // Output: undefined, because obj1 has been garbage collected
    console.log(myWeakMap.get(obj2)); // Output: value 2

}());

// WeakSets (like Sets) are like Arrays.
// WeakSet Example:
(function () {
    'use strict';

    let myWeakSet = new WeakSet();

    let obj1 = {};
    let obj2 = {};

    myWeakSet.add(obj1);
    myWeakSet.add(obj2);

    console.log(myWeakSet); // Output: WeakSet {{…}, {…}}
    console.log(myWeakSet.has(obj1)); // Output: true

    obj1 = null; // Removing the strong reference to obj1

    console.log(myWeakSet.has(obj1)); // Output: false, obj1 has been garbage collected
    console.log(myWeakSet.has(obj2)); // Output: true

}());
