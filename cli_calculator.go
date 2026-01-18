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
		// Splits the numbers
		numbers := strings.SplitAfter(*add, " ")

		// Converts the numbers in the arguement to integers
		num_1, err := strconv.Atoi(numbers[0])
		num_2, err := strconv.Atoi(numbers[1])

		// Checks if the error is not nil
		if err != nil{
			// Outputs the error message
			fmt.Println("Error", err)
		} else {
			// Calculates the sum
			sum := num_1 + num_2
			
			// Outputs the sum
			fmt.Println(sum)
		}
	// Checks for subtraction
	} else if *sub != "" {
		// Splits the numbers
		numbers := strings.SplitAfter(*add, " ")

		// Converts the numbers in the arguement to integers
		num_1, err := strconv.Atoi(numbers[0])
		num_2, err := strconv.Atoi(numbers[1])

		// Checks if the error is not nil
		if err != nil{
			// Outputs the error message
			fmt.Println("Error", err)
		} else {
			// Calculates the sum
			sum := num_1 - num_2
			
			// Outputs the sum
			fmt.Println(sum)
		}
	// Checks for multiplication
	} else if *multi != "" {
		// Splits the numbers
		numbers := strings.SplitAfter(*add, " ")

		// Converts the numbers in the arguement to integers
		num_1, err := strconv.Atoi(numbers[0])
		num_2, err := strconv.Atoi(numbers[1])

		// Checks if the error is not nil
		if err != nil{
			// Outputs the error message
			fmt.Println("Error", err)
		} else {
			// Calculates the sum
			sum := num_1 * num_2
			
			// Outputs the sum
			fmt.Println(sum)
		}
	// Checks for division
	} else if *div != "" {
		// Splits the numbers
		numbers := strings.SplitAfter(*add, " ")

		// Converts the numbers in the arguement to integers
		num_1, err := strconv.Atoi(numbers[0])
		num_2, err := strconv.Atoi(numbers[1])

		// Checks if the error is not nil
		if err != nil{
			// Outputs the error message
			fmt.Println("Error", err)
		} else {
			// Calculates the sum
			sum := num_1 / num_2
			
			// Outputs the sum
			fmt.Println(sum)
		}
	} else{
		fmt.Println("Invaild input")
	}
}