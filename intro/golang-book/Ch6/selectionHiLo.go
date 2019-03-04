package main 

import "fmt"

func main() {
	var hi int = 0
	var lo int = 0
	//var lo int = 100		This is the easy fix for this problem. 
	//We will come up with a more elegant/typical solution
	x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17}
	for i, v := range x{
		if i==0 {
			hi := v
			lo := v
			break
		}
		
	}
	
	for index, value := range x {
		if value > hi{
			hi = value
		}
		if value < lo{
			lo = value
		}
		index++
	}
	fmt.Println(hi)
	fmt.Println(lo)
}


//Current problem
//lo value will not work because 0
//is lower than all values.
//A simple fix would be to initialize lo to 100.
//However, this is not practical for scalability.
//I will need to initialize each value, hi & lo, 
//to the first value in the array/list.