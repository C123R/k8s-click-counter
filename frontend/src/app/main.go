package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"io/ioutil"
	"os"
)

type templateVariables struct {
	PageTitle string
}

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/count", getCount)
	fmt.Println("Listening on port 9000") 

	err := http.ListenAndServe(":9000", nil)
    if err != nil {
        log.Fatal("Could not listen: ", err)
    }
}

func homePage(w http.ResponseWriter, r *http.Request) {

	templateVars := templateVariables {
		PageTitle:      "Click-Counter",
	}
	tmpl:= template.Must(template.ParseFiles("index.html"))

	err := tmpl.Execute(w, templateVars) 
	if err != nil {                    
		log.Print("error executing template: ", err) 
	}
}

func getCount(w http.ResponseWriter, r *http.Request) {

	backendSvc := os.Getenv("BACKEND_SERVICE")

    response, err := http.Get(backendSvc)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
 
    count, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
 
    clickCount := string(count)
    fmt.Fprint(w, clickCount)
}