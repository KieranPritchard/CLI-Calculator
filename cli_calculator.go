package main

import (
	"flag"
	"fmt"
)

func main(){
	// Defines command line flags i will be using
	add := flag.String("add","", "The numbers you want to add together")
	sub := flag.String("sub", "", "The numbers you want to subtract ")
	multi := flag.String("multi", "", "The numbers you want to multiply")
	div := flag.String("div", "", "The numbers you want to multiply")
}