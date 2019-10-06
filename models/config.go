package models

type Config struct {
	Data data
}

type data struct {
	Organization string
	User         string
	Ticket       string
}
