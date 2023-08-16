//10.6

(function () {
    'use strict';
    // Asynchronous function to simulate fetching data from different sources
    function fetchDataFromSource1() {
        return new Promise((resolve) => {
            setTimeout(() => {
                resolve('Data from source 1');
            }, 1000);
        });
    }

    function fetchDataFromSource2() {
        return new Promise((resolve) => {
            setTimeout(() => {
                resolve('Data from source 2');
            }, 1500);
        });
    }

    function fetchDataFromSource3() {
        return new Promise((resolve) => {
            setTimeout(() => {
                resolve('Data from source 3');
            }, 2000);
        });
    }

    // Asynchronous function that uses await multiple times
    async function getDataFromMultipleSources() {
        try {
            console.log('Fetching data from sources...');
            const data1 = await fetchDataFromSource1();
            console.log('Source 1:', data1);

            const data2 = await fetchDataFromSource2();
            console.log('Source 2:', data2);

            const data3 = await fetchDataFromSource3();
            console.log('Source 3:', data3);

            console.log('All data fetched successfully!');
            return [data1, data2, data3];
        } catch (error) {
            console.error('Error:', error);
            return null;
        }
    }

    getDataFromMultipleSources();


}());