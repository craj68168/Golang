package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	performJson()
}

func performJson() {
	const myUrl string = "https://localhost:8080/api/post"
	requestData := strings.NewReader(`
	{"name":"raj","age":26}
	`)

	response, err := http.Post(myUrl, "applicaiton/json", requestData)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("content", string(content))
}
