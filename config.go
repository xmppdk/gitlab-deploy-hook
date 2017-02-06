package main

type configRepo struct {
	Name    string
	Command string
}

type config struct {
	Logfile    string
	Address    string
	Port       int64
	Repository configRepo
}
