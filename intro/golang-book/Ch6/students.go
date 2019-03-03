package main

import "fmt"

func main() {
	fmt.Print("Please enter the number of students to enter: ")
	var x float32
	fmt.Scanf("%f", &x)
	var name [x]string
	for i:=0; i<(x); i++
	{
		fmt.Printf("Please enter the first name")
		var tmp string
		name[i] = fmt.("%s", &tmp)
	}
}