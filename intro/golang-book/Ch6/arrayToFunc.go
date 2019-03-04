package main

import "fmt"

func main(){
  var x [5]int
  for i := 0; i<4; i++{
    x[i] = timesTwelve(i)
    fmt.Println(x)
  }
  fmt.Println(x)
}

func timesTwelve(i int){
  var y int = i * 12
  return y
}