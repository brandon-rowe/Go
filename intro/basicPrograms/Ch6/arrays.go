package main

import "fmt"

func main(){
  // An array is a numbered sequence of elements
  // of a single type with a fixed length.
  // In go, they look like this
  var x [5]int
  for i := 0; i<4; i++{
    x[i] = i * 12
    fmt.Println(x)
  }
  fmt.Println(x)
  x[4] = 100
  fmt.Println(x)
}
