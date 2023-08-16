//7.2

(function () {
    'use strict';
    // AJAX request using the Fetch API:
    fetch('https://api.example.com/data')
        .then(response => response.json())
        .then(data => {
            console.log(data);  // Process the received data
        })
        .catch(error => {
            console.error('Error:', error);
        });
}());

(function () {
    'use strict';
    // open(method, url[, async[, user[, password]]]):
    const xhr = new XMLHttpRequest();
    xhr.open('GET', 'https://api.example.com/data', true);
}());

(function () {
    'use strict';
    // send([body]):
    const xhr = new XMLHttpRequest();
    xhr.open('POST', 'https://api.example.com/submit', true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({ name: 'John', age: 30 }));
}());

(function () {
    'use strict';
    // onreadystatechange:
    const xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            console.log(xhr.responseText);  // Handle the response
        }
    };
    xhr.open('GET', 'https://api.example.com/data', true);
    xhr.send();
}());

(function () {
    'use strict';
    // onload:
    const xhr = new XMLHttpRequest();
    xhr.onload = function () {
        console.log(xhr.responseText);  // Handle the response
    };
    xhr.open('GET', 'https://api.example.com/data', true);
    xhr.send();
}());

(function () {
    'use strict';
    // onerror:
    const xhr = new XMLHttpRequest();
    xhr.onerror = function () {
        console.error('An error occurred during the request');
    };
    xhr.open('GET', 'https://api.example.com/data', true);
    xhr.send();
}());

(function () {
    'use strict';

}());