// Wordle game prep (step 4) - introduces a Go string function
// Step 4 adds case-handling for user guesses 

package main

import (
	"fmt"
	"strings"
)

// main is the entry point - Go automatically calls this function when the program starts
func main() {
	target := "HELLO" // Define and initialise variable using short assignment operator (:=)
	var guess string  // Declare variable with explicit type (alternative to short assignment)
	maxAttempts := 3
	attempts := 0

	fmt.Println("Welcome to Wordle! (mode: TRICKY)")

	for attempts < maxAttempts {
		fmt.Println("Guess the 5-letter word:")
		// Read user input from console (&guess passes the memory address so Scan can modify the variable)
		fmt.Scan(&guess)
		guess = strings.ToUpper(guess) // Convert user input to upper case for comparison
		attempts++ // Increment counter by 1 (this is a STATEMENT not expression as in JavaScript)

		// Check if user's guess exactly matches target (case-sensitive)
		if guess == target {
			fmt.Println("Correct! You guessed it in", attempts, "attempts")
			break // Exit loop on correct guess
		} else {
			fmt.Println("Incorrect")
		
			// Show letter-by-letter feedback
			// Loop through each position in the 5-letter words
			for i := 0; i < 5; i++ {
				// Compare char at position i in both words
				if guess[i] == target[i] {
					// %d gets replaced by (i+1), %c gets replaced by guess[i]
					fmt.Printf("GREEN - Position %d: %c is correct!\n", i+1, guess[i])
				} else if strings.Contains(target, string(guess[i])) {
					fmt.Printf("YELLOW - Position %d: %c is in the target word but not in this position.\n", i+1, guess[i])
				} else {
					// (Placeholders get filled in order by the arguments provided after the format string)
					fmt.Printf("GREY - Position %d: %c is not in the target word.\n", i+1, guess[i])
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
