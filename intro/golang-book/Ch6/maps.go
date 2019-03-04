// A map is an unordered collection of key-value pairs
// AKA associative array, hash table, dictionary
// Maps are used to look up value by associated key
package main 

import "fmt"

func main() {
	/*x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])*/
	elements := make(map[string]string)
  	elements["H"] = "Hydrogen"
  	elements["He"] = "Helium"
  	elements["Li"] = "Lithium"
  	elements["Be"] = "Beryllium"
  	elements["B"] = "Boron"
  	elements["C"] = "Carbon"
  	elements["N"] = "Nitrogen"
  	elements["O"] = "Oxygen"
 	elements["F"] = "Fluorine"
  	elements["Ne"] = "Neon"

  	fmt.Println("Elements in order: ", elements)

  	fmt.Println("Elements called individually")

  	fmt.Println(elements["Li"])
  	fmt.Println(elements["He"])
  	fmt.Println(elements["C"])
  	fmt.Println(elements["H"])
  	fmt.Println(elements["Be"])
  	fmt.Println(elements["N"])
  	fmt.Println(elements["Ne"])
  	fmt.Println(elements["F"])
  	fmt.Println(elements["O"])
  	fmt.Println(elements["B"])
}