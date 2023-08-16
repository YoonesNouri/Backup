// 8.3

// 1. parentNode:
(function () {
    'use strict';
    const parentElement = document.getElementById('child1').parentNode;
    console.log('Parent Node:', parentElement); // Output: Parent Node: <div id="parentElement">...</div>
})();

// 2. childNodes:
(function () {
    'use strict';
    const childNodes = document.getElementById('parentElement').childNodes;
    console.log('Child Nodes:', childNodes); // Output: Child Nodes: NodeList [ #text "\n ", <div#child1>, #text "\n ", <div#child2>, #text "\n ", <div#child3>, #text "\n"]
})();

// 3. children: 
(function () {
    'use strict';
    const childElements = document.getElementById('parentElement').children;
    console.log('Child Elements:', childElements); // Output: Child Elements: HTMLCollection [ <div#child1>, <div#child2>, <div#child3> ]
})();

// 4. firstChild and lastChild:
(function () {
    'use strict';
    const firstChildNode = document.getElementById('parentElement').firstChild;
    const lastChildNode = document.getElementById('parentElement').lastChild;
    console.log('First Child Node:', firstChildNode); // Output: First Child Node: #text "\n "
    console.log('Last Child Node:', lastChildNode); // Output: Last Child Node: #text "\n"
})();

// 5. firstElementChild and lastElementChild:
(function () {
    'use strict';
    const firstChildElement = document.getElementById('parentElement').firstElementChild;
    const lastChildElement = document.getElementById('parentElement').lastElementChild;
    console.log('First Child Element:', firstChildElement); // Output: First Child Element: <div id="child1">Child 1</div>
    console.log('Last Child Element:', lastChildElement); // Output: Last Child Element: <div id="child3">Child 3</div>
})();

// 6. nextSibling and previousSibling:
(function () {
    'use strict';
    const nextSiblingNode = document.getElementById('child1').nextSibling;
    const previousSiblingNode = document.getElementById('child1').previousSibling;
    console.log('Next Sibling Node:', nextSiblingNode); // Output: Next Sibling Node: #text "\n "
    console.log('Previous Sibling Node:', previousSiblingNode); // Output: Previous Sibling Node: null
})();

// 7. nextElementSibling and previousElementSibling:
(function () {
    'use strict';
    const nextSiblingElement = document.getElementById('child1').nextElementSibling;
    const previousSiblingElement = document.getElementById('child1').previousElementSibling;
    console.log('Next Sibling Element:', nextSiblingElement); // Output: Next Sibling Element: <div id="child2">Child 2</div>
    console.log('Previous Sibling Element:', previousSiblingElement); // Output: Previous Sibling Element: null
})();
