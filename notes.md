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

| Pro   | Con    |
|--------------- | --------------- |
| Icecream   | cake   |



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
  
