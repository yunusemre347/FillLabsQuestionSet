package main

import "fmt"

//given input
var input int = 9

//number of times the recursion will happen
var x int = 5

//counter that i will use inside the funciton.
//since the recursive is going in reverse my index counter is also subtracts
var c int = x + 1

//+1 is because function subtracts in the beginning of the function
//so it wouldnt be equal to -1 when x is 0 and so on
var result int

func recursive(input int, x int, c int) int {
	c--
	//b(0) valuation for recursive base function
	if x == 0 {
		result = 2
		fmt.Printf("output b(%v):%v \n", x, result)
		return 2
	} else {
		//b(n)=b(n-1)+(c*input/3)-1 is the recursive formula
		result = recursive(input, x-1, c) + (c * input / 3) - 1
		fmt.Printf("output b(%v) :%v \n", x, result)
		return result
	}
}
func main() {
	fmt.Printf("input :%v \n", input)
	recursive(input, x, c)

}
