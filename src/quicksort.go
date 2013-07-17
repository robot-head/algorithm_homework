package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var array []int
	// open input file
	fi, err := os.Open("QuickSort.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	// make a read buffer
	r := bufio.NewReader(fi)
	scanner := bufio.NewScanner(r)
	num := 0
	for scanner.Scan() {
		num, err = strconv.Atoi(scanner.Text())
		array = append(array, num)
	}
	
	fmt.Println(array[0:100])
	quicksort(array, 0, len(array) - 1)
	fmt.Println(array[0:100])
}
