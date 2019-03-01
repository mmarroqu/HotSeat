package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

type test_struct struct {
	Test string
}

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
    /*params := mux.Vars(r)
    var person string
    _ = json.NewDecoder(r.Body).Decode(&person)
    fmt.Println(params["name"])
	*/
	decoder := json.NewDecoder(r.ParseForm())
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
  	fmt.Println(t.Test)

}


func GetQuestion(w http.ResponseWriter, r *http.Request) {
   
}

func SendQuestion(w http.ResponseWriter, r *http.Request) {

}
