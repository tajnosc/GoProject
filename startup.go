package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func readJsonFromFile(filePATH string, outStruct interface{}) {
	jsonFile, err := os.Open(queriesPATH)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &outStruct)
}

func loadQuestions() {
	readJsonFromFile(queriesPATH, &questions)
}

func loadAnswers() {
	readJsonFromFile(answersPATH, &answers)
}

func loadRanks() {
	file, err := os.Open(ranksPATH)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s, _ := strconv.ParseFloat(scanner.Text(), 64)
		ranks = append(ranks, s)
	}
}

func loadFiles() {
	loadQuestions()
	loadAnswers()
	loadRanks()
}
