//4.4

(function () {
    'use strict';

    // 1 - for loop: Executes a block of code for a specified number of times.
    console.log('normal loop');
    for (let i = 0; i < 5; i++) {
        console.log(i);
    }

    console.log('****************************************************');

    console.log('manual loop (better to be avoided)');
    let x = 0;
    for (; ;) {
        if (x >= 5) {
            break;
        }
        console.log(x);
        x += 1;
    }


}());