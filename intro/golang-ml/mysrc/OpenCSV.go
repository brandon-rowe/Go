package main 

import "fmt"

func main() {
	//Open CSV File
f, err := os.Open("blah.csv")
if err != nil {
	log.Fatal(err)
}

//Read in the CSV records.
r := csv.NewReader(f)
records, err := r.ReadAll()
if err != nil {
	log.Fatal(err)
}

//Get the max val in int col
var intMax int
for _, record := range records {
	//Parse the int val
	intVal, err := strconv.Atoi(record[0])
	if err != nil {
		log.Fatal(err)
	}

	//Replace the max val if nec
	if intVal > intMax {
		intMax = intVal
	}
 }

 fmt.Println(intMax)
}


