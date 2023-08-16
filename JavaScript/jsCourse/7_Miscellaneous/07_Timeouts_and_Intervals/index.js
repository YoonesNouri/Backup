//7.7

(function () {
    'use strict';
    // setTimeout():
    setTimeout(function () {
        console.log("Delayed message");
    }, 2000);
}());

(function () {
    'use strict';
    // clearTimeout():
    const timeoutId = setTimeout(function () {
        console.log("Delayed message");
    }, 2000);
    clearTimeout(timeoutId); // Cancel the timeout
}());

(function () {
    'use strict';
    // setInterval():
    setInterval(function () {
        console.log("Repeating message");
    }, 1000);
}());

(function () {
    'use strict';
    // clearInterval():
    const intervalId = setInterval(function () {
        console.log("Repeating message");
    }, 1000);
    clearInterval(intervalId); // Stop the interval
}());

(function () {
    'use strict';
    // counting:
    let count = 0;
    const intervalId = setInterval(function () {
        console.log(count);
        count++;
        if (count > 5) {
            clearInterval(intervalId); // Clear the interval
        }
    }, 1000);
}());

