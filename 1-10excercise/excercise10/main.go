package main

import "fmt"

func main() {

	//10.1 and 10.2 and 10.3 and 10.4
	const (
		minsperday   = 60 * 24
		WeekDays     = 7
		hoursInDay   = 24
		hoursPerWeek = hoursInDay * WeekDays
		home         = "Milky Way Galaxy"
		lengt        = len(home)
		pi           = 3.14159265358979323846264
		tau          = pi * 2
	)
	fmt.Printf("minutes in 2 weeks is %d \n", minsperday*WeekDays*2) //10.1
	fmt.Printf("There are %d hours in 2 weeks.\n", hoursPerWeek*5)   //10.2
	fmt.Printf("Ther are %d char in %q \n ", lengt, home)            //10.3
	fmt.Printf("tau = %g\n", tau)                                    //10.4

}
