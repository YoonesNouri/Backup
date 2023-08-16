//2.10
//Nullish Coalescing ادغام تهی 
const name = null;
const defaultName = 'John';
const result = name ?? defaultName;

console.log(result); // Output: "John"

//زنجیربندی اختیاری Optional Chaining 
const person = {
    name: 'John',
    age: 30,
    address: {
        city: 'New York',
        zipCode: 12345
    }
};

const zipCode = person?.address?.zipCode;
console.log(zipCode); // Output: 12345

const phoneNumber = person?.contact?.phoneNumber;
console.log(phoneNumber); // Output: undefined
//In the above example, the optional chaining operator (?.) is used
// to access the zipCode property of the address object.
//It works even if any intermediate property(address or contact) is null or undefined.

