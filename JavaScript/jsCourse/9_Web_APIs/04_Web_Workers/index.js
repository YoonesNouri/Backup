//9.4

// main.js
const worker = new Worker('worker.js');
worker.addEventListener('message', (event) => {
    // Receive response from the worker
    const response = event.data;
    console.log('Received response from worker:', response);
});
// Send a message to the worker
const dataToSend = 5;
console.log('Sending message to worker:', dataToSend);
worker.postMessage(dataToSend);
// setTimeout(worker.terminate(), 5000);
