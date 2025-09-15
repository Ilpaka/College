package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	sorted_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	unsorted_slice := GenerateSlice(10)
	bad_slice := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	fmt.Println("sorted_slice:", sorted_slice)
	fmt.Println("unsorted_slice:", unsorted_slice)
	fmt.Println("bad_slice:", bad_slice)

	sortedBubble_best := BubbleSort(sorted_slice)
	sortedBubble := BubbleSort(unsorted_slice)
	sortedBubble_worst := BubbleSort(bad_slice)

	fmt.Println("Sorted with Bubble Sort (best case):", sortedBubble_best)
	fmt.Println("Sorted with Bubble Sort (worst case):", sortedBubble_worst)
	fmt.Println("Sorted with Bubble Sort:", sortedBubble)

	index := binary_search(sortedBubble, 5)
	fmt.Println("Index of 5 in sorted array:", index)

}
func GenerateSlice(size int) []int {
	if size <= 0 {
		size = 100
	}

	var slice []int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		slice = append(slice, rand.Intn(1000))
	}
	return slice
}

func BubbleSort(slice []int) []int {
	var counter int
	slice = append([]int{}, slice...)
	counter++
	for i := 0; i < len(slice); i++ {
		counter++
		for j := 0; j < len(slice)-i-1; j++ {
			counter++
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				counter += 3
			}
			counter++
		}
	}
	counter++
	fmt.Println(counter)
	return slice
}

//T = N^2 * 5
//T = N
//T = N^2

//В лучшем случае O(n) - массив уже отсортирован операций 102
//В худшем случае O(n^2) - массив отсортирован в обратном порядке операций 200
//В среднем O(n^2) - массив не отсортирован операций 140

func binary_search(slice []int, target int) int {
	var counter int
	low := 0
	counter++
	high := len(slice) - 1
	counter++

	for low <= high {
		counter++
		mid := (low + high) / 2
		counter++
		if slice[mid] == target {
			counter += 2
			return mid
		} else if slice[mid] < target {
			counter += 2
			low = mid + 1
		} else {
			counter += 2
			high = mid - 1
		}
	}
	counter++
	fmt.Println(counter)
	return -1
}

//T = N*10logN

//Во всех случаях (n log n)

func QuickSort(slice []int) []int {
	var counter int
	slice = append([]int{}, slice...)
	counter++
	if len(slice) <= 1 {
		counter += 2
		return slice
	}

	pivot := slice[len(slice)/2]
	counter++
	less := []int{}
	counter++
	more := []int{}
	counter++

	for i := 0; i < len(slice); i++ {
		counter++
		if slice[i] == slice[len(slice)/2] {
			counter += 2
			continue
		}
		if slice[i] < pivot {
			counter += 2
			less = append(less, slice[i])
		} else {
			counter += 2
			more = append(more, slice[i])
		}
	}
	counter++
	fmt.Println(counter)
	return append(append(QuickSort(less), pivot), QuickSort(more)...)
}

//T = 11 * N log N
// В лучшем случае (n log n) - маленькое кол-во n
// В среднем  (n log n) - большое кол-во n
// В худшем случае (n^2)

//В лучшем случае

func InsertionSort(slice []int) []int {
	slice = append([]int{}, slice...)
	for i := 1; i < len(slice); i++ {
		num := slice[i]
		j := i - 1
		for j >= 0 && slice[j] > num {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = num
	}
	return slice
}

func IIsort(slice []int) []int {
	slice = append([]int{}, slice...)

	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1 //предыдущий элемент
		for j >= 0 && slice[j] > key {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = key
	}
	return slice
}

func ShellSort(slice []int) []int {
	slice = append([]int{}, slice...)
	step := len(slice) / 2

	for step > 0 {
		for i := step; i < len(slice); i++ {
			temp := slice[i]
			j := i

			for j >= step && slice[j-step] > temp {
				slice[j] = slice[j-step]
				j -= step
			}
			slice[j] = temp
		}
		step /= 2
	}
	return slice
}
