# Hangman Game

Welcome to the Hangman Game! This is a console-based word-guessing game where you try to guess the hidden word one letter at a time. The game is written in Go and features a randomized word selection from a predefined dictionary. Enjoy guessing the word before the hangman drawing is completed!

## How to Play

1. A random word is selected from the dictionary.

2. The first and last letters of the word are revealed.

3. You guess letters one at a time by typing them in.

4. If you guess correctly, the letter is revealed in the word.

5. If you guess incorrectly, the hangman drawing progresses.

6. The game ends when:

    * You guess the entire word correctly (You Win!).

    * The hangman drawing is completed after 9 incorrect guesses (You Lose!).

## Installation and Setup

To run the game locally, follow these steps:

### Prerequisites

* Go installed on your system (version 1.20 or later).

* A terminal or command prompt to run the game.

### Steps

1. Clone the repository:

```
git clone https://github.com/ruslan-symonenko/hangman-game.git
```

2. Navigate to the project directory:

```
cd hangman-game
```

3. Build the game:

```
go build -o hangman
```

4. Run the game:

```
./hangman
```

## Customization

You can modify the word dictionary by editing the dictionary variable in the main.go file.

Update or replace the files in the states folder to customize the hangman drawings.