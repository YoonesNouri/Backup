//9.2

(function () {
    'use strict';

    // Sample asynchronous functions returning Promises
    function fetchUserData() {
        return new Promise((resolve) => {
            setTimeout(() => {
                const userData = { name: 'John', age: 30 };
                resolve(userData);
            }, 1000);
        });
    }

    function fetchUserPosts(userId) {
        return new Promise((resolve) => {
            setTimeout(() => {
                const posts = [
                    { id: 1, title: 'First Post', body: 'This is the first post.' },
                    { id: 2, title: 'Second Post', body: 'This is the second post.' }
                ];
                resolve(posts);
            }, 1500);
        });
    }

    function fetchPostComments(postId) {
        return new Promise((resolve) => {
            setTimeout(() => {
                const comments = [
                    { id: 101, text: 'Comment 1 on post 1' },
                    { id: 102, text: 'Comment 2 on post 1' }
                ];
                resolve(comments);
            }, 2000);
        });
    }

    // Promise chaining example
    fetchUserData()
        .then((user) => {
            console.log('User:', user);
            return fetchUserPosts(user.id); // Chaining the next operation
        })
        .then((posts) => {
            console.log('Posts:', posts);
            return fetchPostComments(posts[0].id); // Chaining the next operation
        })
        .then((comments) => {
            console.log('Comments:', comments);
        })
        .catch((error) => {
            console.error('Error:', error); // Handling any errors that occur during the chain
        });

}());