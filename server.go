package main

import (
    "encoding/json"
    "log"
    "net/http"
    "net/url"
    "github.com/gorilla/mux"
    "fmt"
    "math/rand"
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
   
    //playerIndex := rand.Intn(len(players))
    //fmt.Println(questionIndex)
	if(numQuestions>0){
        // Get random player and question
        questionIndex := rand.Intn(len(questions))
        playerIndex := rand.Intn(len(players))
    	q := questions[questionIndex]
        playerName := players[playerIndex]

        // Remove Question from list of questions
        questions[questionIndex] = questions[len(questions)-1]
    	questions = questions[:len(questions)-1]
        formData := url.Values{
            "question": {q},
            "name": {playerName},
         }

    	json.NewEncoder(w).Encode(formData)
    	numQuestions--

    }else{
    	json.NewEncoder(w).Encode(nil)
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
