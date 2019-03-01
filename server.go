package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)



var players [] string
var questions [] string
var numQuestions = 0
// our main function
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/start", PostPlayer).Methods("POST")
    router.HandleFunc("/game", GetQuestion).Methods("GET")
    router.HandleFunc("/game", SendQuestion).Methods("POST")

    log.Fatal(http.ListenAndServe(":8000", router))
}


func PostPlayer(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println("Name received:" + r.Form["name"][0])
    person := r.Form["name"][0]

    players=append(players,person)
    fmt.Print("Current players: ")
    fmt.Println(players)
   
	
}


func GetQuestion(w http.ResponseWriter, r *http.Request) {
	if(numQuestions>0){
    	q := questions[len(questions)-1]
    	questions = questions[:len(questions)-1]
    	json.NewEncoder(w).Encode(q)
    	numQuestions--

    }else{
    	json.NewEncoder(w).Encode("No Questions Left")
    }

}


func SendQuestion(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
    fmt.Println("Question received:" + r.Form["question"][0])
    question := r.Form["question"][0]

    questions = append(questions,question)
    numQuestions++
    fmt.Print("Current Questions: ")
    fmt.Println(questions)
}
