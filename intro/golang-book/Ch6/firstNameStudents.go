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
}


//What I learned.
//go does not permit assignment of array length 
//from a calculated var at runtime.
//Ex. var fname [x]string

//Must use make like in line 14