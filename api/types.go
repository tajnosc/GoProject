package api_types

type Question struct {
	Query    string   `json:"query"`
	Variants map[string]string `json:"variants"`
}

type Questions = map[string]Question

type Answers = map[string]string

type SubmitResponse struct {
	CorrectAnswers   int     `json:"correct"`
	IncorrectAnswers int     `json:"incorrect"`
	Position         float64 `json:"position"`
}