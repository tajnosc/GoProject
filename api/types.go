package api_types

type Question struct {
	Query    string   `json:"query"`
	Variants []string `json:"variants"`
}

type Questions = map[string]Question

type Answers = map[string]string