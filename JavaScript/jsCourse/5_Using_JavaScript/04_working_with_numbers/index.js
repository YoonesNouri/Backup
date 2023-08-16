//5.4

(function () {
    'use strict';
    // isNaN() result:
    //if cannot be converted to a number = true
    //if can be converted to a number = false
    console.log(isNaN(42)); // Output: false
    console.log(isNaN(5 * 'oops')); // Output: true
    console.log(isNaN("Hello")); // Output: true
    console.log(isNaN("123")); // Output: false
    console.log(isNaN(true)); // Output: false
    console.log(isNaN(undefined)); // Output: true

    console.log('*************************************************************');
    // Number.isNaN() : It returns true only if the value is NaN
    console.log(Number.isNaN(42)); // Output: false
    console.log(Number.isNaN("Hello")); // Output: false
    console.log(Number.isNaN(NaN)); // Output: true

    console.log('*************************************************************');
    // Number.isFinite() :  It returns true if the value is a finite number, and false otherwise.
    console.log(Number.isFinite(42)); // Output: true
    console.log(Number.isFinite(Infinity)); // Output: false

    console.log('*************************************************************');
    // Number.isInteger() : It returns true if the value is an integer, and false otherwise.
    console.log(Number.isInteger(42)); // Output: true
    console.log(Number.isInteger(3.14)); // Output: false

    console.log('*************************************************************');
    // Regular expressions:
    // /^\d{4,}$/
    // ^ asserts the start of the string.
    // \d matches any digit character(0 - 9).
    // { 4,} specifies that there should be at least 4 occurrences of the previous element, which is \d.
    // $ asserts the end of the string.
    const regex = /^\d{4,}$/;
    console.log(regex.test("1234")); // Output: true
    console.log(regex.test("56789")); // Output: true
    console.log(regex.test("0123456789")); // Output: true
    console.log(regex.test("123")); // Output: false
    console.log(regex.test("12a34")); // Output: false
    console.log(regex.test("abcd")); // Output: false
    console.log(regex.test("")); // Output: false

    console.log('*************************************************************');
    // Syntax: parseFloat(string)
    // Purpose: Parses a string and returns a floating - point number.
    // Behavior: It extracts and parses the leading numeric portion of the string, ignoring any non - numeric characters after the valid number.
    console.log(parseFloat("3.14")); // Output: 3.14
    console.log(parseFloat("10.5 meters")); // Output: 10.5 (Extracts the number from the string if it is at the beginning of the string)
    console.log(parseFloat("meters10.5")); // Output: NaN (If there is no number at the beginning, it becomes NaN)
    console.log(parseFloat("abc")); // Output: NaN (Not-a-Number)
    console.log('type=', typeof parseFloat("abc")); // Output: type= number

    console.log('*************************************************************');
    // Syntax: parseInt(string, radix)
    // Purpose: Parses a string and returns an integer.
    // Behavior: It extracts and parses the leading numeric portion of the string, optionally considering a specified radix(base).If no radix is provided, it assumes base 10.
    console.log(parseInt("42")); // Output: 42
    console.log(parseInt("1010", 2)); // Output: 10 (binary number)
    console.log(parseInt("12 meter")); // Output: 12 (Extracts the number from the string if it is at the beginning of the string)
    console.log(parseInt("meter12")); // Output: NaN (If there is no number at the beginning, it becomes NaN)
    console.log(parseInt("abc")); // Output: NaN (Not-a-Number)
    console.log('type=', typeof parseInt("abc")); // Output: type= number

    console.log('*************************************************************');
    // toFixed() : rounds the number to the specified number
    const num = 3.14159;
    console.log(num.toFixed()); // Output: "3"
    console.log(num.toFixed(2)); // Output: "3.14"
    console.log(num.toFixed(4)); // Output: "3.1416"

    console.log('*************************************************************');
    // toFixed(digits): Formats a number with a fixed number of decimal places and returns it as a string.
    // toPrecision(precision): Formats a number with a specified precision(total number of significant digits) and returns it as a string.
    // toExponential(fractionDigits): Formats a number in exponential notation(scientific notation) and returns it as a string.
    // toString(radix): Converts a number to a string representation in the specified radix(base).If no radix is provided, it defaults to base 10.
    const numb = 3.14159;
    console.log(numb.toFixed(2)); // Output: "3.14"
    console.log(numb.toPrecision(4)); // Output: "3.142"
    console.log(numb.toExponential(3)); // Output: "3.142e+0"
    console.log(numb.toString(16)); // Output: "3.243f6a8885a3"

    console.log('*************************************************************');
    // JSON.stringify(): converts a JavaScript value, including numbers, to a JSON string representation
    const numbe = 3.14159;
    const jsonString = JSON.stringify(numbe);
    console.log(jsonString); // Output: "3.14159"

    console.log('*************************************************************');
    // Static Methods:
    // Defined on the class itself, rather than on instances of the class.
    // Accessed directly on the class itself, not on individual instances.
    // Do not have access to instance - specific data or behavior.
    // Typically used for utility functions, helper methods, or operations that are not tied to specific instances.
    class MathUtils {
        static multiply(a, b) {
            return a * b;
        }
    }
    console.log(MathUtils.multiply(2, 3)); // Output: 6

    console.log('**********************************************************');
    // Instance Methods:
    // Defined on the prototype of a class and are accessible through instances of the class.
    // Have access to instance - specific data and behavior through the this keyword.
    // Typically used for methods that operate on or interact with specific instances of the class.
    class Circle {
        constructor(radius) {
            this.radius = radius;
        }
        calculateArea() {
            return Math.PI * this.radius * this.radius;
        }
    }
    const circle = new Circle(5);
    console.log(circle.calculateArea()); // Output: 78.53981633974483

}());