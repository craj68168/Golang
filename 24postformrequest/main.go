package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	postForm()
}

func postForm() {
	const myUrl string = "https://localhost:8080/api/postform"

	data := url.Values{}

	data.Add("firstName", "raj")
	data.Add("lastName", "chaudhary")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("content", string(content))
}
