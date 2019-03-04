// A map is an unordered collection of key-value pairs
// AKA associative array, hash table, dictionary
// Maps are used to look up value by associated key
package main 

import "fmt"

func main() {
	elements := map[string]map[string]string{
    "H": map[string]string{
      "name":"Hydrogen",
      "state":"gas",
    },
    "He": map[string]string{
      "name":"Helium",
      "state":"gas",
    },
    "Li": map[string]string{
      "name":"Lithium",
      "state":"solid",
    },
    "Be": map[string]string{
      "name":"Beryllium",
      "state":"solid",
    },
    "B":  map[string]string{
      "name":"Boron",
      "state":"solid",
    },
    "C":  map[string]string{
      "name":"Carbon",
      "state":"solid",
    },
    "N":  map[string]string{
      "name":"Nitrogen",
      "state":"gas",
    },
    "O":  map[string]string{
      "name":"Oxygen",
      "state":"gas",
    },
    "F":  map[string]string{
      "name":"Fluorine",
      "state":"gas",
    },
    "Ne":  map[string]string{
      "name":"Neon",
      "state":"gas",
    },
  }

  if el, ok := elements["Li"]; ok {
    fmt.Println(el["name"], el["state"])
  }
  if el, ok := elements["O"]; ok {
    fmt.Println(el["name"], el["state"])
  }
  if el, ok := elements["C"]; ok {
    fmt.Println(el["name"], el["state"])
  }
  if el, ok := elements["B"]; ok {
    fmt.Println(el["name"], el["state"])
  }
  if el, ok := elements["Ne"]; ok {
    fmt.Println(el["name"], el["state"])
  }
  if el, ok := elements["F"]; ok {
    fmt.Println(el["name"], el["state"])
  }
  if el, ok := elements["Be"]; ok {
    fmt.Println(el["name"], el["state"])
  }
  if el, ok := elements["H"]; ok {
    fmt.Println(el["name"], el["state"])
  }
  if el, ok := elements["Li"]; ok {
    fmt.Println(el["name"], el["state"])
  }
  if el, ok := elements["He"]; ok {
    fmt.Println(el["name"], el["state"])
    if el, ok := elements["N"]; ok {
    fmt.Println(el["name"], el["state"])
  }
  }

}























