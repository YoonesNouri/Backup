//6.1

(function () {
    'use strict';
    console.log(this); // output: undefined
}());

(function () {
    // 'use strict';
    console.log(this); // output: Window {window: Window, self: Window, document: document, name: '', location: Location, …} هرگز توصیه نمیشه
}());

(function () {
    'use strict';
    let obj = {
        method() {
            console.log(this); // output: {method: ƒ}
        }
    };
    obj.method();
}());

(function () {
    'use strict';
    function Person(namee) {
        this.namee = namee
    }
    let bob = new Person('I\'m bob')
    console.log(bob);
}());



