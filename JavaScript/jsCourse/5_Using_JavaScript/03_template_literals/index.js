//5.3

(function () {
    'use strict';
    // وقتی از backticks `` بجای '' یا ""  استفاده بشه ، با استفاده از ${ } این ویژگیها رو ساپورت میکنه:

    // 1 - جاسازی کلپمه String Interpolation:
    const name = "Alice";
    console.log(`Hello, ${name}!`); // Output: "Hello, Alice!"

    // 2 - استرینگ چند خطی Multiline Strings:
    const message = `This is a
multiline
string.`;
    console.log(message);
    /*
    Output:
    This is a
    multiline
    string.
    */

    // 3 - ارزیابی عبارت Expression Evaluation:
    const a = 5;
    const b = 10;
    console.log(`The sum of ${a} and ${b} is ${a + b}.`); // Output: "The sum of 5 and 10 is 15."

    // 4 - از کارانداختن خاصیت Escaping Characters:
    // با نوشتن \ backslash قبل از ` backtick ، خاصیت بک¬تیک که نقل قول کردن است از بین میرود و خودش پرینت میشود.
    const message1 = `This is a backtick: \` and a newline: \n`;
    console.log(message1);
    /*
        Output:
        This is a backtick: ` and a newline:
        */

}());