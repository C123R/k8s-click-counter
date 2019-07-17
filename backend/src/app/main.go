// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var counter int

func main() {
	counter = 0
	http.HandleFunc("/api/counter", countIt)

	fmt.Println("Listening on port 9001") 

	err := http.ListenAndServe(":9001", nil)
    if err != nil {
        log.Fatal("Could not listen: ", err)
    }
}


func countIt(w http.ResponseWriter, r *http.Request) {
	messages := make(chan string)
	enableCors(&w)
	go func() {
		counter++
		messages <- strconv.Itoa(counter)
	}()
	myMessage := <-messages
	fmt.Fprintln(w, myMessage)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
