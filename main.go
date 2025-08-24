// Wordle game prep (step 10) - improves guess feedback accuracy to users
// Step 10 Uses a helper function to track each letter count and provide accurate feedback for duplicate letters in guesses

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

// Helper function to track each letter in the target word to give accurate feedback for duplicates
func getWordleFeedback(guess, target string) []string {
	// Create a slice to store feedback for each position (always 5 letters)
	feedback_slice := make([]string, 5)
	// Create a map to count how many of each letter the target has
	targetLetterCounts := make(map[rune]int)

	// Count letters in target word
	for _, char := range target { // _ means ignore the index, just give each character
		targetLetterCounts[char]++ // +1 to the count for this char
	}

	// First pass: mark all exact matches (GREEN)
	for i := 0; i < 5; i++ {
		// Letter matches at this exact position
		if guess[i] == target[i] {
			feedback_slice[i] = "GREEN"
			targetLetterCounts[rune(guess[i])]-- // Subtract 1 from available count
		}
	}

	// Second pass: check remaining letters for wrong-position (YELLOW) or incorrect (GREY) matches
	for i := 0; i < 5; i++ {
		if feedback_slice[i] == "" { // Not already GREEN
			if targetLetterCounts[rune(guess[i])] > 0 {
				// Still available count for this letter in the target
				feedback_slice[i] = "YELLOW"
				targetLetterCounts[rune(guess[i])]-- // Subtract 1 from available count
			} else { // No more of this letter available in target count
				feedback_slice[i] = "GREY"
			}
		}
	}
	// Returns colour match pattern for each letter in target word
	return feedback_slice
}

// main is the entry point - Go automatically calls this function when the program starts
func main() {
	possibleWords := loadWords()                           // Define and initialise slice with values
	target := possibleWords[rand.Intn(len(possibleWords))] // Select random word (Intn returns 0 to length-1)
	var guess string                                       // Declare variable with explicit type (alternative to short assignment)
	maxAttempts := 3
	attempts := 0

	fmt.Println("Welcome to Wordle! (mode: REGULAR)")

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

			// getWordleFeedback handles duplicated letters in the guess properly
			guessFeedback := getWordleFeedback(guess, target)

			// Loop through feedback results from output of letter-tracking function
			for i := 0; i < 5; i++ {
				switch guessFeedback[i] {
				case "GREEN":
					// %s (1st) sets font colour, %s (2nd) resets colour, %d gets replaced by (i+1), %c gets replaced by guess[i]
					fmt.Printf("%sGREEN%s - Position %d: %c is correct!\n", colourGreen, colourReset, i+1, guess[i])
				case "YELLOW":
					fmt.Printf("%sYELLOW%s - Position %d: %c is in the target word but not in this position.\n", colourYellow, colourReset, i+1, guess[i])
				case "GREY":
					// Placeholders get filled in order by the arguments provided after the format string
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
