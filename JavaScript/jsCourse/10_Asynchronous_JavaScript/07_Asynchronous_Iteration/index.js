//10.7

(function () {
    'use strict';
    // Asynchronous data producer (asynchronous iterator)
    async function* asyncDataGenerator() {
        yield await new Promise((resolve) => setTimeout(resolve, 1000, 'Data 1'));
        yield await new Promise((resolve) => setTimeout(resolve, 1500, 'Data 2'));
        yield await new Promise((resolve) => setTimeout(resolve, 2000, 'Data 3'));
    }

    // Asynchronous consumer using for-await-of loop
    async function consumeAsyncData() {
        for await (const data of asyncDataGenerator()) {
            console.log(data);
        }
    }

    consumeAsyncData();

    // output:
    // Data 1(after 1 second)
    // Data 2(after 1.5 seconds)
    // Data 3(after 2 seconds)

}());

(function () {
    'use strict';
    const asyncDataGenerator = {
        [Symbol.asyncIterator]: async function* () {
            yield await new Promise((resolve) => setTimeout(resolve, 1000, 'Data 1'));
            yield await new Promise((resolve) => setTimeout(resolve, 1500, 'Data 2'));
            yield await new Promise((resolve) => setTimeout(resolve, 2000, 'Data 3'));
        }
    };

    // Asynchronous consumer using for-await-of loop
    async function consumeAsyncData() {
        for await (const data of asyncDataGenerator) {
            console.log(data);
        }
    }

    consumeAsyncData();

    // output:
    // Data 1(after 1 second)
    // Data 2(after 1.5 seconds)
    // Data 3(after 2 seconds)

}());