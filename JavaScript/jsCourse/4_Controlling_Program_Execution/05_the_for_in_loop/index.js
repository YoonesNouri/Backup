//4.5

(function () {
    'use strict';

    // for...in loop: Iterates over the enumerable properties of an object.
    let obj = { k1: 'v1', k2: 'v2', k3: 'v3' };
    for (let pr in obj) {
        if (obj.hasOwnProperty(pr)) { // این رو اضافه کرده تا مطمئن بشه که کلیدها برای 
            // آبجکت خودش هست و از آبجکت دیگه ای به ارث نبرده باشه
            console.log(pr, obj[pr]);
        }
    }
    //output: 
    // k1 v1
    // k2 v2
    // k3 v3

}());