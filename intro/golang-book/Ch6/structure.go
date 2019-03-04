//I am jumping ahead a little from the book.
//I made a personal small project to create
//a class in go with at least 1 method(func)
//outside of the main.

package main

import "fmt"

type Student struct{
	FristName, LastName string
}

func (u Student) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName,"/t Welcome to go!")
}

func main() {
	fmt.Print("Please enter the number of students to enter: ")
	var x int
	fmt.Scanln(&x)
	fname := make([]string, x)
	lname := make([]string, x)
	for i:=0; i<x; i++{
		fmt.Print("Please enter the first name of student ",i+1,": " )
		var tmp string
		fmt.Scanln(&tmp)
		fname[i] = tmp
		fmt.Print("Please enter the last name of student ",i+1,": " )
	    fmt.Scanln(&tmp)
	    lname[i] = tmp
	}
  	fmt.Println()

  	for i:=0; i<x; i++{
  		fmt.Println("Student ",i+1,": ", fname[i], lname[i])
  	}
}

