//10.4

// basic syntax of using the async keyword:
(function () {
    'use strict';
    async function functionName() {
        // Asynchronous operations using await keyword
        const result = await someAsyncFunction();
        return result;
    }
}());

// with arrow functions:
(function () {
    'use strict';
    const functionName = async () => {
        // Asynchronous operations using await keyword
        const result = await someAsyncFunction();
        return result;
    };
}());

// async keyword with a simple asynchronous function:
(function () {
    'use strict';
    function fetchData() {
        return new Promise((resolve) => {
            setTimeout(() => {
                resolve('Data fetched successfully.');
            }, 2000);
        });
    }

    async function getData() {
        try {
            const result = await fetchData();
            console.log(result); // Output: 'Data fetched successfully.'
            return result;
        } catch (error) {
            console.error(error);
        }
    }

    getData();

}());

(function () {
    'use strict';

}());

