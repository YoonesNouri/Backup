//10.1

(function () {
    'use strict';
    // Example asynchronous function that returns a Promise
    function fetchData() {
        return new Promise((resolve, reject) => {
            // Simulate an asynchronous operation (e.g., fetching data from an API)
            setTimeout(() => {
                const data = { name: 'John', age: 30 };
                if (data) {
                    resolve(data); // Resolve the Promise with the data
                } else {
                    reject('Data not available.'); // Reject the Promise with an error message
                }
            }, 2000); // Simulate a 2-second delay
        });
    }

    // Using the Promise
    fetchData()
        .then((data) => {
            console.log('Data:', data); // This will execute when the Promise is fulfilled
        })
        .catch((error) => {
            console.error('Error:', error); // This will execute when the Promise is rejected
        });

}());