package services

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Search(fileName string, field string, value interface{}) []map[string]interface{} {
	var results []map[string]interface{}
	var unique bool
	if field == "_id" {
		unique = true
	}
	if value == nil {
		return results
	}

	file, _ := os.Open(fileName)
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Token()

	for decoder.More() {
		data := map[string]interface{}{}
		decoder.Decode(&data)
		if strings.Contains(fmt.Sprintf("%v", data[field]), fmt.Sprintf("%v", value)) {
			results = append(results, data)
			if unique {
				break
			}
		}
	}
	return results
}
