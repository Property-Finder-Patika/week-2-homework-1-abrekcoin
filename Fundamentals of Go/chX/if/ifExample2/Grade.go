package main

import (
	"fmt"
	"strconv"
)

func main() {

	var testScore string
	fmt.Println("Please enter test score as an integer:")
	fmt.Scanln(&testScore)
	score, err := strconv.Atoi(testScore)

	if err != nil {
		fmt.Errorf("you did not enter an integer %v", err)
	}

	var grade rune

	if score > 100 {
		fmt.Println("Test score can't be more than 100!")
		return
	} else if score >= 90 {
		grade = 'A'
	} else if score >= 80 {
		grade = 'B'
	} else if score >= 70 {
		grade = 'C'
	} else if score >= 60 {
		grade = 'D'
	} else { // Comment out this else see what happens to grade variable
		grade = 'F'
	}
	fmt.Printf("Your grade is %c ", grade)

	if grade == 'A' {
		fmt.Println("👏👏👏")
	} else if grade == 'F' {
		fmt.Println("😩")
	}
}
