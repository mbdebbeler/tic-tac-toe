package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	msg, err := Welcome()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}

func askHuman() string {
	fmt.Print("\nPlease hit X or O to choose your marker: ")
	var userInput string
	fmt.Scan(&userInput)
	return userInput
}

func Welcome() (string, error) {
	message := fmt.Sprintf("Welcome to TicTacToe!")
	fmt.Println(askHuman())
	return message, nil
}
