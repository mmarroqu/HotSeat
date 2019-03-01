package main

import (
   // "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)



var players [] string
var questions [] string

// our main function
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/start", PostPlayer).Methods("POST")
    router.HandleFunc("/game/", GetQuestion).Methods("GET")
    router.HandleFunc("/game/", SendQuestion).Methods("POST")

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
   
}

func SendQuestion(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
    fmt.Println("Question received:" + r.Form["question"][0])
    question := r.Form["question"][0]

    questions = append(questions,question)
    fmt.Print("Current Questions: ")
    fmt.Println(questions)
}
