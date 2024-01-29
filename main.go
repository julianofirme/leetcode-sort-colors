package main

import "fmt"

func quicksort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)

		quicksort(arr, low, pivotIndex-1)
		quicksort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func sortColors(nums []int) {
	quicksort(nums, 0, len(nums)-1)

	fmt.Println(nums)
	return
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}

	sortColors(nums)
}
