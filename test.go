package main

import (
	"net/http"
	"net/url"
	"log"
	//"io/ioutil"
	"encoding/json"
	//"bytes"
	//"fmt"
)

func main() {
	MakeRequest()
}

func MakeRequest() {
/*
	name := "manuel"
	formData := url.Values{
		"name": {name},
	}

	resp, err := http.PostForm("http://localhost:8000/start", formData)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result["form"])
*/

	question := "what is your iq?"
	formData := url.Values{
		"question": {question},
	}

	resp, err := http.PostForm("http://localhost:8000/game", formData)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result["form"])
}