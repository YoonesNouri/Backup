// worker.js
self.addEventListener('message', (event) => {
    // Receive message from the main thread
    const data = event.data;
    console.log('Received message in worker:', data);

    // Process the data (You can do some work here)
    const processedData = data * 2;

    // Send the response back to the main thread
    self.postMessage(processedData);
});
