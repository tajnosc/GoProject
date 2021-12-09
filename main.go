package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tajnosc/GoProject/api"
)

const queriesPATH = "assets/queries.json"
const answersPATH = "assets/answers.json"
const ranksPATH = "assets/ranks.txt"

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
