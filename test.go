package main

import (
	"net/http"
	//"net/url"
	"log"
	//"io/ioutil"
	"encoding/json"
	"bytes"
	"fmt"
)

func main() {
	MakeRequest()
}

func MakeRequest() {

	formData := url.Values{
		"name": {"masnun"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result["form"])
/*
	message := map[string]interface{}{
		"name": "poopyhead",
		"hello": "world",
		"life":  42,
		"embedded": map[string]string{
			"yes": "of course!",
		},
	}
	fmt.Println(message)
	fmt.Println("")

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(bytesRepresentation)
	fmt.Println("")

	resp, err := http.Post("http://localhost:8000/start", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
	fmt.Println("")

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])
*/
}