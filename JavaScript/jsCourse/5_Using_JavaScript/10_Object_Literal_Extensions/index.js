//5.10

(function () {
    'use strict';
    let x = 'test', y = 42;
    // let obj = {
    //     x : x, //یعنی ایکس، کلید باشد و ولیوی ایکس، ولیوی متغیر/وریبل ایکس که تست است باشد
    //     y : y //یعنی وای، کلید باشد و ولیوی وای، ولیوی متغیر/وریبل وای که 42 است باشد
    // };
    // output: {x: 'test', y: 42}

    // equivalent:
    let obj = {
        x, //یعنی ایکس، کلید باشد و ولیوی ایکس، ولیوی متغیر/وریبل ایکس که تست است باشد
        y //یعنی وای، کلید باشد و ولیوی وای، ولیوی متغیر/وریبل وای که 42 است باشد
    };
    console.log(obj);
    // output: {x: 'test', y: 42}
}());

(function () {
    'use strict';
    // Computed Property Names:
    // Property names can be computed dynamically
    // using square brackets[] and an expression enclosed in them.
    const key = 'name';
    const person = {
        [key]: 'John'
    };
    console.log(person); // Output: { name: 'John' }
}());

(function () {
    'use strict';
    // Method Shorthand:
    // When defining a function as a property value, you can use
    //  a shorter syntax by omitting the function keyword and the colon:.
    const person = {
        sayHello() {
            console.log('Hello!');
        }
    };
    person.sayHello(); // Output: Hello!
}());

(function () {
    'use strict';
    // Object Method Definition:
    // You can define object methods using concise method syntax,
    // which omits the function keyword, the colon:, and the comma,.
    const person = {
        sayHello() {
            console.log('Hello!');
        },
        sayGoodbye() {
            console.log('Goodbye!');
        }
    };
    person.sayHello();    // Output: Hello!
    person.sayGoodbye();  // Output: Goodbye!
}());

(function () {
    'use strict';
    // Object.assign():
    // to copy the values of all enumerable properties
    // from one or more source objects to a target object.
    // It allows you to merge multiple objects into a single object.
    const target = { a: 1, b: 2 };
    const source1 = { b: 3, c: 4 };
    const source2 = { b: 5, e: 6 };
    const mergedObject = Object.assign(target, source1, source2);
    // The properties a and b from the target object are preserved.
    // However, the b property is overwritten multiple times by the subsequent sources.
    // The value of b in source2(5) overwrites the value from source1(3).
    console.log(mergedObject);
    // Output: { a: 1, b: 5, c: 4, d: 5, e: 6 }
}());

(function () {
    'use strict';
    // you can use the object spread syntax({ ...source }) for shallow object merging:
    const target = { a: 1, b: 2 };
    const source = { b: 3, c: 4 };
    const mergedObject = { ...target, ...source };
    console.log(mergedObject);
    // Output: { a: 1, b: 3, c: 4 }
}());


