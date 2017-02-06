package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func checkForError(e error, msg ...string) {
	if e != nil {
		if len(msg) == 0 {

			panic(e.Error())
		}
		panic(errors.New(e.Error() + msg[0]))
	}
}

func main() {
	var conf = LoadConfig()
	//setting handler
	http.HandleFunc("/", hookHandler)

	address := conf.Address + ":" + string(conf.Port)

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
