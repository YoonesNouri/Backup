//5.13

(function () {
    'use strict';
    const person = {
        firstName: 'John',
        lastName: 'Doe',

        get fullName() {
            return `${this.firstName} ${this.lastName}`;
        },

        set fullName(value) {
            const [firstName, lastName] = value.split(' ');
            this.firstName = firstName;
            this.lastName = lastName;
        }
    };

    console.log(person.fullName); // Output: "John Doe"

    person.fullName = 'Jane Smith';
    console.log(person.firstName); // Output: "Jane"
    console.log(person.lastName);  // Output: "Smith"
    console.log(person.fullName);  // Output: "Jane Smith"

}());