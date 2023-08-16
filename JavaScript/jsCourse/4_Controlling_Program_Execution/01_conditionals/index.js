//4.1

(function () {
    'use strict';

    //simple condition
    if (1 + 1 === 2) {
        console.log('the condition is true');
    }

    //multi branches
    if (1 === 2) {
        console.log('in if {} will not be logged');
    } else if (2 !== 2) {
        console.log('in else if {} will not be logged');
    } else {
        console.log('in else {} will be logged');
    }

    //truthy condition
    if (2) {
        console.log('logged because 2 is truthy');
    }

    //falsey condition
    if (0) {
        console.log('not logged because 0 is falsey');
    }

    //ternary condition
    // condition ? expression1 : expression2
    let age = 25;
    let message = (age >= 18) ? 'You are an adult' : 'You are a minor';
    console.log(message);

}());