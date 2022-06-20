package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Please enter an integer: ")
	fmt.Scanln(&input)
	anInt, err := strconv.Atoi(input)

	if err != nil {
		fmt.Errorf("you did not enter an integer %v", err)
	}

	switch {
	case anInt < 0:
		fmt.Printf("%d is negative!\n", anInt)
	case anInt == 0:
		fmt.Printf("%d is zero!\n", anInt)
	case anInt > 0:
		fmt.Printf("%d is positive!\n", anInt)
	}

	if anInt < 0 {
		fmt.Printf("%d is negative!\n", anInt)
	} else if anInt == 0 {
		fmt.Printf("%d is zero!\n", anInt)
	} else {
		fmt.Printf("%d is positive!\n", anInt)
	}
}
