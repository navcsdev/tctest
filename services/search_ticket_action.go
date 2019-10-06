package services

import (
	"encoding/json"
	"fmt"

	"github.com/tc_test/models"
)

func SearchTicketAction(conf models.Config, field string, value interface{}) {
	var tickets []map[string]interface{}
	tickets = Search(conf.Data.Ticket, field, value)

	for _, ticket := range tickets {
		assignees := Search(conf.Data.User, "_id", ticket["assignee_id"])
		if len(assignees) > 0 {
			ticket["assignee_name"] = assignees[0]["name"]
		}
		submitters := Search(conf.Data.User, "_id", ticket["submitter_id"])
		if len(submitters) > 0 {
			ticket["submitter_name"] = submitters[0]["name"]
		}

		organizations := Search(conf.Data.Organization, "_id", ticket["organization_id"])
		if len(organizations) > 0 {
			ticket["organization_name"] = organizations[0]["name"]
		}
	}
	searchResult, _ := json.MarshalIndent(models.SearchResult{NumberOfResult: len(tickets), Data: tickets}, "", "  ")
	fmt.Println(string(searchResult))
}
