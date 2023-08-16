function repeatLetters(word) {
    let result = '';
    let count = 1;

    for (let i = 0; i < word.length; i++) {
        for (let j = 0; j < count; j++) {
            result += word[i];
        }

        count++;
    }

    return result;
}

const word = 'hello';
const result = repeatLetters(word);
console.log(result); // Output: heelllllllooooo


