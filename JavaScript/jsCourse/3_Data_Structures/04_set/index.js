//3.3

(function () {
    'use strict';

    let mySet = new Set();
    //iterable object
    let myArray = [1, 1, 1, 2, 2, 3];
    let mySet2 = new Set(myArray);

    console.log(mySet2); //output: Set(3) {1, 2, 3}
    console.log(mySet2.size); //output: 3

    mySet2.add(4).add(5); //chainادد زنجیر شده 
    mySet2.delete(1);
    console.log(mySet2); //output: Set(3) {2, 3, 4, 5}
    console.log(mySet2.delete(6)); //output: false چون ولیو 6 در ست وجود نداره
    console.log(mySet2.has(2)); //output: true چون ولیو 2 در ست وجود داره

}());
