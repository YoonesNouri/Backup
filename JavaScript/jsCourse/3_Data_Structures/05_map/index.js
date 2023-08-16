//3.4

(function () {
    'use strict';

    let myMap = new Map();

    //2 dimensional array ولیوی مپ آرایه دو بعدی است
    let twoDArray = [['k1', 'v1'], ['k2', 'v2']]
    let myMap2 = new Map(twoDArray);
    console.log(myMap2); //output: Map(2) {'k1' => 'v1', 'k2' => 'v2'}

    //add k-v by set method
    myMap2.set('k3', 'v3');
    console.log(myMap2); //output: Map(3) {'k1' => 'v1', 'k2' => 'v2', 'k3' => 'v3'}
    console.log(myMap2.get('k3')); //output: v3

    // Deleting a specific key-value pair
    myMap2.delete('k2');
    console.log(myMap2); //output: Map(2) {'k1' => 'v1', 'k3' => 'v3'}
    console.log(myMap2.delete('k4')); //output: false چون کا4 در مپ موجود نیست

    //iterate over a Map by for...of loop
    for (let [key, value] of myMap2) {
        console.log(key, value);
    }

    //iterate over a Map by forEach() method
    myMap2.forEach((value, key) => {
        console.log(key, value);
    });

    //iterate over a Map by entries() method
    for (let [key, value] of myMap2.entries()) {
        console.log(key, value);
    }

    // Deleting all key-value pairs
    myMap2.clear();
    console.log(myMap2); //output: Map(0) {size: 0}

}());
