package main

import (
 "fmt"
 "math"
)

func main() {
 fmt.Println("Number of apples?")
 var appleNum float64
 fmt.Scanln(&appleNum)
 int(math.Trunc(appleNum))
 fmt.Printf("%d\n", appleNum)
}
