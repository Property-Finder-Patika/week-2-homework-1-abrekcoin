package main

import (
	"fmt"
	"strconv"
)

func main() {
	var aString string
	fmt.Println("Please enter an integer for a month:")
	fmt.Scanln(&aString)
	monthInt, err := strconv.Atoi(aString)

	if err != nil {
		fmt.Errorf("you did not enter an integer %v", err)
	}

	var month string
	switch monthInt {
	case 1:
		month = "January"
	case 2:
		month = "February"
	case 3:
		month = "March"
	case 4:
		month = "April"
	case 5:
		month = "May"
	case 6:
		month = "June"
	case 7:
		month = "July"
	case 8:
		month = "August"
	case 9:
		month = "September"
	case 10:
		month = "October"
	case 11:
		month = "November"
	case 12:
		month = "December"
	default:
		fmt.Printf("No mathching month for %s is found!", aString)
		return
	}
	fmt.Printf("Month for %s is %s", aString, month)

	var season string

	switch monthInt {
	case 12, 1, 2:
		season = "Winter"
	case 3, 4, 5:
		season = "Spring"
	case 6, 7, 8:
		season = "Summer"
	case 9, 10, 11:
		season = "Fall"
	}
	fmt.Printf(" and season for %s is %s", month, season)
}
