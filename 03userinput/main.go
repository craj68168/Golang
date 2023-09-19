package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fullName := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter Your Name:")

	input, err := fullName.ReadString('\n') // \n means hitting enter

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Your Name is", input)

}
