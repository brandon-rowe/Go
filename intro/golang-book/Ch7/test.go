package main 

import "fmt"

func add(x uint, y uint) uint {
	return x + y
}

func sub(x uint, y uint) uint {
	return x - y
}

func mul(x uint, y uint) uint {
	return x * y
}

func div(x uint, y uint) uint {
	return x / y
}

func main() {
	fmt.Println(add(5,9))
	fmt.Println(sub(10,333))
	fmt.Println(mul(1000, 450))
	fmt.Println(div(234,975))
}