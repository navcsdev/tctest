package models

type Config struct {
	Data Data
}

type Data struct {
	Organization string
	User         string
	Ticket       string
}
