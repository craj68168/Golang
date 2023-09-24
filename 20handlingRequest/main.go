package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://www.google.com"

func main() {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response types is %T\n", response) //*http.Response
	defer response.Body.Close()                    //close the connection or it will request again and again

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	res := string(content)
	fmt.Println("The content is", res)
}
