package main

import "fmt"

func main() {
    fmt.Println("This is the String Test")
    fmt.Println(len("Hello go lang"))
    fmt.Println("Hello go lang"[1])
    fmt.Println("Hello " + "go lang")
    fmt.Println()
    fmt.Println("This is the Boolean Test")
    fmt.Println("True & True Evaluates to", true && true)
    fmt.Println("True & False Evaluates to",true && false)
    fmt.Println("True or True Evaluates to", true || true)
    fmt.Println("True or False Evaluates to", true || false)
    fmt.Println("Not True Evaluates to",!true)
    fmt.Println("Not False Evaluates to",!false)
}
