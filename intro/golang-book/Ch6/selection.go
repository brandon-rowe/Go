package main 

import "fmt"

func main() {
	var hi int = 0;
	var lo int = 0;
	x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17}
	for index, value := range x {
		if value > hi
			hi = value
		if value < lo
			lo = value
	}
	fmt.Println(hi)
	fmt.Println(lo)
}


//Current error message
//./selection.go:11:14: syntax error: assignment hi = value used as value
//./selection.go:13:14: syntax error: assignment lo = value used as value