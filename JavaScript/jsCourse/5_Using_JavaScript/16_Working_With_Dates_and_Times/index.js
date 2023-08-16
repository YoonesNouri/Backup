//5.16

(function () {
    'use strict';
    let now1 = Date();
    console.log(now1); // output: Sun Jul 16 2023 16:15:36 GMT+0330 (Iran Standard Time)
    console.log(typeof now1); // output: string
    console.log('********************************');
    let now2 = new Date();
    console.log(now2); // output: Sun Jul 16 2023 16:15:36 GMT+0330 (Iran Standard Time)
    console.log(typeof now2); // output: object
}());

(function () {
    'use strict';
    // Creating a Date object:
    const currentDate = new Date(); // Current date and time
    const specificDate = new Date(2022, 6, 15); // July 15, 2022
    const dateString = new Date('2022-07-15'); // July 15, 2022 (parsed from a string)
}());

(function () {
    'use strict';
    // Getting different components of a date:
    const date = new Date();
    const year = date.getFullYear(); // Output: 2023
    const month = date.getMonth(); // Output: 6 (July)
    const day = date.getDate(); // Output: 19
    const hours = date.getHours(); // Output: 12
    const minutes = date.getMinutes(); // Output: 30
    const seconds = date.getSeconds(); // Output: 45
    const milliseconds = date.getMilliseconds(); // Output: 123
}());

(function () {
    'use strict';
    // Working with UTC dates and times
    const date = new Date();
    const utcYear = date.getUTCFullYear(); // Output: 2023
    const utcMonth = date.getUTCMonth(); // Output: 6 (July)
    const utcDay = date.getUTCDay(); // Output: 1 (Monday)
    const utcHours = date.getUTCHours(); // Output: 9
    const utcMinutes = date.getUTCMinutes(); // Output: 30
    const utcSeconds = date.getUTCSeconds(); // Output: 45
}());

(function () {
    'use strict';
// the number of milliseconds elapsed since January 1, 1970, 00:00:00 UTC,
// commonly referred to as the Unix timestamp or epoch time.
    const date = new Date(415236534);// یعنی انقد میلی ثانیه بعد از اون تاریخ اپاک / مبدأ 
    console.log(Date); // ƒ Date() { [native code] }
}());

(function () {
    'use strict';
    // Invalid argument or Incorrect format or Out of range number.
    const invalidDate = new Date('This is not a valid date');
    console.log(invalidDate); // Output: Invalid Date
}());

(function () {
    'use strict';
    // setDate()
    const date = new Date(); // Current date and time
    console.log(date); // Output: current date and time
    date.setDate(10); // Set the day of the month to 10
    console.log(date); // Output: the same month and year, but with the day set to 10
}());

// >>> convertors available on date and time <<<

(function () {
    'use strict';
    // toDateString():
    const date = new Date();
    const dateString = date.toDateString();
    console.log(dateString); // Output: "Mon Aug 02 2021"
}());

(function () {
    'use strict';
    // toTimeString():
    const date = new Date();
    const timeString = date.toTimeString();
    console.log(timeString); // Output: "14:30:45 GMT+0300 (Eastern European Summer Time)"
}());

(function () {
    'use strict';
    // toLocaleDateString():
    const date = new Date();
    const localizedDateString = date.toLocaleDateString();
    console.log(localizedDateString); // Output: "8/2/2021" (format can vary based on the locale)
}());

(function () {
    'use strict';
    // toLocaleTimeString():
    const date = new Date();
    const localizedTimeString = date.toLocaleTimeString();
    console.log(localizedTimeString); // Output: "2:30:45 PM" (format can vary based on the locale)
}());

(function () {
    'use strict';
    // toISOString():
    const date = new Date();
    const isoString = date.toISOString();
    console.log(isoString); // Output: "2021-08-02T11:30:45.000Z"
}());

(function () {
    'use strict';
    // toUTCString():
    const date = new Date();
    const utcString = date.toUTCString();
    console.log(utcString); // Output: "Mon, 02 Aug 2021 11:30:45 GMT"
}());

(function () {
    'use strict';
    // Setting different components of a date:
    const date = new Date();
    date.setFullYear(2023);
    date.setMonth(11); // December (0-11)
    date.setDate(25);
    date.setHours(12);
    date.setMinutes(30);
    date.setSeconds(0);
    date.setMilliseconds(0);
    console.log(date); // Output: Sun Dec 25 2023 12:30:00 GMT+0300 (Eastern European Summer Time)
}());

