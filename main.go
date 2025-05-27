// Wordle game prep (step 2) - very basic program to introduce some basics in Go
// Step 2 allows the user more that 1 guess (to guess the full word!) with a for loop and counter

package main

import "fmt"

// main is the entry point - Go automatically calls this function when the program starts
func main() {
	target := "HELLO" // Define and initialise variable using short assignment operator (:=)
	var guess string  // Declare variable with explicit type (alternative to short assignment)
	maxAttempts := 3
	attempts := 0

	// Print welcome message to console
	fmt.Println("Welcome to Wordle! (mode: TRICKY)")

	for attempts < maxAttempts {
		fmt.Println("Guess the 5-letter word:")
		// Read user input from console (&guess passes the memory address so Scan can modify the variable)
		fmt.Scan(&guess)
		attempts++ // Increment counter by 1 (this is a STATEMENT not expression as in JavaScript)

		// Check if user's guess exactly matches target (case-sensitive)
		if guess == target {
			fmt.Println("Correct! You guessed it in", attempts, "attempts")
			break // Exit loop on correct guess
		} else {
			fmt.Println("Incorrect. Try again! Remaining guesses:", maxAttempts-attempts)
		}
	}
	// If all attempts are used without a correct guess
	if guess != target {
		fmt.Println("Game over! The word was", target)
	}
}
