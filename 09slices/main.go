package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Apple", "Tomato"} // we donot declare number of element in array

	fmt.Printf("Type of fruitList %T\n", fruitList) // []string

	fruitList = append(fruitList, "Banana", "Mango")
	fmt.Println("Appended Fruits value are ", fruitList) //[Apple Tomato Banana Mango]

	// fruitList = append(fruitList[1:]) // [Tomato Banana Mango] start from 1 length and endling endless
	// fruitList = append(fruitList[:3]) // [Apple Tomato Banana ] start from 0 length and not count 3 length
	// fruitList = append(fruitList[:3]) // start from 1 and not include 3 [Tomato Banana]
	// fmt.Println(fruitList)

	highScores := make([]int, 4) // array cause we give how much element to store but its not array it slice

	highScores[0] = 234
	highScores[1] = 145
	highScores[2] = 321
	highScores[3] = 876
	// fmt.Println("High scores are", highScores) // [234 145 321 876]
	// if highScores[4] = 123 than error cause it has only 4 element to store
	highScores = append(highScores, 123, 432)

	fmt.Println("High scores are", highScores) // [234 145 321 876,123,432] it can be extended through append()

	sort.Ints(highScores)
	fmt.Println(highScores)                     // [123 145 234 321 432 876] short the value
	fmt.Println(sort.IntsAreSorted(highScores)) // true
}
