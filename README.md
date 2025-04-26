# tlang
Repo for TLang [Timmy's Programming Language]. An interpreted programming language built while reading Thorsten Ball's Writing an interpreter in Go book


  <body>
    <div class="container">
      <h1>TLang Programming Documentation</h1>
      <p>
        Welcome to the official documentation for Timmy's Interpreted
        Programming Language, a modern and simple programming language.
        <a href="https://rafmme.github.io/tlang_editor" target="_blank">Try it out here</a>
        <a href="https://rafmme.github.io/tlang_editor/doc" target="_blank">More documentation</a>
      </p>
      <p>
        Author: Timileyin E. Farayola |
        <span>
          <a href="https://linkedin.com/in/timileyin-farayola"
            >https://linkedin.com/in/timileyin-farayola</a
          ></span
        >
        |
        <span
          ><a href="https://github.com/rafmme"
            >https://github.com/rafmme</a
          ></span
        >
        |
        <span
          ><a href="mailto:timileyin.e.farayola@gmail.com"
            >timileyin.e.farayola@gmail.com</a
          ></span
        >
        | &copy; APRIL, 2025
      </p>
    </div>
  </body>

  <h2>Data Types</h2>

  ```
String: 

make str_one = "string 1";
make str_two = `string 2`;
make str_th = 'string 3';

count(str_two) // Returns the length of the string. //Output: 8

display(str_one * 3) // Output: "string 1string 1string 1"

Integer: 

make num = 26;

Boolean: 

make isCorrect = true;
make notTrue = false;

Null: 

make a = null;
```



<h2>Operators</h2>
<p>Supported operators.</p>

  ```
Arithmetic Operators: 

make num = 3;

Multiplication "*" display(num * num) //Output: 9

Division "/" display(num / num) //Output: 1

Addition "+" display(num + num) //Output: 6

Subtraction "-" display(num - num) //Output: 0

display(((num * num) + (num + num)) / num) //Output: 5


Logical Operators: 

!true // false

!false // true


Comparison operators: 

3 == 3 // true

3 > 4 // false

4 < 5 // true
```


<h2>Arrays</h2>
<p>Arrays store multiple values in a single variable.</p>

```
make arr = [1, 2, 3, 4]

display(arr[1]) // Output: 2
```


<h2>Hashmaps</h2>
<p>Hashmaps (dictionaries) store key-value pairs.</p>

```
make map = {"name": "Alice", "age": 25}

display(map["name"]) // Output: Alice
```

<h2>Built-in Functions</h2>
<p>TLang provides several built-in functions for handling data andcomputations. </p>

```
display(element) // Outputs the element value.

count(arr) // Returns the number of elements in an array.

push(arr, element) //Adds a new element to an array. Returns a new array with the added element.

first(arr) // Returns the first element in an array.

last(arr)  // Returns the last element in an array.

rest(arr) // Returns a new array without the first element.



TODOS:  // To be added!

keys(map)  // Returns all keys in a hashmap

values(map) // Returns all values in a hashmap
```

<h2>Functions</h2>
<p>Functions are reusable blocks of code. Supports High-Order functions & Closures. </p>

```
make greet = fxn(name) {
  return "Hello, " + name;
}

display(greet("Bob")); // Output: Hello, Bob

A recursive function that loops through an array and print all its element: 

make loopThroughArray = fxn(arr) {
  if(count(arr) == 0){
    return null;
  } else {
      display(first(arr));

      loopThroughArray(rest(arr));
      return null;
  }
}

make a = [74,0,1,2,3,4,5,6,7,8,9];

loopThroughArray(a);
```

<h2>If-Else Conditional</h2>
<p>Conditional statements control program flow.</p>

```
make x = 10;

if (x > 5) {
    display("x is greater than 5");
} else {
    display("x is less than or equal to 5");
}
```

<h2>More features to be added soon. [WIP]</h2>