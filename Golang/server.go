package main

import (
	"fmt"
	"html/template"
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func render(w http.ResponseWriter, tmpl string) {
    tmpl = fmt.Sprintf("templates/%s", tmpl)
    t, err := template.ParseFiles(tmpl)
    if err != nil {
        log.Print("template parsing error: ", err)
    }
    err = t.Execute(w, "")
    if err != nil {
        log.Print("template executing error: ", err)
    }
}

func Home(w http.ResponseWriter, req *http.Request) {
    render(w, "index.html")
}

func Game(w http.ResponseWriter, req *http.Request) {
    render(w, "game.html")
}

func main() {
//	flag.Parse()
//	log.Printf("listening on %q...", *listen)
	http.HandleFunc("/", Home)
  http.HandleFunc("/game/", Game)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
			log.Fatal("ListenAndServe: ", err)
	}
//	log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir))))
}
