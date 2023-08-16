//3.1
//An array is a single object which can contain multiple values.

//Array constructor
let myArray = new Array(2);
console.log(myArray); // Output: [undefined, undefined]

//Array literal
let myArray1 = ['value1', 'value2'];
console.log(myArray1); // Output: value1 value2

//last value of an array
let lastValue = myArray1[myArray1.length - 1]
console.log(lastValue); // Output: value2

//طول 100 و از ایندکس 0 تا 98 آندیفایند هستند
let myArray2 = []; // Define myArray2 as an empty arrayاگر این کار را نکنیم ارور آندیفایند میدهد
myArray2[99] = 'something'
console.log(myArray2); // Output: [empty × 99, 'something']
console.log(myArray2[0]); // Output: undefined

