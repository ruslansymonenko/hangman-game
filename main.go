package main

import (
	"fmt"
	"math/rand"
	"os"
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

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	gameWord := getRandomWord(rng)
	guessedLetters := initializeGuessedWords(gameWord)

	printGameState(gameWord, guessedLetters)
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

func printGameState(gameWord string, guessedLetters map[rune]bool) {
	for _, character := range gameWord {
		if character == ' ' {
			fmt.Print("    ")
			continue
		}

		if guessedLetters[unicode.ToLower(character)] {
			fmt.Print(string(character))
		} else {
			fmt.Print(" _ ")
		}
	}

	fmt.Println()
}

func printHangman(hangmanState int) string {
	data, err := os.ReadFile("states/hangman6")

	if err != nil {
		panic(err)
	}

	return string(data)
}
