//5.2

(function () {
    'use strict';
    // 1.padStart(targetLength, padString): Pads the beginning of the string with the specified padString until the resulting string reaches the targetLength.
    const str = "Hello";
    console.log(str.padStart(10, "*")); // Output: "*****Hello"

    // 2.padEnd(targetLength, padString): Pads the end of the string with the specified padString until the resulting string reaches the targetLength.
    const str1 = "Hello";
    console.log(str1.padEnd(10, "*")); // Output: "Hello*****"

    // trim و replace جلسه قبلی اومدن.
}());