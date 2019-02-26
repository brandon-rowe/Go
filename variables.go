package main

import "fmt"

func main() {
    fmt.Println("Print String Variables")
    var x string = "Hello World"
    fmt.Println(x)
    var y string
    y = "Goodbye World"
    fmt.Println(y)
    fmt.Println()

    fmt.Println("Concat Strings with Variables")
    var z string
    z = "first "
    fmt.Println(z)
    z = z + "second"
    fmt.Println(z)
    fmt.Println()

    fmt.Println("Evaluate 2 Strings")
    var a string = "hello"
    fmt.Println("String a =",a)
    var b string = "world"
    fmt.Println("String b =",b)
    fmt.Println("a==b evaluates to",a == b)
}
