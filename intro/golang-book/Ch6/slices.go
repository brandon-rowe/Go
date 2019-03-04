
//We can manipulate arrays by slicing them.
package main 

import "fmt"

func main() {
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("Slice 1: ",slice1)
	fmt.Println("Slice 2: ",slice2)
	slice3 := []int{6,7,8}
	fmt.Println("Slice 3: ",slice1, "before make slice")
	slice4 := make([]int, 2)
	copy(slice3, slice4)
	fmt.Println("Slice 3 & 4 after make slice and copy")
	fmt.Println(slice3, slice4)
}