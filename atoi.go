package main

import (
  "strconv"
  "fmt"
) 

func start() {
  i, _ := strconv.Atoi("10")
  y := i * 2
  fmt.Println(y)
}

func main() {
  start()
}
