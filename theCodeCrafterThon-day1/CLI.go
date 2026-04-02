package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)
func init() {
	rand.Seed(time.Now().UnixNano())
}
func getValidInput(opt string) (float64, error) {
	var log string
	for {
		fmt.Println(opt)
		fmt.Println(&log)
		value, err := strconv.ParseFloat(log, 64)
		if err != nil {
			fmt.Println("Invalid Input. Please input Numbers.")
			continue
		}
		return value, nil
	}
}

func cliCalculator(){
	var option int

	greetMessages := []string {
		"Welcome on board!",
		"Hey, there!",
		"Welcome back, champ!",
		"Today's looking warm, man!",
		"Look who we have here",
	}
	
	fmt.Println(greetMessages[rand.Intn(len(greetMessages))])

	fmt.Println()
start1:
	fmt.Print("\nProceed: ")
	fmt.Scanln(&option)

	if option == 1 {
		fmt.Println("Access granted!")
	} else if option != 1 {
		
		fmt.Println("Error Found!...Wrong Input Enter 1 to proceed_]")
		goto start1
	}

	var name string
	fmt.Println()
	fmt.Println("[Welcome To Simplified Calculator]")
	fmt.Println()
	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)
	fmt.Printf("\nWelcome %s To CLI-CALCULATOR", name)
	fmt.Println()
	fmt.Println("SELECT OPERATION")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("0. Close - Quit program!")
	fmt.Println("#. HELP")
	fmt.Println("")

	var choice int
	fmt.Scanln(&choice)

	for {
		switch choice {
		case 1:
			Num1 := getValidInput("Enter first number: ")
			Num2 := getValidInput("Enter second number: ")
			fmt.Printf("%g + %g = %g", Num1, Num2, Num1 + Num2)
		}
	}
	
}
func main(){
	cliCalculator()
}