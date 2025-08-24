// Wordle game prep (step 9) - introduces file I/O for a larger word list
// Step 9 Uses os.ReadFile() to read a plain text file and strings.Split() to break into lines

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// Set up colour formatting for player's guess feedback
const (
	colourReset  = "\033[0m"
	colourGreen  = "\033[32m"
	colourYellow = "\033[33m"
	colourGrey   = "\033[90m"
)

// Helper function for user input validation
func isValidGuess(guess string) bool {
	return len(guess) == 5
}

func loadWords() []string {
	fileContents, err := os.ReadFile("words.txt") // os.ReadFile() returns []byte (not string)
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(fileContents), "\n") // Need to convert []byte to string
	
	// Remove whitespace
	for i, word := range words {
		words[i] = strings.TrimSpace(word)
	}
	return words
}

// main is the entry point - Go automatically calls this function when the program starts
func main() {
	possibleWords := loadWords() // Define and initialise slice with values
	target := possibleWords[rand.Intn(len(possibleWords))] // Select random word (Intn returns 0 to length-1)
	var guess string                                       // Declare variable with explicit type (alternative to short assignment)
	maxAttempts := 3
	attempts := 0

	fmt.Println("Welcome to Wordle! (mode: TRICKY)")

	for attempts < maxAttempts {
		fmt.Println("Guess the 5-letter word:")
		// Read user input from console (&guess passes the memory address so Scan can modify the variable)
		fmt.Scan(&guess)
		if !isValidGuess(guess) {
			fmt.Println("Invalid input. Your guess needs exactly 5 letters.")
			continue // Skip rest of loop if guess is invalid and go to next for loop iteration
		}
		guess = strings.ToUpper(guess) // Convert user input to upper case for comparison
		attempts++                     // Increment counter by 1 (this is a STATEMENT not expression as in JavaScript)

		// Check if user's guess exactly matches target
		if guess == target {
			fmt.Printf("%sCorrect!%s You guessed it in %d attempts\n", colourGreen, colourReset, attempts)
			break // Exit loop on correct guess
		} else {
			fmt.Println("Incorrect")

			// Show letter-by-letter feedback
			// Loop through each position in the 5-letter word
			for i := 0; i < 5; i++ {
				// Compare char at position i in both words
				if guess[i] == target[i] {
					// %s (1st) sets font colour, %s (2nd) resets colour, %d gets replaced by (i+1), %c gets replaced by guess[i]
					fmt.Printf("%sGREEN%s - Position %d: %c is correct!\n", colourGreen, colourReset, i+1, guess[i])
				} else if strings.Contains(target, string(guess[i])) {
					fmt.Printf("%sYELLOW%s - Position %d: %c is in the target word but not in this position.\n", colourYellow, colourReset, i+1, guess[i])
				} else {
					// (Placeholders get filled in order by the arguments provided after the format string)
					fmt.Printf("%sGREY%s - Position %d: %c is not in the target word.\n", colourGrey, colourReset, i+1, guess[i])
				}
			}
			fmt.Println("Try again! Remaining guesses:", maxAttempts-attempts)
		}
	}
	// If all attempts are used without a correct guess
	if guess != target {
		fmt.Println("Game over! The word was", target)
	}
}
