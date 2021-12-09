package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/manifoldco/promptui"
	"github.com/tajnosc/GoProject/server/api"

	"github.com/spf13/cobra"
)

const getQuestionsURL = "http://localhost:8080/getQuestions"
const submitURL = "http://localhost:8080/submit"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Starts quiz",
	Long:  `Starts quiz`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	var questions = getQuestions()
	var answers = startQuiz(questions)
	var results = submitQuiz(answers)
	printResults(results)
}

func getQuestions() api_types.Questions {
	var questions api_types.Questions

	if err := json.Unmarshal(getData(getQuestionsURL), &questions); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	return questions
}

func getData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)

	if err != nil {
		log.Printf("Could not send request: %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "dummy")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}

func startQuiz(questions api_types.Questions) api_types.Answers {
	var answers = api_types.Answers{}

	for number, q := range questions {
		fmt.Println()

		optionsKeys := make([]string, 0, len(q.Variants))
		optionsValues := make([]string, 0, len(q.Variants))

		for key, value := range q.Variants {
			optionsKeys = append(optionsKeys, key)
			optionsValues = append(optionsValues, value)
		}

		prompt := promptui.Select{
			Label: q.Query,
			Items: optionsValues,
		}

		index, _, _ := prompt.Run()
		answers[number] = optionsKeys[index]
	}

	return answers
}

func submitQuiz(answers api_types.Answers) api_types.SubmitResponse {
	var response = submit(submitURL, answers)
	var result api_types.SubmitResponse
	json.Unmarshal(response, &result)
	return result
}

func submit(baseAPI string, values api_types.Answers) []byte {
	json_data, _ := json.Marshal(values)

	request, err := http.NewRequest(
		http.MethodPost,
		baseAPI,
		bytes.NewBuffer(json_data),
	)

	if err != nil {
		log.Printf("Could not send request: %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "dummy")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}

func printResults(result api_types.SubmitResponse) {
	fmt.Println("======================")
	fmt.Println("Your results:")
	fmt.Println("Correct answers: ", result.CorrectAnswers)
	fmt.Println("Incorrect answers: ", result.IncorrectAnswers)
	fmt.Println("Your score is better than: ", result.Position, "% of contestants!")
}

func init() {
	rootCmd.AddCommand(initCmd)
}
