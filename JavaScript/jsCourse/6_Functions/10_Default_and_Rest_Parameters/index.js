//6.10

(function () {
    'use strict';
    // Default Parameters:
    function greet(Defval = 'default value') {
        console.log(Defval);
    }
    greet();  // Output: default value
    greet('non-default');  // Output: non-default
}());

(function () {
    'use strict';
    // Rest Parameters:
    function sum(...numbers) {
        return numbers.reduce((total, num) => total + num, 0);
    }
    console.log(sum(1, 2, 3));  // Output: 6
    console.log(sum(4, 5, 6, 7));  // Output: 22
}());

(function () {
    'use strict';

}());

(function () {
    'use strict';

}());

(function () {
    'use strict';

}());

(function () {
    'use strict';

}());

(function () {
    'use strict';

}());

(function () {
    'use strict';

}());

