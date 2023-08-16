//8.2


// 1. Changing InnerHTML and outerHTML:

(function () {
    'use strict';
    // InnerHTML:
    const innerDivElement = document.getElementById("innerDiv");
    innerDivElement.innerHTML = "Updated content"; // Change the content of the div
    console.log(innerDivElement);
}());
(function () {
    'use strict';
    // outerHTML:
    const outerDiv = document.getElementById('outerDiv');
    console.log(outerDiv.outerHTML);
    // Update the content of the div using outerHTML
    outerDiv.outerHTML = '<div id="myDiv">This is the updated content</div>';
})();


(function () {
    'use strict';
    // 2. Changing Text Content: 
    const myParagraph = document.getElementById("myParagraph");
    myParagraph.textContent = "Greetings, everyone!"; // Change the text content
    console.log(myParagraph);
}());

(function () {
    'use strict';
    // 3. Changing Attributes: 
    const myImage = document.getElementById("myImage");
    myImage.setAttribute("src", "new-image.jpg"); // Change the 'src' attribute
    myImage.setAttribute("alt", "New Image"); // Change the 'alt' attribute
    console.log(myImage);
}());

(function () {
    'use strict';
    // 4. Styling Elements: 
    const myBox = document.getElementById("myBox");
    myBox.style.backgroundColor = "blue"; // Change the background color
    myBox.style.color = "white"; // Change the text color
    console.log(myBox);
}());

(function () {
    'use strict';
    // 5. Adding and Removing Elements: 
    const myList = document.getElementById("myList");
    // Add a new list item
    const newItem = document.createElement("li");
    newItem.textContent = "Item 3";
    myList.appendChild(newItem);
    // Remove the first list item
    const firstItem = myList.querySelector("li");
    myList.removeChild(firstItem);
    console.log(myList);

    // Using "contains" method:
    if (myList.contains(newItem)) {
        console.log("newItem is in myList");
    } else {
        console.log("newItem is not in myList");
    }
}());

(function () {
    'use strict';
    // 6. Event Handling: 
    const myButton = document.getElementById("myButton");
    myButton.addEventListener("click", function () {
        alert("Button clicked!");
    });
    console.log(myButton);
}());

(function () {
    'use strict';
    let heading = document.getElementById('mainHeading')
    console.log(heading.id);
    heading.id = 'newHeading'
    console.log(heading.id);
}());

