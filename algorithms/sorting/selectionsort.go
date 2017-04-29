package main

import (
	"fmt"
)

func selectionSort(unsorted []int) []int {
	var min int = 0
	var temp int = 0
	
	for i := 0; i < len(unsorted); i++ {
	min = i
	for x := i + 1; x < len(unsorted); x++{
	
	if(unsorted[x] < unsorted[min]){
		min = x
	}
	temp = unsorted[i]
	unsorted[i] = unsorted[min]
	unsorted[min] = temp
	
	}
	
	
	}
	
	return unsorted
}


func main() {
	unsorted := []int{2,7,4,1,5,3}
	fmt.Println(selectionSort(unsorted))
}
