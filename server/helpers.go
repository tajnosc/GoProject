package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/tajnosc/GoProject/api"
)

func checkUserAnswers(userAnswers api_types.Answers) api_types.SubmitResponse {
	var result = api_types.SubmitResponse{
		CorrectAnswers:   0,
		IncorrectAnswers: 0,
		Position:         0.0,
	}

	for key, val := range answers {
		if userAnswers[key] == val {
			result.CorrectAnswers++
		} else {
			result.IncorrectAnswers++
		}
	}

	saveResultAndGetPosition(&result)

	return result
}

func saveResultAndGetPosition(result *api_types.SubmitResponse) {
	var noOfQuestions = float64(len(answers))
	var correctAnswers = float64(result.CorrectAnswers)
	var userResult = correctAnswers / noOfQuestions

	var userPosition = float64(saveResult(userResult))
	var totalUsers = float64(len(ranks) - 1)

	if totalUsers > 0 {
		result.Position = (userPosition / totalUsers) * 100
	} else {
		// Don't divide by zero!
		result.Position = 100
	}
}

/**
 * Insert a new result value into a sorted array of current ranks
 * and saves the new value to a file
 */
func saveResult(s float64) int {
	i := sort.SearchFloat64s(ranks, s)
	ranks = append(ranks, 0)
	copy(ranks[i+1:], ranks[i:])
	ranks[i] = s
	saveRanks()
	return i
}

func saveRanks() {
	file, err := os.Create(ranksPATH)
	if err != nil {
		return
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, line := range ranks {
		fmt.Fprintln(w, line)
	}
	w.Flush()
}
