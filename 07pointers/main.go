package main

import "fmt"

func main() {
	fmt.Println("Hello pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr) //nil

	myNumber := 50
	var ptr = &myNumber
	fmt.Println("& pointer is ", ptr)  //0x14000112008 memory location of myNumber
	fmt.Println("* pointer is ", *ptr) // gives memory location of value which is 50

	*ptr = *ptr + 50
	fmt.Println("Final value of pointer", myNumber) // 100 because action + perform on that memory location value not copy of that value

}
