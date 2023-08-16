//8.5

(function () {
    'use strict';
    // JavaScript code to create and append a comment
    const commentText = "This is a comment.";
    const comment = document.createComment(commentText);
    // Get the target element
    const commentContainer = document.getElementById("commentContainer");
    // Append the comment to the target element
    commentContainer.appendChild(comment);
}());

(function () {
    'use strict';
    // JavaScript code to create and append elements using DocumentFragment
    const fragment = document.createDocumentFragment();
    for (let i = 1; i <= 5; i++) {
        const newElement = document.createElement("p");
        newElement.textContent = `Paragraph ${i}`;
        fragment.appendChild(newElement);
    }
    // Get the target element
    const container = document.getElementById("container");
    // Append all elements to the container at once
    container.appendChild(fragment);
}());