package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func getValidInput(opt string) float64 {
	var log string
	for {
		fmt.Print(opt)
		fmt.Scan(&log)

		if value, err := strconv.ParseInt(log, 0, 64); err == nil {
			return float64(value)
		}
		if value, err := strconv.ParseFloat(log, 64); err == nil {
			return value
		}
		fmt.Println("Invalid Input. PLease input Numbers.")
	}
}

func cliCalculator() {
	var option int
	var choice int

	greetMessages := []string{
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
	fmt.Printf("\nHey %s Welcome To CLI-CALCULATOR", name)
	fmt.Println()
start2:

    fmt.Println()
	fmt.Println("SELECT OPERATION")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("0. Close - Quit program!")
	fmt.Println("10. HELP")
	fmt.Println("")
	fmt.Scanln(&choice)

	// if choice == 1 || choice == 2 || choice == 3 || choice == 4 || choice == 0 || choice == 00 {
	// 	return
	// }else if choice != 1 || choice != 2 || choice != 3 || choice != 4 || choice != 0 || choice != 00 {
	// 	fmt.Println("Invalid command: Press # for help ")
	// }else {
	// 	// if err != nil {
	// 	// 	return 
	// 	// }
	// 	fmt.Println("Error: Input Number")
	// 	goto start2
	// }

	switch choice {
	case 1:
		num1 := getValidInput("Enter first number: ")
		num2 := getValidInput("Enter second number: ")
		fmt.Printf("%g + %g = %g\n", num1, num2, num1+num2)
		goto start2

	case 2:
		num1 := getValidInput("Enter first number: ")
		num2 := getValidInput("Enter second number: ")
		fmt.Printf("%g - %g = %g\n", num1, num2, num1-num2)
		goto start2

	case 3:
		num1 := getValidInput("Enter first number: ")
		num2 := getValidInput("Enter second number: ")
		fmt.Printf("%g * %g = %g\n", num1, num2, num1*num2)
		goto start2
		
	case 4:
		num1 := getValidInput("Enter first number: ")
		num2 := getValidInput("Enter second number: ")
		fmt.Printf("%g / %g = %g\n", num1, num2, num1/num2)
		goto start2

	case 0:
		fmt.Println()
		fmt.Println("Good bye, champ!")

	default:
		fmt.Println()
		fmt.Println("Unknown command!. Press 10 for HELP")
		fmt.Println()
		goto start2

	}

}
func main() {

	cliCalculator()
}
