package services

import (
	"encoding/json"
	"fmt"

	"github.com/tc_test/models"
)

func SearchOrganizationAction(conf models.Config, field string, value interface{}) {
	var organizations []map[string]interface{}
	organizations = Search(conf.Data.Organization, field, value)

	for _, ogt := range organizations {
		var tickets []map[string]interface{}
		tickets = Search(conf.Data.Ticket, "organization_id", ogt["_id"])
		var subjects []interface{}
		for _, ticket := range tickets {
			subjects = append(subjects, ticket["subject"])
		}
		ogt["tickets"] = subjects

		var users []map[string]interface{}
		users = Search(conf.Data.User, "organization_id", ogt["_id"])
		var userNames []interface{}
		for _, user := range users {
			userNames = append(userNames, user["name"])
		}
		ogt["users"] = userNames
	}
	searchResult, _ := json.MarshalIndent(models.SearchResult{NumberOfResult: len(organizations), Data: organizations}, "", "  ")
	fmt.Println(string(searchResult))
}
