package main

import "fmt"

func main() {
	fmt.Print("Please enter the number of students to enter: ")
	var x int
	fmt.Scanln(&x)
	fname := make([]string, x)
	for i:=0; i<x; i++{
		fmt.Print("Please enter the first name of student ",i+1,": " )
		var tmp string
		fmt.Scanln(&tmp)
		fname[i] = tmp
	}

	lname := make([]string, x)
  	for i:=0; i<x; i++{
	    fmt.Print("Please enter the last name of student ",i+1,": " )
	    var tmp string
	    fmt.Scanln(&tmp)
	    lname[i] = tmp
	}
	fmt.Println("Students' first names are: ")
	fmt.Println(fname)
	fmt.Println("Students' last names are: ")
  	fmt.Println(lname)
  	fmt.Println()

  	for i:=0; i<x; i++{
  		fmt.Println("Student ",i+1,": ", fname[i], lname[i])
  	}
}

