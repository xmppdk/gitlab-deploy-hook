package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var conf Config

func main() {
	conf = LoadConfig()

	http.HandleFunc("/", hookHandler)

	address := conf.Address + ":" + strconv.FormatInt(conf.Port, 10)

	log.Println("Listening on " + address)

	e := http.ListenAndServe(address, nil)
	if e != nil {
		log.Println(e)
	}
}

func checkForError(e error, msg ...string) {
	if e != nil {
		if len(msg) == 0 {

			panic(e.Error())
		}
		panic(errors.New(e.Error() + msg[0]))
	}
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	// Check incoming requests for proper access token
	e := verifyToken(r)
	checkForError(e)

	var hook webhook
	var data, _ = ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &hook)
}

func verifyToken(r *http.Request) error {
	var token = r.Header.Get("X-GitLab-Token")
	if token != conf.Token {
		return errors.New("invalid token received")
	}
	return nil
}
