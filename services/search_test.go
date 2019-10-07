package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tc_test/models"
)

func TestFuncSearch(t *testing.T) {
	var conf models.Config
	conf = models.Config{
		models.Data{
			Organization: "../data/organizations.json",
			User:         "../data/users.json",
			Ticket:       "../data/tickets.json",
		},
	}
	users := Search(conf.Data.User, "_id", 50)
	if len(users) != 1 {
		t.Errorf("Search user by id failed, results %v", len(users))
	}
	assert.Equal(t, users[0]["_id"], float64(50))
	assert.Equal(t, users[0]["external_id"], "e1378651-f998-4181-8b1b-35e99a30b900")
	assert.Equal(t, users[0]["tags"], []interface{}{"Kaka", "Abrams", "Genoa", "Yettem"})
}
