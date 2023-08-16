//7.3

(function () {
    'use strict';

}());

(function () {
    'use strict';
    // 1. Creating Regular Expressions:
    // Regular expressions are created using the RegExp constructor or 
    // by using a regex literal notation(enclosed in forward slashes /.../).
    const regex1 = new RegExp('pattern');
    const regex2 = /pattern/;
}());

(function () {
    'use strict';
    // 2. Matching Patterns:
    const regex = /abc/;
    console.log(regex.test('abcdef'));  // Output: true
    console.log('abcdef'.match(regex));  // Output: ['abc']
}());

(function () {
    'use strict';
    // 3. Pattern Modifiers and Flags:
    const regex = /abc/gi;
    console.log('ABCabcdef'.match(regex));  // Output: ['ABC', 'abc']
}());

(function () {
    'use strict';
    // 1. Character Classes:
    const regex1 = /[abc]/;
    console.log(regex1.test('a'));  // Output: true
    console.log(regex1.test('d'));  // Output: false

    const regex2 = /[a-z]/;
    console.log(regex2.test('h'));  // Output: true
    console.log(regex2.test('5'));  // Output: false

    const regex3 = /[^0-9]/;
    console.log(regex3.test('7'));  // Output: false
    console.log(regex3.test('x'));  // Output: true

    // 2. Shorthand Character Classes:
    const regex4 = /\d/;
    console.log(regex4.test('5'));  // Output: true
    console.log(regex4.test('x'));  // Output: false

    const regex5 = /\w/;
    console.log(regex5.test('A'));  // Output: true
    console.log(regex5.test('#'));  // Output: false

    const regex6 = /\s/;
    console.log(regex6.test(' '));  // Output: true
    console.log(regex6.test('x'));  // Output: false

    // 3. Quantifiers:
    const regex7 = /a*/;
    console.log(regex7.test(''));    // Output: true
    console.log(regex7.test('a'));   // Output: true
    console.log(regex7.test('aaa')); // Output: true
    console.log(regex7.test('b'));   // Output: false

    const regex8 = /b+/;
    console.log(regex8.test(''));    // Output: false
    console.log(regex8.test('b'));   // Output: true
    console.log(regex8.test('bbb')); // Output: true
    console.log(regex8.test('a'));   // Output: false

    const regex9 = /c?/;
    console.log(regex9.test(''));    // Output: true
    console.log(regex9.test('c'));   // Output: true
    console.log(regex9.test('cc'));  // Output: false

    const regex10 = /d{3}/;
    console.log(regex10.test('d'));   // Output: false
    console.log(regex10.test('ddd')); // Output: true
    console.log(regex10.test('dddd')); // Output: true

    const regex11 = /e{2,5}/;
    console.log(regex11.test('e'));    // Output: false
    console.log(regex11.test('ee'));   // Output: true
    console.log(regex11.test('eeee')); // Output: true
    console.log(regex11.test('eeeee')); // Output: true
    console.log(regex11.test('eeeeee')); // Output: true
    console.log(regex11.test('eeeeeee')); // Output: false

}());

(function () {
    'use strict';
    // Metacharacters and Escape Sequences:

    // 1. Metacharacters:
    const regex1 = /\./;
    console.log(regex1.test('hello.'));  // Output: true
    console.log(regex1.test('hello'));   // Output: false

    const regex2 = /\d+/;
    console.log(regex2.test('123'));     // Output: true
    console.log(regex2.test('abc'));     // Output: false

    const regex3 = /\w+/;
    console.log(regex3.test('hello'));    // Output: true
    console.log(regex3.test('123'));      // Output: true
    console.log(regex3.test('#$%'));     // Output: false

    const regex4 = /\s\w+/;
    console.log(regex4.test(' Hello'));   // Output: true
    console.log(regex4.test('World'));    // Output: false

    // 2. Escape Sequences:
    const regex5 = /\d\+/;
    console.log(regex5.test('123+'));    // Output: true
    console.log(regex5.test('123'));     // Output: false

    const regex6 = /\w\+/;
    console.log(regex6.test('hello+'));  // Output: true
    console.log(regex6.test('hello'));   // Output: false

    const regex7 = /[\[\]]/;
    console.log(regex7.test('['));       // Output: true
    console.log(regex7.test(']'));       // Output: true
    console.log(regex7.test('abc'));     // Output: false

    const regex8 = /a\/b/;
    console.log(regex8.test('a/b'));     // Output: true
    console.log(regex8.test('a'));       // Output: false

}());

(function () {
    'use strict';
    // dollar sign($):

    // 1. Anchoring at the End of the String:
    const regex1 = /abc$/;
    console.log(regex1.test('abcdef'));    // Output: false
    console.log(regex1.test('xyzabc'));    // Output: true
    console.log(regex1.test('abc'));       // Output: true

    // 2. Matching the End of a Line in Multiline Mode:
    const regex2 = /^abc$/m;
    console.log(regex2.test('abc\ndef'));  // Output: true
    console.log(regex2.test('xyz\nabc'));  // Output: false
    console.log(regex2.test('abc'));       // Output: true

    // 3. Capturing Groups:
    const regex3 = /(\d+)-\1/;
    console.log(regex3.test('123-123'));   // Output: true
    console.log(regex3.test('456-789'));   // Output: false

}());

