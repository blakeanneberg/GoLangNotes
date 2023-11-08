Links: 
- GoLangDocs: https://go.dev/doc/#learning
- Coursera: https://www.coursera.org/learn/golang-getting-started/home/welcome

Notes: 
Object-Oriented Programming:
-Organize your code though encapsulation
-Group together data and functions which are related.
-User defined type which is specific to an application. Example: hints have data (number, and functions +,-, etc.). Functions but no class (defines data and functions)

Objects in Go:
-Go uses struts with associated methods (not using term class)
-Can associate methods or functions with those struts

Concurrency in Go:
-Managing multiple tasks at the same time.
-Go includes Concurrency primitives
-`Goroutines` represent concurrent tasks
-`channels` are used to communicate between tasks
-`select` enables task synchronization

Packages
-First line of file names the package
-Package called Main `main`, makes an executable
  -Needs a function `main()`
  -`main()` is where the code execution starts

| Pro   | Con    |
|--------------- | --------------- |
| Icecream   | cake   |



Go Tool:
-`go build`
  -compiles the program
  -creates an executable for the main package, same name as the first .go file
  -.exe suffix for executable in winds
-`go doc`
  -prints documentation for a package
-`go format`
  -formats source code files

Variables:
'var x int' `var x int`
