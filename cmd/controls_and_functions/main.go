package main

// Multiple imports
import (
	"errors"
	"fmt"
)

func main() {
	/*
		Functions
	*/
	messageStr := "John Doe"
	printMe(messageStr)

	numerator := 11
	denominator := 2
	var result int = division(numerator, denominator)
	fmt.Println("result:", result)

	var quotient, remainder int = divAndMod(numerator, denominator)
	// Formatted printing using fmt.Printf() function, which allows you to format the output using verbs like %v, %d, %s, etc.
	fmt.Printf("quotient: %v, remainder: %v\n", quotient, remainder)

	var anotherResult, err = divWithError(numerator, denominator)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Result:", anotherResult)
	}

	switch {
	case err != nil:
		fmt.Println("Error:", err.Error())
		// break is implied in Go, so you don't need to write it explicitly.
	case anotherResult > 0:
		fmt.Println("Result is positive:", anotherResult)
	default:
		fmt.Println("Result is zero or negative:", anotherResult)
	}

	switch anotherResult {
	case 0:
		fmt.Println("Result is zero")
	case 1, 2, 3:
		fmt.Println("Result is between 1 and 3")
	default:
		fmt.Println("Result is greater than 3")
	}
}

func printMe(message string) {
	fmt.Println("Hello from printMe() function!", message)
}

func division(numerator int, denominator int) int {
	return numerator / denominator
}

func divAndMod(numerator int, denominator int) (int, int) {
	return numerator / denominator, numerator % denominator
}

func divWithError(numerator int, denominator int) (int, error) {
	var err error // error is a builtin type, with default value of nil, which means no error.
	if denominator == 0 {
		err = errors.New("denominator cannot be zero") // errors.New() function creates a new error with the given message.
		return 0, err
	}
	return numerator / denominator, err
}
