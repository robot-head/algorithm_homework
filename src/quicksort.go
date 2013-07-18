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

func hoarePartition(A []int, pp, l, r int) int {
	fmt.Println("Before partition: ", A[l:r], ". P: ", pp, ", l:", l, ", r:", r)
	A[pp], A[l] = A[l], A[pp]
	p := A[l]
	i := l + 1
	for j := l + 1; j <= r; j++ {
		if A[j] < p {
			A[j], A[i] = A[i], A[j]
			i++
		}
	}
	fmt.Println("Before final swap: ", A[l:r], ". P: ", pp, ", l:", l, ", r:", r)
	A[l], A[i-1] = A[i-1], A[l]
	fmt.Println("After partition: ", A[l:r], ". P: ", i, ", l:", l, ", r:", r)
	return i
}

func quicksort(array []int, left int, right int) int {
	comparisons := 0
	//fmt.Print(left, ",", right, "(", right-left, ")")
	if right-left <= 0 {
		//fmt.Print(" | ")
		return comparisons
	}

	pivotIndex := getPivot(left, right)
	//pivotIndex := getMedianOfThreePivotIndex(array, left, right)
	//pivotNewIndex := partition(array, left, right, pivotIndex)
	//fmt.Println(pivotIndex, ": ", array[pivotIndex], ". ", left, ": ", array[left])
	
	pivotNewIndex := hoarePartition(array, pivotIndex, left, right)
	comparisons += (right - left) - 1
	//fmt.Print(": ", left, ",", pivotNewIndex, ",", right, " | ")
	comparisons += quicksort(array, left, pivotNewIndex-1)
	comparisons += quicksort(array, pivotNewIndex+1, right)
	return comparisons
}

func getPivot(left int, right int) int {
	return left
}

func getMedianOfThreePivotIndex(array []int, left int, right int) int {
	middleIndex := left + int((right-left)/2)
	middleItem := array[middleIndex]
	leftItem := array[left]
	rightItem := array[right]
	if rightItem > leftItem {
		if middleItem > leftItem {
			return middleIndex
		} else {
			return left
		}
	} else {
		if middleItem > rightItem {
			return middleIndex
		} else {
			return right
		}
	}
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
	fmt.Println(array[0:100])
	fmt.Println(array[len(array)-100 : len(array)-1])
}
