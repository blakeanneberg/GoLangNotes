package main

import (
  "fmt"
  "math"
)

func main () {
  fmt.Println("Number of apples?")
  var appleNum int8 
  fmt.Scanln(&appleNum)
  appleNum := int8(math.Trunc(appleNum))
  fmt.Printf("%d\n", appleNum)
}
