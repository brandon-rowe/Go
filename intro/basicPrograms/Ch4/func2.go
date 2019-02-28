package main

import "fmt"

func main() {
  fmt.Print("Enter hours worked this week: ")
  var hours float64
  fmt.Scanf("%f", &hours)
  fmt.Print("Enter hourly pay rate: ")
  var payrate float64
  fmt.Scanf("%f", &payrate)
  output := hours * payrate
  fmt.Println(output)
}
