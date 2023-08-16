//7.6

(function () {
    'use strict';
    class Person {
        constructor(name, age) {
            this.name = name;
            this.age = age;
        }
        greet() {
            console.log(`Hello, my name is ${this.name} and I am ${this.age} years old.`);
        }
    }
    const person1 = new Person('John', 30);
    person1.greet();  // Output: "Hello, my name is John and I am 30 years old."
}());

(function () {
    'use strict';
    // Extend(Inheritance):
    class Animal {
        constructor(name) {
            this.name = name;
        }

        speak() {
            console.log(`${this.name} makes a sound.`);
        }
    }
    class Dog extends Animal {
        speak() {
            console.log(`${this.name} barks.`);
        }
    }
    const dog = new Dog('Buddy');
    dog.speak();  // Output: "Buddy barks."
}());

(function () {
    'use strict';
    // Static:
    class MathUtils {
        static add(x, y) {
            return x + y;
        }
    }
    console.log(MathUtils.add(3, 5));  // Output: 8
}());

(function () {
    'use strict';
    // super keyword :
    // Calling the Parent Class Constructor:
    class Animal {
        constructor(name) {
            this.name = name;
        }
    }
    class Dog extends Animal {
        constructor(name, breed) {
            super(name); // Call the parent class constructor
            this.breed = breed;
        }
    }
    const dog = new Dog('Buddy', 'Labrador');
    console.log(dog.name);  // Output: "Buddy"
    console.log(dog.breed); // Output: "Labrador"
}());

(function () {
    'use strict';
    // super keyword :
    // Accessing Parent Class Methods:
    class Animal {
        speak() {
            console.log('Animal speaks.');
        }
    }
    class Dog extends Animal {
        speak() {
            super.speak(); // Call the parent class method
            console.log('Dog barks.');
        }
    }
    const dog = new Dog();
    dog.speak();
}());
