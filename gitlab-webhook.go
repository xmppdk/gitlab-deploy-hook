package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

var conf Config

func main() {
	conf = LoadConfig()
	if conf.LogToFile == true {
		writer, e := os.OpenFile(conf.Logfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		checkForError(e)

		defer func() {
			writer.Close()
		}()

		log.SetOutput(writer)
	}
	setupWebServer()
}

func checkForError(e error, msg ...string) {
	if e != nil {
		if len(msg) == 0 {

			panic(e.Error())
		}
		panic(errors.New(e.Error() + msg[0]))
	}
}

func setupWebServer() {
	http.HandleFunc("/", hookHandler)
	address := conf.Address + ":" + strconv.FormatInt(conf.Port, 10)

	log.Println("Listening on " + address)

	e := http.ListenAndServe(address, nil)
	if e != nil {
		log.Println(e)
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
	verifyRepositoryName(&hook)

	go runCommand()
}

func verifyToken(r *http.Request) error {
	var token = r.Header.Get("X-GitLab-Token")
	if token != conf.Token {
		return errors.New("Invalid token received")
	}
	return nil
}

func verifyRepositoryName(hook *webhook) error {
	if hook.Repository.Name != conf.RepositoryName {
		return errors.New("Repository name does not match configured setting")
	}
	return nil
}

func runCommand() error {
	var cmd = conf.Command
	var command = exec.Command(cmd)
	out, err := command.Output()

	if err != nil {
		return err
	}

	log.Println("Executed: " + cmd)
	log.Println("Output: " + string(out))

	return nil
}
