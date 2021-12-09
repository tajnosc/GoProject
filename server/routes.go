package main

import (
	"net/http"

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
