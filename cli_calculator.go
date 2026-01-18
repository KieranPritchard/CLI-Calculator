package main

import (
	"flag"
	"strings"
	"fmt"
	"strconv"
)

// Function to parse two ints
func parseTwoInts(input string) (int, int, error) {
	// Splits the strings at a space
	parts := strings.Split(input, " ")
	// Checks if the parts are greater than two
	if len(parts) != 2 {
		// Returns an error message
		return 0, 0, fmt.Errorf("please provide exactly two numbers")
	}

	// converts the string to integer
	num_1, err := strconv.Atoi(parts[0])
	// Checks for error
	if err != nil {
		// Returns failed conversion
		return 0, 0, err
	}

	// converts the string to integer
	num_2, err := strconv.Atoi(parts[1])
	// Checks for error
	if err != nil {
		// Returns failed conversion
		return 0, 0, err
	}

	// Returns numbers and errors
	return num_1, num_2, nil
}

func main(){
	// Defines command line flags i will be using
	add := flag.String("add","", "The numbers you want to add together")
	sub := flag.String("sub", "", "The numbers you want to subtract ")
	multi := flag.String("multi", "", "The numbers you want to multiply")
	div := flag.String("div", "", "The numbers you want to multiply")

	// Parses the flags
	flag.Parse()

	// Checks for the add flag
	if *add != "" {
		// Gets the two numbers and the error variaable from the function
		num_1, num_2, err := parseTwoInts(*add)
		// Checks if error isnt nil
		if err != nil {
			// Outputs the error
			fmt.Println("Error:", err)
			// Returns nothing
			return
		}
		// Outputs the sum
		fmt.Println(num_1 + num_2)
	// Checks for subtraction
	} else if *sub != "" {
		// Gets the two numbers and the error variaable from the function
		num_1, num_2, err := parseTwoInts(*sub)
		// Checks if error isnt nil
		if err != nil {
			// Outputs the error
			fmt.Println("Error:", err)
			// Returns nothing
			return
		}
		// Outputs the sum
		fmt.Println(num_1 - num_2)
	// Checks for multiplication
	} else if *multi != "" {
		// Gets the two numbers and the error variaable from the function
		num_1, num_2, err := parseTwoInts(*multi)
		// Checks if error isnt nil
		if err != nil {
			// Outputs the error
			fmt.Println("Error:", err)
			// Returns nothing
			return
		}
		// Outputs the sum
		fmt.Println(num_1 * num_2)
	// Checks for division
	} else if *div != "" {
		// Gets the two numbers and the error variaable from the function
		num_1, num_2, err := parseTwoInts(*sub)
		// Checks if error isnt nil
		if err != nil {
			// Outputs the error
			fmt.Println("Error:", err)
			// Returns nothing
			return
		}
		// Outputs the sum
		fmt.Println(num_1 / num_2)
	} else {
		fmt.Println("Invalid Input")
	}
}