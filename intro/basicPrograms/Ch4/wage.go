package main

import "fmt"

func main() {
  fmt.Print("Enter hours worked this week: ")
  var hours float32
  fmt.Scanf("%f", &hours)
  fmt.Print("Enter hourly pay rate: ")
  var payrate float32
  fmt.Scanf("%f", &payrate)
  var regHrs float32 = 40
  var output float32

  if hours > regHrs{
    var otHrs float32 = hours - regHrs
    var otWage float32 = payrate * 1.5
    var otPay float32 = otHrs * otWage
    var regPay = payrate * regHrs
    output = regPay + otPay
  } else {
    output = hours * payrate
  }
  fmt.Println("Your pay for this week is: ", output)
}
