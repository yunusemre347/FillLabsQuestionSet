package main

import "fmt"

//Given input. Result is going to change proportionally to the input value
var input int = 9

//Number of times the recursion will happen
var runXTimes int = 5

//Counter that i will use inside the funciton.
//Since the recursive is going in reverse my index counter is also subtracts
var counter int = runXTimes + 1

//+1 is because function subtracts in the beginning of the function
//so it wouldnt be equal to -1 when x is 0 and so on
var result int

func recursive(input int, runXTimes int, counter int) int {
	counter--
	//b(0) valuation for recursive base function
	if runXTimes == 0 {
		result = 2
		fmt.Printf("output b(%v):%v \n", runXTimes, result)
		return 2
	} else {
		//b(n)=b(n-1)+(c*input/3)-1 is the recursive formula
		result = recursive(input, runXTimes-1, counter) + (counter * input / 3) - 1
		fmt.Printf("output b(%v) :%v \n", runXTimes, result)
		return result
	}
}
func main() {
	fmt.Printf("input :%v \n", input)
	recursive(input, runXTimes, counter)

}
