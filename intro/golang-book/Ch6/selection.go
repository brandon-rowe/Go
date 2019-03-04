package main 

import "fmt"

func main() {
	var hi int = 0;
	var lo int = 0;
	x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17}
	for i, v := range x{
		if v > hi
			hi = v
		if v < lo
			lo = v
	}
	fmt.Println(hi)
	fmt.Println(lo)
}