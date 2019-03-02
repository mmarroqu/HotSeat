package main

import (
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func main() {
	//adds user name
	TestPostName()
	//adds two questions
	TestPostQuestion()
	TestPostQuestion()
	//gets the two questions and tries one extra time
	TestGetQuestion()
	TestGetQuestion()
	TestGetQuestion()
}

//tests the post request for sending a user's name
func TestPostName() {
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
}

//tests the post request for submitting a question
func TestPostQuestion() {
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

//tests the get request for getting a question
func TestGetQuestion() {
	response, err := http.Get("http://localhost:8000/game")
    if err != nil {
		log.Fatalln(err) 
	} else {
   		defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
			log.Fatalln(err)
		}
        fmt.Printf("%s", string(contents))
    }
}