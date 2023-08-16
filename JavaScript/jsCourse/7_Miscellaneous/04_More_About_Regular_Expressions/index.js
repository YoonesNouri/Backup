//7.4

(function () {
    'use strict';
    // pipe symbol (|) Matching Either of Two Patterns:

    const regex = /apple|banana/;
    console.log(regex.test('apple'));     // Output: true
    console.log(regex.test('banana'));    // Output: true
    console.log(regex.test('orange'));    // Output: false
    console.log(regex.test('grapefruit')); // Output: false

    const regex1 = /(apple|banana) juice/;
    console.log(regex1.test('apple juice'));    // Output: true
    console.log(regex1.test('banana juice'));   // Output: true
    console.log(regex1.test('orange juice'));   // Output: false
    console.log(regex1.test('grapefruit juice')); // Output: false

    const regex2 = /red|green|blue/;
    console.log(regex2.test('red car'));      // Output: true
    console.log(regex2.test('green tree'));   // Output: true
    console.log(regex2.test('blue sky'));     // Output: true
    console.log(regex2.test('yellow sun'));   // Output: false
}());

(function () {
    'use strict';
    // Logical AND: (?=...) and(?!...)

    const regex1 = /apple(?= juice)/;
    console.log(regex1.test('apple juice'));      // Output: true
    console.log(regex1.test('apple pie'));        // Output: false

    const regex2 = /apple(?! pie)/;
    console.log(regex2.test('apple juice'));      // Output: true
    console.log(regex2.test('apple pie'));        // Output: false
}());

(function () {
    'use strict';
    // Logical NOT: (? !...)
    const regex = /apple(?! juice)/;
    console.log(regex.test('apple pie'));        // Output: true
    console.log(regex.test('apple juice'));      // Output: false
    console.log(regex.test('banana'));           // Output: true
}());

(function () {
    'use strict';
    // /^http(s?)/
    const regex = /^http(s?)/;
    console.log(regex.test('http://www.example.com'));     // Output: true
    console.log(regex.test('https://www.example.com'));    // Output: true
    console.log(regex.test('htt://www.example.com'));      // Output: false (doesn't start with "http" or "https")
    console.log(regex.test('ftp://www.example.com'));      // Output: false (doesn't start with "http" or "https")
    console.log(regex.test('httpsa://www.example.com'));   // Output: false (starts with "httpsa" which is not "http" or "https")
}());

(function () {
    'use strict';
    // global flag(g) match multiple occurrences:
    const regex = /a/g;
    const str = 'abcabc';
    console.log(str.match(regex));  // Output: ['a', 'a']
}());

(function () {
    'use strict';
    // global flag with methods:
    const regex = /a/g;
    const str = 'abcabc';
    console.log(regex.test(str));           // Output: true (matches first occurrence)
    console.log(regex.test(str));           // Output: true (matches second occurrence)
    console.log(str.match(regex));          // Output: ['a', 'a'] (returns all matches as an array)
}());

(function () {
    'use strict';
    let camelCased = 'camelCaseString'
    console.log(camelCased.replace(/([A-Z])/g, '_$1')) // Output: camel_Case_String
}());

(function () {
    'use strict';


}());