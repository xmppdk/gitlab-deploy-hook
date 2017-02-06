package main

import (
	"encoding/json"
	"os"
)

// Config struct to carry configuration items
type Config struct {
	Logfile        string
	Address        string
	Port           int64
	RepositoryName string
	Command        string
	Token          string
}

// LoadConfig loads the configuration file and returns a config struct
func LoadConfig() Config {
	var conf Config
	var file, e = os.Open("config.json")
	checkForError(e)

	// close file on exit and check for its returned error
	defer func() {
		e := file.Close()
		checkForError(e)
	}()

	buffer := make([]byte, 1024)
	count := 0

	count, e = file.Read(buffer)
	checkForError(e)

	e = json.Unmarshal(buffer[:count], &conf)
	checkForError(e)

	return conf
}
