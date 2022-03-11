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

function letAndConst() {
    const PI = Math.PI;
    let r = readLine();
    
    console.log(PI*r*r);
    console.log(PI*2*r);
}

function getGrade(score) {
    if ( score <= 5 ) {
        return 'F';
    }
    if ( score <= 10 ) {
        return 'E';
    }
    if ( score <= 15 ) {
        return 'D';
    }
    if ( score <= 20 ) {
        return 'C';
    }
    if ( score <= 25 ) {
        return 'B';
    }

    return 'A';
}

function getLetter(s) {
    switch(s[0]) {
        case 'a':
        case 'e':
        case 'i':
        case 'o':
        case 'u':
            return 'A';
        case 'b':
        case 'c':
        case 'd':
        case 'f':
        case 'g':
            return 'B';
        case 'h':
        case 'j':
        case 'k':
        case 'l':
        case 'm':
            return 'C';
        default:
            return 'D';
    }
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

function getSecondLargest(nums) {
    let first = nums[0];
    let second = -1;
    
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] > first) {
            second = first;
            first = nums[i]
        }
        
        if (nums[i] > second && nums[i] < first) {
            second = nums[i];
        }
    }
    
    return second;
}

function reverseString(s) {
    try{
        console.log(s.split("").reverse().join("")) 
    } catch(e) {
        console.log(e.message);
        console.log(s);
    }
}
