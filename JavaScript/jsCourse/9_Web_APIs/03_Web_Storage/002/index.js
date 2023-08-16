//9.3

(function () {
    'use strict';
    // setItem(key, value): This method is used to store data in the storage.
    // Session Storage
    sessionStorage.setItem('username', 'John');
    // Local Storage
    localStorage.setItem('token', 'abc123');

    // getItem(key): This method is used to retrieve data from the storage based on the specified key.
    // Session Storage
    const username = sessionStorage.getItem('username');
    // Local Storage
    const token = localStorage.getItem('token');
   
    // removeItem(key): This method is used to remove a specific item from the storage based on the specified key.
    // Session Storage
    sessionStorage.removeItem('username');
    // Local Storage
    localStorage.removeItem('token');
   
    // clear(): This method is used to clear all data stored in the storage.
    // Session Storage
    sessionStorage.clear();
    // Local Storage
    localStorage.clear();

}())

