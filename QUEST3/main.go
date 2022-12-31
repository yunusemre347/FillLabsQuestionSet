package main

import (
	"fmt"
)

func main() {
	//this var will be our final result
	var selectedItem string
	//create a struct so that later we can assign counters
	type newItems struct {
		fruit        string
		matchCounter int
	}
	var fruits = []string{"apple", "pie", "apple", "red", "red", "red"}
	//var fruits = []string{"apple", "pie", "apple", "blue", "red", "red", "banana", "berry", "banana", "banana"}
	//var fruits = []string{"berry", "berry", "apple", "pie", "apple", "banana", "berry"}

	//new slice for counters
	var newFruits = make([]newItems, 0)
	var counter int = 0
	//loop to check for matches and assign counters
	for _, item := range fruits {
		var matchCounter int = 0
		for _, otherItems := range fruits {
			if item == otherItems {
				matchCounter++
				//fmt.Printf("match between %v -- %v matchCounter:%v \n", item, otherItems, matchCounter)
			}
		}
		var data = newItems{
			fruit:        item,
			matchCounter: matchCounter,
		}
		newFruits = append(newFruits, data)
		counter = counter + 1

		//loop to find our final result:The most repeated item.
		var selectedCounter int = 0
		for _, item := range newFruits {
			if item.matchCounter > selectedCounter {
				selectedCounter = item.matchCounter
				selectedItem = item.fruit
			}
		}
	}
	//fmt.Println(newFruits)
	//fmt.Printf("The most repeated item of the array is:###--%v--### \n", selectedItem)
	fmt.Println(selectedItem)

}
