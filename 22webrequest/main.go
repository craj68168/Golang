package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	MyGetrequest()
}

func MyGetrequest() {
	const myUrl string = "https://localhost:8080/api/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content is", string(content))

}
