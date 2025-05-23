// Wordle game prep (step 1) - very basic program to introduce some basics in Go
// Step 1 asks for user input (to guess the full word!) & checks that against the target word

package main // Declares this file as part of the 'main' package (required for executable programs)

import "fmt" // Imports the fmt package for formatted I/O functions like Println and Scan

// main is the entry point - Go automatically calls this function when the program starts
func main() {
	target := "HELLO" // Define and initialise variable using short assignment operator (:=)
	var guess string  // Declare variable with explicit type (alternative to short assignment)

	// Print welcome messages to console
	fmt.Println("Welcome to Wordle! (mode: HARD)")
	fmt.Println("Guess the 5-letter word:")

	// Read user input from console (&guess passes the memory address so Scan can modify the variable)
	fmt.Scan(&guess)

	// Check if user's guess exactly matches target (case-sensitive)
	if guess == target {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect. The word was:", target)
	}
}
