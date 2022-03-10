function greeting(parameterVariable) {
    console.log('Hello, World!');
    console.log(parameterVariable)
}

function performOperation(secondInteger, secondDecimal, secondString) {
    const firstInteger = 4;
    const firstDecimal = 4.0;
    const firstString = 'HackerRank ';
    
    console.log(firstInteger + parseInt(secondInteger))
    console.log(firstDecimal + parseFloat(secondDecimal))
    console.log(firstString  + secondString)
}

function getArea(length, width) {
    return length * width;
}

function getPerimeter(length, width) {
    return 2 * (length + width);
}

function factorial(n) {
    if (n === 0) {
        return 1
    }
    
    return n * factorial(n - 1)
}

function vowelsAndConsonants(s) {
    const vowels = 'aeiou';
    var consonants = '';

    for (var i = 0; i < s.length; i++) {
        if (vowels.includes(s[i])) {
            console.log(s[i]);
            continue;
        }

        consonants = consonants + s[i] + '\n';
    }
    
    console.log(consonants);
}

function letAndConst() {
    const PI = Math.PI;
    let r = readLine();
    
    console.log(PI*r*r);
    console.log(PI*2*r);
}
