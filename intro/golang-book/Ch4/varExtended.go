package main

import "fmt"
//Establishing var outside of the func makes them accessible to other functions.
var x string = "together we stand"

func main() {
	//fmt.Println("A shorter method for assigning variables")
	//x := 5
    //fmt.Println(x)
    //y := "Max"
    //fmt.Println(y)
    fmt.Println(x)
}

func f(){
	fmt.Println(x)
}
