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
	//adds user names
	TestPostName("Manuel")
	TestPostName("Joey")
	TestPostName("Silvia")
	TestPostName("Marroquin")
	TestPostName("DiMaria")
	TestPostName("Camara")

	//adds two questions
	TestPostQuestion("Test Question 1?")
	TestPostQuestion("Test Question 2?")
	TestPostQuestion("Test Question 3?")
	TestPostQuestion("Test Question 4?")
	TestPostQuestion("Test Question 5?")
	TestPostQuestion("Test Question 6?")
	//gets the two questions and tries one extra time
	fmt.Println("Getting Question...")
	TestGetQuestion()
	fmt.Println("Getting Question...")
	TestGetQuestion()
	fmt.Println("Getting Question...")	
	TestGetQuestion()
	fmt.Println("Getting Question...")
	TestGetQuestion()
	fmt.Println("Getting Question...")
	TestGetQuestion()
	fmt.Println("Getting Question...")
	TestGetQuestion()
	fmt.Println("Getting Question...")
	TestGetQuestion()
	
}

//tests the post request for sending a user's name
func TestPostName(name string) {
	
	formData := url.Values{
		"name": {name},
	}

	resp, err := http.PostForm("http://localhost:8000/start", formData)
	if err != nil {
		log.Fatalln(err)
	}else{
		fmt.Println("Sent Name: " + name)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	//log.Println(result["form"])
}

//tests the post request for submitting a question
func TestPostQuestion(question string) {
	formData := url.Values{
		"question": {question},
	}

	resp, err := http.PostForm("http://localhost:8000/game", formData)
	if err != nil {
		log.Fatalln(err)
	}else{
		fmt.Println("Sent Question: " + question)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	//log.Println(result["form"])
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