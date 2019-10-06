package services

import (
	"encoding/json"
	"fmt"

	"github.com/tc_test/models"
)

func SearchUserAction(conf models.Config, field string, value interface{}) {
	var users []map[string]interface{}
	users = Search(conf.Data.User, field, value)

	for _, user := range users {
		var organizations []map[string]interface{}
		organizations = Search(conf.Data.Organization, "_id", user["organization_id"])
		if len(organizations) != 0 {
			organizationName := organizations[0]["name"]
			user["organization_name"] = organizationName
		}
		var assigneeTickets []map[string]interface{}
		var submittedTickets []map[string]interface{}
		assigneeTickets = Search(conf.Data.Ticket, "assignee_id", user["_id"])
		if len(assigneeTickets) != 0 {
			var subjects []interface{}
			for _, assigneeTicket := range assigneeTickets {
				subjects = append(subjects, assigneeTicket["subject"])
			}
			user["assigned_tickets"] = subjects
		}
		submittedTickets = Search(conf.Data.Ticket, "submitter_id", user["_id"])
		if len(submittedTickets) != 0 {
			var subjects []interface{}
			for _, submittedTicket := range submittedTickets {
				subjects = append(subjects, submittedTicket["subject"])
			}
			user["submitted_tickets"] = subjects
		}
	}
	searchResult, _ := json.MarshalIndent(models.SearchResult{NumberOfResult: len(users), Data: users}, "", "  ")
	fmt.Println(string(searchResult))
}
