package diff

import (
	"strconv"
	"strings"
)

//CompareFunc is Type declaration for comparator func to be used in InsertionSort
type CompareFunc func(int, int) bool

//MaxDiff counts the maximum difference between 2 distinct int slices
func MaxDiff(first, second []int) int {
	//return -1 if the length id different
	if len(first) != len(second) {
		return -1
	}
	//sort the first slice ascending and the second descending
	insertionSort(first, func(a, b int) bool { return a < b })
	insertionSort(second, func(a, b int) bool { return a > b })

	diffTotal := 0
	//loop through the first slice and diff-ing the element to the corresponding value on second slice
	//and add it to the total diff
	for index, num := range first {
		diff := second[index] - num
		//if the diff is negative then flip it
		//Note: using math.Abs required to convert each number into float64
		//then converting the result to int
		//for this usecase, it's better to use this method
		if diff < 0 {
			diff *= -1
		}
		diffTotal += diff
	}
	//return the total diff
	return diffTotal
}

//MaxDiffString is a string wrapper for MaxDiff func to be compliant with JumatHek's problem
func MaxDiffString(input string) int {
	//split by "-" to get the first and second slice
	splitted := strings.Split(input, "-")
	//if the result is not 2 string then return -1
	if len(splitted) != 2 {
		return -1
	}
	//split each string by space to get the individual num
	//if the length of the result is not same then return -1
	firstString := strings.Split(splitted[0], " ")
	secondString := strings.Split(splitted[1], " ")
	if len(firstString) != len(secondString) {
		return -1
	}
	//convert the number into int type
	first, second := make([]int, len(firstString)), make([]int, len(firstString))
	for index, num := range firstString {
		firstConverted, err := strconv.Atoi(num)
		if err != nil {
			return -1
		}
		secondConverted, err := strconv.Atoi(secondString[index])
		if err != nil {
			return -1
		}
		first[index] = firstConverted
		second[index] = secondConverted
	}
	return MaxDiff(first, second)
}

//comparer(a, b int) { return a < b } -> ASC
//comparer(a, b int) { return a > b } -> DESC
func insertionSort(arr []int, comparer CompareFunc) {
	for i := 1; i < len(arr); i++ {
		for currentValIndex := i; currentValIndex != 0 && comparer(arr[currentValIndex], arr[currentValIndex-1]); currentValIndex-- {
			temp := arr[currentValIndex-1]
			arr[currentValIndex-1] = arr[currentValIndex]
			arr[currentValIndex] = temp
		}
	}
}
