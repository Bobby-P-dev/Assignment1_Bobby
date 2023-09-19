package main

import (
	"assignment1/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func cleanedText(text string) string {
	return strings.ToLower(strings.Replace(text, "-", "", -1))
}

func isStudentExist(userInput, studentCode string) bool {
	userinputCleaned := cleanedText(userInput)
	studentCodeCleaned := cleanedText(studentCode)

	return userinputCleaned == studentCodeCleaned
}

func ShowStudent(userInput string) {
	jsonPath := filepath.Join("data", "data.json")

	data, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		log.Fatal(err)
	}

	var students model.Students

	if err := json.Unmarshal(data, &students); err != nil {
		log.Fatal(err)
	}

	if userInput == "" {
		fmt.Println("input student code")
		return
	}

	for _, student := range students.Students {
		if isStudentExist(userInput, student.Code) {
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Student Code: %s\n", student.Code)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Address: %s\n", student.Address)
			fmt.Printf("Occupation: %s\n", student.Occupation)
			fmt.Printf("Joining Reason: %s\n", student.JoiningReason)
			return
		}
	}

	fmt.Printf("The student'%s' was not found. \n", userInput)
}

func getUserInput() string {
	args := os.Args

	if len(args) < 2 {
		return ""
	}

	return args[1]
}

func main() {
	userInput := getUserInput()

	ShowStudent(userInput)
}
