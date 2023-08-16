//7.5

(function () {
    'use strict';
    // positive Lookbehind:
    const str = 'Hello, World!';
    const regex = /(?<=Hello, )World/;

    const match = str.match(regex);
    if (match) {
        const precedingPattern = match[0]; // Extract the preceding pattern
        console.log(precedingPattern); // Output: "Hello, "

        const newStr = str.replace(regex, 'Universe');
        console.log(newStr); // Output: "Hello, Universe!"
    } else {
        console.log('No match found.');
    }
}());

(function () {
    'use strict';
    // negative Lookbehind:
    const str = 'Hello, World!';
    const regex = /(?<!Hello, )World/;

    const match = str.match(regex);
    if (match) {
        const notPrecedingPattern = match[0]; // Extract the matched pattern
        console.log(notPrecedingPattern); // Output: "World"

        const newStr = str.replace(regex, 'Universe');
        console.log(newStr); // Output: "Hello, Universe!"
    } else {
        console.log('No match found.');
    }
}());

(function () {
    'use strict';
    // Unicode Escape Syntax: \u{ codepoint }
    const unicodeString = '\u{1F600}'; // Represents the "grinning face" emoji (U+1F600)
    console.log(unicodeString); // Output: 😄
}());

(function () {
    'use strict';
    // Unicode property escape:
    const regex = /^\p{Uppercase}+$/u;
    console.log(regex.test('HELLO'));          // Output: true
    console.log(regex.test('Hello'));          // Output: false
    console.log(regex.test('123'));            // Output: false
    console.log(regex.test('WORLD'));          // Output: true
    console.log(regex.test('UPPERCASE'));      // Output: true
    console.log(regex.test('UPPERCASE123'));   // Output: false (contains digits)
}());

(function () {
    'use strict';
    // Matching lowercase alphabetic characters:
    const regexLowercase = /\p{Lowercase}/u;
    console.log(regexLowercase.test('a'));   // Output: true
    console.log(regexLowercase.test('A'));   // Output: false
    console.log(regexLowercase.test('1'));   // Output: false
    console.log(regexLowercase.test('!'));   // Output: false


    // Matching alphabetic characters:
    const regexAlphabetic = /\p{Alphabetic}/u;
    console.log(regexAlphabetic.test('a'));   // Output: true
    console.log(regexAlphabetic.test('A'));   // Output: true
    console.log(regexAlphabetic.test('1'));   // Output: false
    console.log(regexAlphabetic.test('!'));   // Output: false
}());

(function () {
    'use strict';
    // Negative Unicode property escape (\P)
    let isEmoji = /^\P{Emoji}+$/u
    console.log(isEmoji.test('😄'));  // Output: false (Contains emoji)
}());
