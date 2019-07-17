package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type templateVariables struct {
	BackendService string
}

func main() {

	http.HandleFunc("/", homePage)

	fmt.Println("Listening on port 9000") 

	err := http.ListenAndServe(":9000", nil)
    if err != nil {
        log.Fatal("Could not listen: ", err)
    }
}

func homePage(w http.ResponseWriter, r *http.Request) {

	backendSvc := os.Getenv("BACKEND_SERVICE")

	templateVars := templateVariables {
		BackendService:      backendSvc,
	}
	tmpl:= template.Must(template.ParseFiles("index.html"))

	err := tmpl.Execute(w, templateVars) 
	if err != nil {                    
		log.Print("error executing template: ", err) 
	}
}