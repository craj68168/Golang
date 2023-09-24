package main

import "fmt"

func main() {

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for d := 0; d < len(days); d++ { // we dont generally use this
	// 	fmt.Println("Loops", days[d])
	// }

	// for i := range days {
	// 	fmt.Println("Range", days[i])
	// }

	// for i, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", i, day)
	// }

	for _, day := range days {
		fmt.Printf("Index is and value is %v\n", day)
	}

	myValue := 1

	for myValue < 10 {

		if myValue == 3 {
			goto lco
		}

		if myValue == 5 {
			break
			// myValue++
			// continue
		}
		fmt.Println("Print", myValue)
		myValue++
	}

lco:
	fmt.Println("Its jumping goto")

}
