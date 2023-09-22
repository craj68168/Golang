package main

import "fmt"

func main() {

	var fruitList = []string{"Apple", "Tomato"} // we donot declare number of element in array

	fmt.Printf("Type of fruitList %T\n", fruitList) // []string

	fruitList = append(fruitList, "Banana", "Mango")
	fmt.Println("Frutis value ", fruitList) //[Apple Tomato Banana Mango]
}
