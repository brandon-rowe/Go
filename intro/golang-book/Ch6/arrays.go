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
  array()
}
//connected two array functions from ch6

func array(){
  var x [5]float64
  x[0] = 99
  x[1] = 93
  x[2] = 77
  x[3] = 82
  x[4] = 83

  var total float64 = 0
  for i:=0; i<5; i++ {
    total += x[i]
  }
  fmt.Println("Your average is ",total/5)
}