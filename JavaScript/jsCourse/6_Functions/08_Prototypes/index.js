//6.8

(function () {
    'use strict';
    let obj = {};
    console.log(obj);
}());

// وقتی آبجکت خالی را در کنسول باز میکنیم چنین زیر مجموعه ای از
// پروتوتایپ ها میاورد که هر کدام هم زیر مجموعه های زیادی دارند
// Object
// [[Prototype]]
// :
// Object
// constructor
// :
// ƒ Object()
// hasOwnProperty
// :
// ƒ hasOwnProperty()
// isPrototypeOf
// :
// ƒ isPrototypeOf()
// propertyIsEnumerable
// :
// ƒ propertyIsEnumerable()
// toLocaleString
// :
// ƒ toLocaleString()
// toString
// :
// ƒ toString()
// valueOf
// :
// ƒ valueOf()
// __defineGetter__
// :
// ƒ __defineGetter__()
// __defineSetter__
// :
// ƒ __defineSetter__()
// __lookupGetter__
// :
// ƒ __lookupGetter__()
// __lookupSetter__
// :
// ƒ __lookupSetter__()
// __proto__
// :
// (...)
// get __proto__
// :
// ƒ __proto__()
// set __proto_

(function () {
    'use strict';
    const person = {
        name: 'John',
        sayHello() {
            console.log(`Hello, my name is ${this.name}`);
        }
    };
    const john = {
        age: 30
    };
    john.__proto__ = person;  // Setting the prototype of john to person
    john.sayHello();  // Output: "Hello, my name is John"
    console.log(john.age);  // Output: 30
}());