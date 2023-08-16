//6.6

(function () {
    'use strict';
    // Syntax:
    function* myGeneratorFunction() {
        // Generator function body
    }
}());

(function () {
    'use strict';
    function* countUpTo(limit) {
        for (let i = 1; i <= limit; i++) {
            yield i;
        }
    }
    const generator = countUpTo(5);
    console.log(generator.next());  // Output: { value: 1, done: false }
    console.log(generator.next());  // Output: { value: 2, done: false }
    console.log(generator.next());  // Output: { value: 3, done: false }
    console.log(generator.next());  // Output: { value: 4, done: false }
    console.log(generator.next());  // Output: { value: 5, done: false }
    console.log(generator.next());  // Output: { value: undefined, done: true }
}());

(function () {
    'use strict';

}());


