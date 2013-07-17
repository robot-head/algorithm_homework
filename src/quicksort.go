package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func quicksort(array []int, left int, right int) int {
	comparisons := 0
	if left < right {
		//pivotIndex := getPivot(left, right)
		pivotIndex := getMedianOfThreePivotIndex(array, left, right)
		pivotNewIndex := partition(array, left, right, pivotIndex)
		comparisons += ((pivotNewIndex - 1) - left) - 1
		comparisons += (right - (pivotNewIndex + 1)) - 1
		comparisons += quicksort(array, left, pivotNewIndex-1)
		comparisons += quicksort(array, pivotNewIndex+1, right)
	}
	return comparisons
}

func getPivot(left int, right int) int {
	return right
}

func getMedianOfThreePivotIndex(array []int, left int, right int) int {
	medianItem := array[(right-left)/2]
	leftItem := array[left]
	rightItem := array[right]
	items := []int{leftItem, medianItem, rightItem}
	is := &intSorter{array: array}
	sort.Sort(is)
	return items[2]
}

type intSorter struct {
	array []int
}

func (i *intSorter) Len() int {
	return len(i.array)
}

func (i *intSorter) Swap(x, y int) {
	i.array[x], i.array[y] = i.array[y], i.array[x]
}

func (i *intSorter) Less(x, y int) bool {
	return x < y
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

	fmt.Println(quicksort(array, 0, len(array)-1))
}
