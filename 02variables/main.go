package main

import "fmt"

const LoginToken string = "mytoken" // capital L means we make it public and it constant and use it inside method as it is globally declare
func main() {
	var username = "Raj"
	fmt.Println(username)
	fmt.Printf("Variable username is type: %T \n", username) // string

	var isBool bool = false
	fmt.Println(isBool)
	fmt.Printf("Variable isBool is type: %T \n", isBool) // bool

	var smallVal uint8 = 255 // limit upto 0-255 max
	fmt.Println(smallVal)
	fmt.Printf("Variable smallVal is type: %T \n", smallVal) //uint8

	var smallFloat float32 = 255.554433
	fmt.Println(smallFloat)
	fmt.Printf("Variable smallFloat is type: %T \n", smallFloat) //255.55443 only 5 value after decimal

	var smallFloat2 float64 = 255.554433221100
	fmt.Println(smallFloat2)
	fmt.Printf("Variable smallFloat2 is type: %T \n", smallFloat2) //it gives 255.5544332211 only

	// default values and some aliases
	var defaultValInt int
	fmt.Println(defaultValInt)                                         // default value will be 0
	fmt.Printf("Variable defaultValInt is type: %T \n", defaultValInt) // int

	var defaultValString string
	fmt.Println(defaultValString)                                            //  nothing
	fmt.Printf("Variable defaultValString is type: %T \n", defaultValString) // string

	// implicit way of declaring variables
	var myString = "RaZz"
	fmt.Println(myString)                                    //  print it
	fmt.Printf("Variable myString is type: %T \n", myString) // decalre to string
	// cannot do myString = 3

	// no var style
	myNumber := 123456.90 // inside any method you are allowed to use this walrus operator not outside of any method.
	fmt.Println("number", myNumber)

	fmt.Println("LoginToken", LoginToken)
}
