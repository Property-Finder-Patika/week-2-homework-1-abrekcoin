package main

import "fmt"

func main() {
	//forClause1(15)
	//forClause2()
	//forClause3()
	//singleCondition()
	//simpleFor5()
	//forRange1()
	forRange2()
}

func forClause1(limit int) {
	for i := 0; i < limit; i++ {
		fmt.Printf("%d\n", i)
	}
}

func forClause2() {
	for i := 0; i < 20; i++ {
		fmt.Printf("%d\n", i)
	}
}

func forClause3() {
	var i int
	for ; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
}

func singleCondition() {
	var i int
	for i < 10 {
		i++
		fmt.Printf("%d\n", i)
	}
}

func forClause4() {
	var i int
	for {
		i++
		if i == 12 {
			break
		}
		fmt.Printf("%d\n", i)
	}
}

func forRange1() {
	slice := []int{1, 3, 5, 7, 9, 11}
	for i, x := range slice {
		fmt.Printf("%dth element is %d\n", i, x)
	}
}

func forRange2() {
	slice := []int{1, 3, 5, 7, 9, 11}
	for x := range slice {
		fmt.Printf("%d\n", x)
	}
	// Sum of what is in slice
	sum := 0
	for x := range slice {
		sum += x
	}
	fmt.Printf("Sum is %d\n", sum)
}
