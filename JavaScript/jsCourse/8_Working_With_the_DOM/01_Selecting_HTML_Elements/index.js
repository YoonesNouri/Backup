//8.1

(function () {
    'use strict';
    // Select by id:
    const myDivElement = document.getElementById("myDiv1");
    console.log(myDivElement); // Output: <div id="myDiv1">This is a div element</div>
}());

(function () {
    'use strict';
    // Select by first element that matches a CSS selector:
    const paragraphElement = document.querySelector(".container p");
    console.log(paragraphElement); // Output: <p>Hello</p>
}());

(function () {
    'use strict';
    // Selects all elements that match a CSS selector, and returns a NodeList:
    const listItems = document.querySelectorAll("ul li");
    console.log(listItems); // Output: NodeList [ <li>Item 1</li>, <li>Item 2</li> ]
}());

(function () {
    'use strict';
    // Select by tag name:
    const paragraphs = document.getElementsByTagName("p");
    console.log(paragraphs); // Output: HTMLCollection [ <p>Paragraph 1</p>, <p>Paragraph 2</p> ]
}());

(function () {
    'use strict';
    // Select by class name:
    const boxes = document.getElementsByClassName("box");
    console.log(boxes); // Output: HTMLCollection [ <div class="box">Box 1</div>, <div class="box">Box 2</div> ]
}());

(function () {
    'use strict';
    // Select by specific element to search within its descendants:
    const myDiv2 = document.getElementById("myDiv2");
    const paragraphsInsideDiv = myDiv2.querySelectorAll("p");
    console.log(paragraphsInsideDiv); // Output: NodeList [ <p>Paragraph 1</p>, <p>Paragraph 2</p> ]
}());

(function () {
    'use strict';
    // Select by :

}());



