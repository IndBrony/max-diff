package diff

type compareFunc func(int, int) bool

func MaxDiff(first, second []int) int {
	if len(first) != len(second) {
		return -1
	}
	insertionSort(first, func(a, b int) bool { return a < b })
	insertionSort(second, func(a, b int) bool { return a > b })
	diffTotal := 0
	for index, num := range first {
		diff := second[index] - num
		if diff < 0 {
			diff *= -1
		}
		diffTotal += diff
	}
	return diffTotal
}

//comparer(a, b int) { return a < b } -> ASC
//comparer(a, b int) { return a > b } -> DESC
func insertionSort(arr []int, comparer compareFunc) {
	for i := 1; i < len(arr); i++ {
		for currentValIndex := i; currentValIndex != 0 && comparer(arr[currentValIndex], arr[currentValIndex-1]); currentValIndex-- {
			temp := arr[currentValIndex-1]
			arr[currentValIndex-1] = arr[currentValIndex]
			arr[currentValIndex] = temp
		}
	}
}
