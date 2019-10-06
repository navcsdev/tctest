package models

type SearchResult struct {
	NumberOfResult int                      `json:"number_of_result"`
	Data           []map[string]interface{} `json:"search_result"`
}
