package main

import "fmt"

func main() {
	//8.1
	fmt.Println(50 + 25)
	fmt.Println(50 - 15.5)
	fmt.Println(50 * 0.5)
	fmt.Println(50 / 0.5)
	fmt.Println(25 % 3)
	fmt.Println(-(5 + 2))
	//8.2
	x := float64(5) / 2
	fmt.Println(x)
	fmt.Println(10 + 5 - (5 - 10))
	fmt.Println(-10 + 0.5 - (1 + 5.5))

	// 8.3
	fmt.Println(5 + 10*(2-5))
	fmt.Println(0.5 * (2 - 1))
	fmt.Println((3+1)/2*10 + 4)
	fmt.Println(10 / 2 * (10 % 7))
	fmt.Println(100 / (5. / 2))

	//8.4

	counter, factor := 45, 0.5

	counter++
	counter++
	counter++
	counter++
	counter++
	factor--
	factor--

	fmt.Println(float64(counter) * factor)
	//8.5
	var counter2 int

	counter2++
	counter2--
	counter2 += 5
	counter2 *= 10
	counter2 /= 5
	fmt.Println(counter2)

	//8.6
	width, height := 10, 2

	/* width = width + 1
	width = width + height
	width = width - 1
	width = width - height
	width = width * 20
	width = width / 25
	width = width % 5 */

	width++
	width += height
	width--
	width -= height
	width *= 20
	width /= 25
	width %= 5

	fmt.Println(width)

	//8.7

}
