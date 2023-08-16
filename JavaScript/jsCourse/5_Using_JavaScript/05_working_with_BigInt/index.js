//5.5

(function () {
    'use strict';
    // In JavaScript, you can produce big integers(BigInts) by using
    // the BigInt constructor or by appending the n suffix to a numeric literal.
    const bigNum1 = BigInt(123456789); // Using the BigInt constructor
    const bigNum2 = 987654321n; // Using the `n` suffix
    console.log(bigNum1); // Output: 123456789n
    console.log(bigNum2); // Output: 987654321n

    console.log('*******************************************');
    // BigInt to JSON:
    const bigNum3 = BigInt(1234567890);
    const jsonString = JSON.stringify({ num: bigNum3.toString() });
    // Output: '{"num":"1234567890"}'
    const parsedObj = JSON.parse(jsonString);
    console.log(parsedObj); // Output: {num: '1234567890'}
    const parsedBigNum = BigInt(parsedObj.num);
    console.log(parsedBigNum); // Output: 1234567890n

    console.log('*******************************************');
    // BigInt.prototype.toString(): Returns the string representation of the BigInt.
    const bigNum4 = BigInt(1234567890);
    const stringRepresentation = bigNum4.toString();
    console.log(stringRepresentation); // Output: "1234567890"

    console.log('*******************************************');
    // BigInt.prototype.toLocaleString(): Returns a string representation using the locale - specific number format.
    const bigNum5 = BigInt(1234567890);
    const localeString = bigNum5.toLocaleString();
    console.log(localeString); // Output: "1,234,567,890" (based on the locale-specific number format)

    console.log('*******************************************');
    // BigInt.prototype.valueOf(): Returns the primitive value of the BigInt.
    const bigNum6 = BigInt(1234567890);
    const primitiveValue = bigNum6.valueOf();
    console.log(primitiveValue); // Output: 1234567890 (primitive value of the BigInt)
}());