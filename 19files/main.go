package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is testing"
	file, error := os.Create("./myfiles.txt")

	if error != nil {
		panic(error) // panic will shut down the execution and shows error
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("The length is ", length) // 15 length and create a file and append the content
	defer file.Close()
	ReadFiles("./myfiles.txt")
}

func ReadFiles(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}
	fmt.Println("The date is", string(databyte))
}
