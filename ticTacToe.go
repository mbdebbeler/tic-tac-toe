package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetPrefix("main: ")
	log.SetFlags(0)

	message, err := startGame(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	present(message)
}

func askHuman(messageKey string, stdin io.Reader) (string, error) {
	present(messageKey)
	userInput, error := getInput(stdin)
	return userInput, error
}

func startGame(stdin io.Reader) (string, error) {
	present("welcome")
	result, error := askHuman("chooseMarker", stdin)
	return result, error
}

func present(messageKey string) {
	fmt.Println(messages[messageKey])
}

var messages = map[string]string{
	"welcome":      "Welcome to TicTacToe!",
	"chooseMarker": "Please hit X or O to choose your marker: ",
	"X":            "X",
	"O":            "O",
}

func readInput(stdin io.Reader) (string, error) {
	reader := bufio.NewReader(stdin)
	text, error := reader.ReadString('\n')
	return text, error
}

func getInput(stdin io.Reader) (string, error) {
	userInput, error := readInput(stdin)
	if error != nil {
		return "", error
	}
	trimmedInput := trimInput(userInput)
	return string(trimmedInput), nil
}

func trimInput(userInput string) string {
	return strings.TrimRight(userInput, "\r\n")
}
