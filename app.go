package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/tc_test/describes"
	"github.com/tc_test/models"
	"github.com/tc_test/services"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func Info() {
	app.Name = "Simple Search"
	app.Usage = "CLI for search data"
	app.Author = "Navcs"
	app.Version = "1.0.0"
}

func ParseQuery(query string) (string, string) {
	queryArr := strings.Split(query, "=")
	return queryArr[0], queryArr[1]
}

func Commands() {
	app.Commands = []cli.Command{
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "Search data",
			Subcommands: []cli.Command{
				{
					Name:  "organizations",
					Usage: "search data organizations",
					Action: func(c *cli.Context) error {
						query := c.Args().Get(0)
						if len(query) == 0 {
							return cli.NewExitError("Missing query arguments. Example correct: search organizations ':tags=West'", 86)
						}
						field, value := ParseQuery(query)
						services.SearchOrganizationAction(conf, field, value)
						return nil
					},
				},
				{
					Name:  "users",
					Usage: "search data users",
					Action: func(c *cli.Context) error {
						query := c.Args().Get(0)
						if len(query) == 0 {
							return cli.NewExitError("Missing query arguments. Example correct: search users 'alias=Miss Coffey'", 86)
						}
						field, value := ParseQuery(query)
						services.SearchUserAction(conf, field, value)
						return nil
					},
				},
				{
					Name:  "tickets",
					Usage: "search data tickets",
					Action: func(c *cli.Context) error {
						query := c.Args().Get(0)
						if len(query) == 0 {
							return cli.NewExitError("Missing query arguments. Example correct: search tickets 'status=pending'", 86)
						}
						field, value := ParseQuery(query)
						services.SearchTicketAction(conf, field, value)
						return nil
					},
				},
			},
		},
		{
			Name:    "describe",
			Aliases: []string{"s"},
			Usage:   "get all search fields",
			Subcommands: []cli.Command{
				{
					Name:  "organizations",
					Usage: "search fields for organizations",
					Action: func(c *cli.Context) error {
						fmt.Println(describes.OrganizationFields())
						return nil
					},
				},
				{
					Name:  "users",
					Usage: "search fields for users",
					Action: func(c *cli.Context) error {
						fmt.Println(describes.UserFields())
						return nil
					},
				},
				{
					Name:  "tickets",
					Usage: "search fields for tickets",
					Action: func(c *cli.Context) error {
						fmt.Println(describes.TicketFields())
						return nil
					},
				},
			},
		},
	}
}

var conf models.Config

func main() {
	if _, err := toml.DecodeFile("./config/default.toml", &conf); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	Info()
	Commands()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
