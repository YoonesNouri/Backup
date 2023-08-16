//10.3

// Promise.all():
(function () {
    'use strict';
    const promise1 = new Promise((resolve) => setTimeout(resolve, 1000, 'Hello'));
    const promise2 = new Promise((resolve) => setTimeout(resolve, 2000, 'World'));
    const promise3 = new Promise((resolve) => setTimeout(resolve, 1500, 'from Promises'));

    Promise.all([promise1, promise2, promise3])
        .then((results) => {
            console.log(results); // Output: ['Hello', 'World', 'from Promises']
        })
        .catch((error) => {
            console.error(error);
        });
}());

// Promise.any():
(function () {
    'use strict';
    const promise1 = new Promise((resolve, reject) => setTimeout(reject, 1000, 'Error 1'));
    const promise2 = new Promise((resolve, reject) => setTimeout(reject, 2000, 'Error 2'));
    const promise3 = new Promise((resolve) => setTimeout(resolve, 1500, 'Success'));

    Promise.any([promise1, promise2, promise3])
        .then((result) => {
            console.log(result); // Output: 'Success'
        })
        .catch((errors) => {
            console.error(errors); // Output: AggregateError: All Promises were rejected
        });
}());

// promise structure:
//         .then((result) => {
//             // Code to handle the fulfilled state
//           })
//         .catch((error) => {
//             // Code to handle the rejected state
//         })
//         .finally(() => {
//             // Code that will be executed after the Promise is settled, whether fulfilled or rejected.
//         });

// finally():
(function () {
    'use strict';
const fetchData = new Promise((resolve, reject) => {
    setTimeout(() => {
        // Simulate a successful asynchronous operation
        resolve('Data successfully fetched.');
        // OR to simulate an error
        // reject('Error: Data not available.');
    }, 2000);
});

fetchData
    .then((data) => {
        console.log(data);
    })
    .catch((error) => {
        console.error(error);
    })
    .finally(() => {
        console.log('Cleanup: This will be executed regardless of the Promise result.');
    });
}());