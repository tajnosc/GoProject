package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tajnosc/GoProject/api"
)

const queriesPATH = "queries.json"
const answersPATH = "answers.json"
const ranksPATH = "ranks.txt"

var questions = api_types.Questions{}
var answers = api_types.Answers{}
var ranks = []float64{}

func main() {
	loadFiles()

	router := gin.Default()

	router.GET("/getQuestions", getQuestions)
	router.POST("/submit", submit)

	router.Run("localhost:8080")
}
