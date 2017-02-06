package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//setting handler
	http.HandleFunc("/", hookHandler)

	address := "0.0.0.0:3344"

	log.Println("Listening on " + address)

	//starting server
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Println(err)
	}
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	var hook webhook

	var data, _ = ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &hook)
	log.Println(fmt.Sprintf("Repository URL: %s", hook.Repository.Url))
}
