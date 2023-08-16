//4.3

(function () {
    'use strict';

    let fruit = 'apple';

    switch (fruit) {
        case 'apple':
            console.log('It is an apple.'); // output: It is an apple.
            break;
        case 'banana':
            console.log('It is a banana.');
            break;
        case 'orange':
            console.log('It is an orange.');
            break;
        default:
            console.log('It is an unknown fruit.');
            break;
    }

}());