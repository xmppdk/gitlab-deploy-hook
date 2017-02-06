package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type webhook struct {
	Before, After, Ref, User_name string
	User_id, Project_id           int
	Repository                    gitlabRepository
	Commits                       []commit
	Total_commits_count           int
}

type commit struct {
	Id, Message, Timestamp, Url string
	Author                      author
}

type author struct {
	Name, Email string
}

type gitlabRepository struct {
	Name, Url, Description, Home string
}

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
