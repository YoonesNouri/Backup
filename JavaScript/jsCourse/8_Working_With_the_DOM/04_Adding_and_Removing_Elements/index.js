//8.4

(function () {
    'use strict';
    // 1. Inserting Elements at Specific Positions:
    const referenceElement = document.getElementById('reference'); // Get the reference element
    const parentElement = referenceElement.parentElement; // Get the parent element
    const newDiv = document.createElement('div'); // Create the new <div> element
    newDiv.textContent = 'This is the new element'; // Set content for the new <div>
    parentElement.insertBefore(newDiv, referenceElement); // Inserts the new <div> before the reference element
}());

(function () {
    'use strict';
    // 2. Replacing Elements:
    const oldElement = document.getElementById('old'); // Get the old element
    const parentElement = oldElement.parentElement; // Get the parent element
    const newDiv = document.createElement('div'); // Create the new <div> element
    newDiv.textContent = 'This is the replacement element'; // Set content for the new <div>
    parentElement.replaceChild(newDiv, oldElement); // Replaces the old element with the new <div>
}());

(function () {
    'use strict';
    // 3. Removing Elements:
    const childElement = document.getElementById('child'); // Get the child element to remove
    const parentElement = childElement.parentElement; // Get the parent element
    parentElement.removeChild(childElement); // Removes the child element from the parent
}());

(function () {
    'use strict';
    // 4. Adding and Removing Classes:

    // Add by assign:
    const elementAssign = document.getElementById('myElementAssign');
    elementAssign.className = 'class1 class2';
    elementAssign.className = 'class2';
    console.log(elementAssign.className); // Result: "class2" (replaces any existing classes with the specified class)

    // Add by ClassList:
    const elementAdd = document.getElementById('myElementAdd');
    elementAdd.className = 'class1 class2';
    elementAdd.classList.add('class2');
    console.log(elementAdd.className); // Result: "class1 class2" (class "class2" is already present, has no effect)

    // Remove Class "class2":
    elementAdd.classList.remove('class2'); // Removes the class 'class2' from elementAdd.
    console.log(elementAdd.className); // Result: "class1" (only "class1" remains after removal)
}());
 