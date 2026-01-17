package main

import (
	"flag"
	"strings"
	"fmt"
	"strconv"
)

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
	} 
}