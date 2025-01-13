package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var dictionary = []string{
	"Apple",
	"Adventure",
	"Calendar",
	"Universe",
	"Pyramid",
	"Journey",
	"Chocolate",
	"Sunshine",
	"Notebook",
	"Library",
	"Winter",
	"Planet",
	"Keyboard",
	"Ghost",
	"Ukraine",
}

var inputReader = bufio.NewReader(os.Stdin)
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for {
		playGame()

		if !askToPlayAgain() {
			fmt.Println("Thanks for playing! Goodbye!")
			break
		}
	}
}

func playGame() {
	gameWord := getRandomWord(rng)
	guessedLetters := initializeGuessedWords(gameWord)
	hangmanState := 0

	for !isGameOver(gameWord, guessedLetters, hangmanState) {
		printGameState(gameWord, guessedLetters, hangmanState)

		input := readInput()

		if len(input) != 1 {
			fmt.Println("Invalid input")
			continue
		}

		letter := unicode.ToLower(rune(input[0]))

		if isCorrectGuess(gameWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
	}

	fmt.Println("Game over!")

	if isWordGuessed(gameWord, guessedLetters) {
		fmt.Println("You Win!!!")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("You Lose!")
		fmt.Printf("The word was: %s\n", gameWord)
	} else {
		panic("Invalid state. Game is over without winner!")
	}
}

func askToPlayAgain() bool {
	for {
		fmt.Print("Do you want to play again? (y/n): ")

		input, err := inputReader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		input = strings.TrimSpace(strings.ToLower(input))

		if input == "y" || input == "yes" {
			return true
		} else if input == "n" || input == "no" {
			return false
		} else {
			fmt.Println("Invalid input, please enter 'y' or 'n'.")
		}
	}
}

func initializeGuessedWords(gameWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(gameWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(gameWord[len(gameWord)-1]))] = true

	return guessedLetters
}

func getRandomWord(rng *rand.Rand) string {
	word := dictionary[rng.Intn(len(dictionary))]
	return word
}

func isWordGuessed(gameWord string, guessedLetters map[rune]bool) bool {
	for _, character := range gameWord {
		if !guessedLetters[unicode.ToLower(character)] {
			return false
		}
	}

	return true
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
}

func isGameOver(gameWord string, guessedLetters map[rune]bool, hangmanState int) bool {
	return isWordGuessed(gameWord, guessedLetters) || isHangmanComplete(hangmanState)
}

func printGameState(gameWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuessingProgress(gameWord, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(hangmanState))
}

func getWordGuessingProgress(gameWord string, guessedLetters map[rune]bool) string {
	result := ""

	for _, character := range gameWord {
		if character == ' ' {
			result += "    "
			continue
		}

		if guessedLetters[unicode.ToLower(character)] {
			result += string(character)
		} else {
			result += " _ "
		}
	}

	return result
}

func getHangmanDrawing(hangmanState int) string {
	fileName := fmt.Sprintf("states/hangman%d", hangmanState)

	data, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func readInput() string {
	fmt.Print("> ")

	input, err := inputReader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

func isCorrectGuess(gameWord string, letter rune) bool {
	return strings.ContainsRune(gameWord, letter)
}
