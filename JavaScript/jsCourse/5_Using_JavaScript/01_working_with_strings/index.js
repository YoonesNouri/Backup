//5.1

(function () {
    'use strict';
    // length: Returns the length of a string.
    const str = "Hello, World!";
    console.log(str.length); // Output: 13

    // charAt(index): Returns the character at the specified index in a string.
    const str1 = "Hello, World!";
    console.log(str1.charAt(0)); // Output: H
    console.log(str1.charAt(7)); // Output: W

    // substring(startIndex, endIndex): Extracts a portion of a string between the specified start and end indexes(end index is optional).
    // negative indexes = 0
    const str2 = "Hello, World!";
    console.log(str2.substring(0, 5)); // Output: Hello
    console.log(str2.substring(7)); // Output: World!

    // slice(startIndex, endIndex): Extracts a portion of a string between the specified start and end indexes(end index is optional).
    // It also supports negative indexes means counting from the end of the string.
    const str3 = "Hello, World!";
    console.log(str3.slice(0, 5)); // Output: Hello
    console.log(str3.slice(7)); // Output: World!
    console.log(str3.slice(-6)); // Output: World!

    // toLowerCase(): Converts a string to lowercase.
    const str4 = "Hello, World!";
    console.log(str4.toLowerCase()); // Output: hello, world!

    // toUpperCase(): Converts a string to uppercase.
    const str5 = "Hello, World!";
    console.log(str5.toUpperCase()); // Output: HELLO, WORLD!

    // trim(): Removes whitespace from both ends of a string.
    const str6 = "   Hello, World!   ";
    console.log(str6.trim()); // Output: Hello, World!

    // split(separator): Splits a string into an array of substrings based on the specified separator.
    const str7 = "Hello World!";
    console.log(str7.split(" ")); // Output: ["Hello" "World!"]چون کلمات استرینگ ورودی با اسپیس از هم جدا شده اند
    // برای متد تعریف میکنیم که هر فاصله بمعنای یک المنت برای آرایه است که خودش جزو آرایه نیست

    // replace(searchValue, replaceValue): Replaces occurrences of a substring within a string with a new substring.
    const str8 = "Hello, World!";
    console.log(str8.replace("World", "Universe")); // Output: Hello, Universe!

    // indexOf(): finds the first occurrence of a specified substring within a string.
    //میگردد و اولین جایگاهی که اون حرف یا کلمه برای اولین بار واقع شده رو پیدا میکنه و ایندکسشو برمیگردونه
    const str9 = "Hello, World!";
    console.log(str9.indexOf("o")); // Output: 4
    console.log(str9.indexOf("World")); // Output: 7
    console.log(str9.indexOf("foo")); // Output: -1 //وقتی وجود نداره 1- رو برمیگردونه

    // lastIndexOf(): finds the last occurrence of a specified substring within a string.
    //میگردد و اولین جایگاهی که اون حرف یا کلمه برای آخرین بار واقع شده رو پیدا میکنه و ایندکسشو برمیگردونه
    console.log(str9.lastIndexOf("o")); // Output: 8

}());