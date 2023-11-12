# Links: 
- GoLangDocs: https://go.dev/doc/#learning
- Coursera: https://www.coursera.org/learn/golang-getting-started/home/welcome

# Notes: 
Object-Oriented Programming:
-Organize your code though encapsulation
-Group together data and functions which are related.
-User defined type which is specific to an application. Example: hints have data (number, and functions +,-, etc.). Functions but no class (defines data and functions)

## Objects in Go:
-Go uses struts with associated methods (not using term class)
-Can associate methods or functions with those struts

## Concurrency in Go:
-Managing multiple tasks at the same time.
-Go includes Concurrency primitives
-`Goroutines` represent concurrent tasks
-`channels` are used to communicate between tasks
-`select` enables task synchronization

## Packages
-First line of file names the package
-Package called Main `main`, makes an executable
  -Needs a function `main()`
  -`main()` is where the code execution starts



## Go Tools 
-`go build`
  -compiles the program
  -creates an executable for the main package, same name as the first .go file
  -.exe suffix for executable in winds
-`go doc`
  -prints documentation for a package
-`go format`
  -formats source code files

## Variables:
- `var x int` `var x int`
- `new()` function creates a variable and returns a pointer to the variable. Variable is initialized to zero by default. 

### Variable Scope:
- Places in code where a variable can be accessed
``` golang
var x = 4 

func f() {
  fmt.Printf("%d", x)
}
func g() {
  fmt.Printf("%d", x)
}
```
- Blocks
  - A sequence of declaratins and statements within matching brackets, {}
  - Including functiion definitions
  - Universe block: all Go source
  - Package block: all source in a package
  - File block: all source in a file
  - Code inside the statement: "if", "for", "switch"
  - Clauses in "switch" or "select": individual clauses each get a block

- Lexical Scoping
  - Go is lexically scoped using blocks: bi >= bj if bj is defined inside bi

## Basic data types
- Pointers: an address to data in memory and have operators
  - `&` operator returns the address of a variable/function, put in front of a variable, will return the address of that variable. 
  - `*` operator returns data at an address (opposite of & called dereferencing), if put * in front of address, will give you data at that address
  ``` golang
  var x int = 1 
  var y int // not inilizied so it default to zero
  var ip *int // ip is pointer to int

  ip = &x // ip now points to x 
  y = *op // y is now 1
  ``` 
  - see `new()` 
  ``` golang
  ptr := new(int)
  *ptr = 3
  ``` 

- Printing
  - Conversion characters for each argument
  ``` golang

  fmt.Printf("hi %s", x)

  ```

- Integers
  - Generic int declaration
    ``` golang 

    var x int 

    ```
  
  - Different lengths and signs: int8, int16, int32, int64, uint8, uint16, uint32, uint64
  - Binary operators
    - Arthmetric: = - * / % << >>
    - Comparison: == != > < >= <=
    - Boolian && || 

  - Type conversion: most binary operations need operatnd of same type. Ex including assignments
  ```golang
  var x int32 = 1
  var y int16 = 2 
  x = y 
  ```
  - Convert type with T()operation
  x = int32(y)

  - Floating points: real numbers, ex `float32` gives your ~6 digitis of precision, `float64` gives you ~15 digitis of precesion  
    - Express using decimals or scientific notation 
    ```golang
    var x float64 = 123.45
    var y float64 = 1.2345e2
    ```
    - Complex numbers represented as two floats: real and imaginary
    ```golang
    var z complex128 =
    complex(2,3)
    ```

- Strings
  - ASCII and Unicode
    - Character coding - each character is associated with an (7) 8-Bit number ex. 'A'=0x41 in hexadecimal 
    - Unicode is 32-bit character. UTF-8 is variable length 8-bit UTF codes are the same as ASCII  
    - Code points are Unicode characters, and a Rune is a code point in Go
  - String literal are notated by double quotes where each byte is a rune (UTF-8 code point)
    ```golang 
    x := "Hi there"
    ```
  - Unicode Package
    - Runes are divided into many different categories
    - Provides a set of functions to test categories of runes
    ```golang
    IsDigit(r rune)
    IsSpace(r rune)
    IsLetter(r rune)
    IsLower(r rune)
    IsPunct(r rune)
    ```
    - Some functions perform conversions
    ```golang
    ToUpper(r rune)
    ToLower(r rune)
    ```
  - Strings Package of functions to manipulate UTF-8 encoded strings, strings are immutable but moddifed strings are returned
    - `compare (a, B)` returns an integer comparing two strings lexicographically. 0 if a==b if a < b, and +1 if a > b. 
    - `Contains(s, substr)` returns true if substring is inside s
    - `Replace(s, old, new, n)` replace returns a copy of the string s with the first n instances of old replaced by new
    - `ToLower(s)`
    - `ToUpper(s)`
    - `TrimSpace(s)` returns a new string with all leading and trailing white space removed
  - Strconv Package
    - conversion to and from string representations of basic data types
    - `Atoi (s)` converts string s to int
    - `Itoa (s)` converts int (base 10) to string
    - `FormatFloat (f, fmt, prec, bitSize)` converts floating point number to a string
    - `ParseFloat (s, bitSize)` converts a string to a floating point number
  - Constants
    - Expression whose value is known at compile time.
    - Type is inferred from right-hand side (boolean, string, number)
    ```golang
    const x = 1.3 
    const ()
      y = 4 
      z = "Hi"
    ```
    - Iota example, each constant is assigned to a unique integer, constants must be different but actual value is not important
    ```golang 
    type Grades int
    const (
      A Grades = iota
      B
      C
      D
      F
    )
    ```
  - Control flow
    - statements which alter control flow, expression `<condition>` is evaluated, `<consequent>` statments are executed if condition is true
    ```golang
    if <condition> {
      <consequent.
    }
    ```
    - For Loops, iterates while a condition is true, may have an initialization nd update operation
    ```golang
    for <init>; <cond>;
    <update> {
      <stmts>
    }
    ```
    - example 1: 
    ```golang
    for i:=0; i<10; i++ {
      fmt.Printf("hi ")
    }
    ```
    - example 2: 
    ```golang
    i = 0 
    for i < 10 {
      fmt.Printf("hi ")
      i++
    }
    ```
    - example 3: 
    ```golang
    for [fmt.Printf("hi ")]
    ```
  - Switch/Case
    - Switch is a multi-way if statement
    - Switch may contain a tag which is a variable to be checked
    - Tag is compared to a constant defined to each case
    - case which matches tag is executed
    - example 1:
    ```golang
    switch x {
    case 1: 
      fmt.Printf("case1")
    case 2: 
      fmt.Printf("case2")  
    default: 
      fmt.Printf("nocase")
    }
    ```
  - Tagless Switch, a switch without a tag, case contains a boolean expression to evaluate, first true case is executed
    - example 1 to use if you dont wnat to use if else, if else...: 
    ```golang
    switch {
    case x > 1: 
      fmt.printf("case1")
    case x < -1:
      fmt.Printf("case2")
    default: 
      fmt.Printf("nocase")
    }
    ```
  - Break and Continue
    - break exits the containing loop, considered bad form
    - example 1
    ```golang
    i := 0 
    for i < 10 {
      i++ 
      if i == 5 { break }
      fmt.Printf("hi ")
    }
    ```
    - Continue skips the rest of the current iteration
    - example 2 
    ```golang
    i := 0 
    for i < 10 {
      i++ 
      if i == 5 { continue }
      fmt.Printf("hi ")
    }
    ```
  - Scan
    - Scan reads user input
    - Takes a pointer as an argument 
    - Typed data is written to pointer
    - example 1, when the user types in five and hits enter, that scan function takes that number five, puts it into the appleNum variable. So on the next line, when I say, printf appleNum, it'll print five or whatever integer they typed in.
    ```golang
    var appleNum int

    fmt.Printf("number of apples?")
    num, err :=
    fmt.scan(&appleNum) // means the address of the AppleNum variable, 
    fmt.Printf(appleNum)
    ```
