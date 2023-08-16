//4.7

(function () {
    'use strict';

    const myiterable = [1, 2, 3];
    const myiterator = myiterable[Symbol.iterator]();

    console.log(myiterator.next()); // { value: 1, done: false }
    console.log(myiterator.next()); // { value: 2, done: false }
    console.log(myiterator.next()); // { value: 3, done: false }
    console.log(myiterator.next()); // { value: undefined, done: true }


}());