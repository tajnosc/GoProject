package main

import (
	"bufio"
	"net/http"
	"os"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/tajnosc/GoProject/api"
)

func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
}

func submit(c *gin.Context) {
	var userAnswers api_types.Answers

	if err := c.BindJSON(&userAnswers); err != nil {
		return
	}

	var userResult api_types.SubmitResponse = checkUserAnswers(userAnswers)

	c.IndentedJSON(http.StatusCreated, userResult)
}

func checkUserAnswers(userAnswers api_types.Answers) api_types.SubmitResponse {
	var answersValidation = api_types.SubmitResponse{
		CorrectAnswers:   0,
		IncorrectAnswers: 0,
		Position:         0.0,
	}

	for key, val := range answers {
		if userAnswers[key] == val {
			answersValidation.CorrectAnswers++
		} else {
			answersValidation.IncorrectAnswers++
		}
	}

	saveResultAndGetPosition(&answersValidation)

	return answersValidation
}

func saveResultAndGetPosition(answersValidation *api_types.SubmitResponse) {
	var noOfQuestions = float64(len(answers))
	var correctAnswers = float64(answersValidation.CorrectAnswers)

	var newResult = noOfQuestions / correctAnswers

	var userPosition = float64(InsertNewResult(newResult))
	var totalUsers = float64(len(ranks) - 1)

	if totalUsers > 0 {
		answersValidation.Position = (userPosition / totalUsers) * 100
	} else {
		// Don't divide by zero!
		answersValidation.Position = 100
	}
}

/**
 * Insert a new result value into a sorted array of current ranks
 * and saves the new value to a file
 */
func InsertNewResult(s float64) int {
	i := sort.SearchFloat64s(ranks, s)
	ranks = append(ranks, 0)
	copy(ranks[i+1:], ranks[i:])
	ranks[i] = s
	saveRanks()
	return i
}

func saveRanks() {
	file, err := os.Create("rank.json")
	if err != nil {
		return
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	w.Flush()
}
