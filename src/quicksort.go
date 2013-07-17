package main

import (
	"fmt"
)

func partition(array []int, left int, right int, pivotIndex int) int {
	pivotValue := array[pivotIndex]
	array[right], array[pivotIndex] = array[pivotIndex], array[right]
	storeIndex := left
	for i := left; i < right; i++ {
		if array[i] <= pivotValue {
			array[i], array[storeIndex] = array[storeIndex], array[i]
			storeIndex++
		}
	}
	array[storeIndex], array[right] = array[right], array[storeIndex]
	return storeIndex
}

func quicksort(array []int, left int, right int) {
	if left < right {
		pivotIndex := getPivot(left, right)
		pivotNewIndex := partition(array, left, right, pivotIndex)
		quicksort(array, left, pivotNewIndex-1)
		quicksort(array, pivotNewIndex+1, right)
	}
}

func getPivot(left int, right int) int {
	return right
}

func main() {
	array := []int{5, 2, 3, 1, 4}
	fmt.Println(array)
	quicksort(array, 0, 4)
	fmt.Println(array)
}
