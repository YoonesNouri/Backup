//3.2

(function () {
    'use strict';

    //literal object:
    let myObject = {
        property1: 'hello',
        property2: true
    };

    //constructor object
    let obj = new Object();

    //Access دسترسی:
    //Dote notation:
    myObject.property2; // true
    //brackets notation:
    myObject['property2']; // true

    //append:
    myObject.property3 = false;
    myObject['property4'] = 'something';
    console.log(myObject); // {property1: 'hello', property2: true, property3: false, property4: 'something'}

    //exercise:
    function getThingByColor(color) {
        let things = {
            red: 'red thing',
            green: 'green thing',
            blue: 'blue thing'
        };

        return things[color] || 'Sorry, the color is not in the "things" object';
    }

    console.log(getThingByColor('red'));
    console.log(getThingByColor('purple'));
}());
