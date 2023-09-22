package main

import "fmt"

func main() {

	var fruitList [4]string
	// var vegList = [3]string{"tomato","potato","cauli"} another way to create array

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is", fruitList)        // declare 4 element in frutiList so [Apple Banana  Mango] we have dobule space after banana
	fmt.Println("Fruit length is", len(fruitList)) // length is 4 coz of space
}

// Note : to seperate array and slice we can define a number of element [4]string it calls array not slice
// Now for slice we just var str = []string{}
// in array we cannot add as much value as we like coz number of element we declare in array [5]
// in slice we can expand as much we like we use a method called append()
