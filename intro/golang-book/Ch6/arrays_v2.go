package main 
//imporvement from arrays.go but counting strings
import "fmt"

func main() {
	x := [10]string {"jon", "brandon", "joe", "jimmy", "josh", "jared", "jordan", "kailyn", "kyle", "jp" }
	
	var total float64 = 0
	for i :=0; i<len(x); i++{
		total += 1
	}
	fmt.Println("Array x consists of: ")
	fmt.Println(x)
	fmt.Println("With a total of ", total, " objects")
	fmt.Println(total)
}