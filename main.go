package main

import "server"

type card struct {
	id          int
	name        string
	category    string
	element     string
	atack_pwr   int
	defence_pwr int
	description string
	quote       string
}

func main() {
	s := server.NewServer()
	s.Run()
}
