function repeatLetters(word) {
    let result = '';

    for (let i = 0; i < word.length; i++) {
        for (let j = 0; j <= i; j++) {
            result += word[i];
        }
        result += '\n';
    }

    return result;
}

const word = 'Amir';
const result = repeatLetters(word);
console.log(result);
