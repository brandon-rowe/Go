//This func creates an array of strings.
//Array size is determined by user input.
//Once array size is determed, a loop allows
//user to input first name os students.

package main

import "fmt"

func main() {
	fmt.Print("Please enter the number of students to enter: ")
	var x int
	fmt.Scanln(&x)
	fname := make([]string, x)
	for i:=0; i<x; i++{
		fmt.Print("Please enter the first name: ")
		var tmp string
		fmt.Scanln(&tmp)
		fname[i] = tmp
	}
	fmt.Println(fname)

	lname := make([]string, x)
  	for i:=0; i<x; i++{
	    fmt.Print("Please enter the first name: ")
	    var tmp string
	    fmt.Scanln(&tmp)
	    lname[i] = tmp
	}
  	fmt.Println(lname)
}

