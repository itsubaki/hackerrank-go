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

function isPositive(a) {
    if (a > 0) {
        return 'YES';
    }
    
    if ( a === 0 ) {
        throw new Error('Zero Error')
    }
    
    throw new Error('Negative Error');
}

function Rectangle(a, b) {
    return {
        length: a,
        width: b,
        perimeter: 2 * (a + b),
        area: a*b
    }
}

function getCount(objects) {
    return objects.filter(item => item.x === item.y).length;
}

class Polygon{
    constructor(sides){        
        this.sides = sides
    }
    
    perimeter() {
        return this.sides.reduce(function add(a, b){return a+b;})
    } 
}

class Rectangle {
    constructor(w, h) {
        this.w = w;
        this.h = h;
    }
}

Rectangle.prototype.area = function () {
    return this.w * this.h;
}

class Square extends Rectangle {
    constructor(s) {
        super(s, s)
    }
}

function sides(literals, ...expressions) {
    const [a, p] = expressions;
    let s1 = (p + Math.sqrt(p*p -16 *a))/4;
    let s2 = (p - Math.sqrt(p*p -16 *a))/4;
    return [s1, s2].sort();
}

function modifyArray(nums) {
    return nums.map(n => n = (n%2==0) ? n*2: n*3);
}

function getMaxLessThanK(n, k) {
    let max = 0;
    for (let b = n; b > 0; b--) {
        for (let a = b-1; a > 0; a--) {
            if ((a & b) < k && (a & b) > max){
                max = (a&b);
            }
        }
    }
    
    return max;
}

function getDayName(dateString) {
    let dayNames = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
    return dayNames[new Date(dateString).getDay()];
}

function regexVar() {
    return /^([aeiou]).+\1$/;
}

function regexVar() {
    return /^(Mr|Mrs|Ms|Dr|Er)(\.)([a-zA-Z])*$/;
}

function regexVar() {
    return /\d+/g;
}

// CSS/button.css
// #btn {
//     width: 96px;
//     height: 48px;
//     font-size:24px;
// }

// js/button.js
function counter() {
    var button = document.getElementById("btn");
    let count = 1;
    
    button.innerHTML = count;
    button.addEventListener("click",()=>{
        button.innerHTML = +(button.innerHTML)+1;
    });
}

// index.html
// <!-- Enter your HTML code here -->
// <!DOCTYPE html>
// <html>
//     <head>
//         <meta charset="utf-8">
//         <title>Button</title>
//     </head>
//     <body>
//         <script src="js/button.js" type="text/javascript"></script>
//         <button id='btn' onclick="counter()">0</button>
//     </body>
// </html>

// css/buttonsGrid.css
// #btns {
//   display: block;
//   overflow: hidden;
//   width: 75%;
// }
//
// #btns button {
//   display: block;
//   width: 30%;
//   float: left;
//   height: 48px;
//   font-size: 24px;
// }

window.onload = () => {
    const button5 = document.getElementById('btn5');
    button5.addEventListener('click', () => {
      let arr = [
        document.getElementById('btn1').innerText,
        document.getElementById('btn2').innerText,
        document.getElementById('btn3').innerText,
        document.getElementById('btn6').innerText,
        document.getElementById('btn9').innerText,
        document.getElementById('btn8').innerText,
        document.getElementById('btn7').innerText,
        document.getElementById('btn4').innerText
      ];
  
      arr = [...arr.slice(arr.length - 1), ...arr.slice(0, arr.length - 1)];
      document.getElementById('btn1').innerText = arr[0];
      document.getElementById('btn2').innerText = arr[1];
      document.getElementById('btn3').innerText = arr[2];
      document.getElementById('btn6').innerText = arr[3];
      document.getElementById('btn9').innerText = arr[4];
      document.getElementById('btn8').innerText = arr[5];
      document.getElementById('btn7').innerText = arr[6];
      document.getElementById('btn4').innerText = arr[7];
    });
  };

// index.html
// <!DOCTYPE html>
// <html lang="en">
//   <head>
//     <meta charset="UTF-8" />
//     <link rel="stylesheet" href="css/buttonsGrid.css" type="text/css" />
//   </head>
//   <body>
//     <div id="btns">
//       <button id="btn1">1</button> <button id="btn2">2</button> <button id="btn3">3</button>
//       <button id="btn4">4</button> <button id="btn5">5</button> <button id="btn6">6</button>
//       <button id="btn7">7</button> <button id="btn8">8</button> <button id="btn9">9</button>
//     </div>
//     <script src="js/buttonsGrid.js" type="text/javascript"></script>
//   </body>
// </html>
