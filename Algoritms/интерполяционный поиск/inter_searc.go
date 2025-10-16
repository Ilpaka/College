package main

import (
	"fmt"
	"time"
)

func main() {
	inter()
	binary()
}

func inter() {
	fmt.Println("-----------------Интерполяционный поиск------------------")
	slice100 := GenerateSlice(100)
	slice500 := GenerateSlice(500)
	slice1000 := GenerateSlice(1000000)

	// Для 100 элементов
	timer_100 := time.Now()
	pos100 := InterpolationSearch(slice100, 56)
	timer_100_end := time.Now()
	duration_100 := timer_100_end.Sub(timer_100).Nanoseconds()
	duration_100_ms := float64(duration_100) / 1e6
	fmt.Printf("Время выполнения 100: %.3f мс (%d нс)\n", duration_100_ms, duration_100)
	fmt.Println("Позиция:", pos100)

	// Для 500 элементов
	timer_500 := time.Now()
	pos500 := InterpolationSearch(slice500, 10)
	timer_500_end := time.Now()
	duration_500 := timer_500_end.Sub(timer_500)
	duration_500_ms := float64(duration_500.Nanoseconds()) / 1e6
	fmt.Printf("Время выполнения 500: %.3f мс (%v)\n", duration_500_ms, duration_500)
	fmt.Println("Позиция:", pos500)

	// Для 10000 элементов
	timer_1000 := time.Now()
	pos1000 := InterpolationSearch(slice1000, 945)
	timer_1000_end := time.Now()
	duration_10000 := timer_1000_end.Sub(timer_1000)
	duration_10000_ms := float64(duration_10000.Nanoseconds()) / 1e6
	fmt.Printf("Время выполнения 10000: %.3f мс (%v)\n", duration_10000_ms, duration_10000)
	fmt.Println("Позиция:", pos1000)
}

func binary() {
	fmt.Println("-----------------Бинарный поиск------------------")
	slice100 := GenerateSlice(100)
	slice500 := GenerateSlice(500)
	slice1000 := GenerateSlice(2000000)

	timer_100 := time.Now()
	pos100 := BinarySearch(slice100, 56)
	timer_100_end := time.Now()
	duration_100 := timer_100_end.Sub(timer_100).Nanoseconds()
	duration_100_ms := float64(duration_100) / 1e6
	fmt.Printf("Время выполнения 100: %.3f мс (%d нс)\n", duration_100_ms, duration_100)
	fmt.Println("Позиция:", pos100)

	// Для 500 элементов
	timer_500 := time.Now()
	pos500 := BinarySearch(slice500, 10)
	timer_500_end := time.Now()
	duration_500 := timer_500_end.Sub(timer_500)
	duration_500_ms := float64(duration_500.Nanoseconds()) / 1e6
	fmt.Printf("Время выполнения 500: %.3f мс (%v)\n", duration_500_ms, duration_500)
	fmt.Println("Позиция:", pos500)

	// Для 10000 элементов
	timer_1000 := time.Now()
	pos1000 := BinarySearch(slice1000, 945)
	timer_1000_end := time.Now()
	duration_10000 := timer_1000_end.Sub(timer_1000)
	duration_10000_ms := float64(duration_10000.Nanoseconds()) / 1e6
	fmt.Printf("Время выполнения 10000: %.3f мс (%v)\n", duration_10000_ms, duration_10000)
	fmt.Println("Позиция:", pos1000)
}

func GenerateSlice(size int) []int {
	if size <= 0 {
		size = 100
	}
	slice := make([]int, 0, size)
	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}
	return slice
}

func InterpolationSearch(slice []int, target int) int {
	slice = append([]int{}, slice...)

	left := 0
	right := len(slice) - 1

	// Основной цикл поиска
	for left <= right && target >= slice[left] && target <= slice[right] {

		if left == right {
			if slice[left] == target {
				return left
			}
			return -1
		}

		if slice[right] == slice[left] {

			if slice[left] == target {
				return left
			}
			return -1
		}

		pos := left + ((target-slice[left])*(right-left))/(slice[right]-slice[left])

		if pos < left || pos > right {
			return -1
		}

		if slice[pos] == target {
			return pos
		}

		if slice[pos] < target {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}

	return -1
}

func BinarySearch(slice []int, target int) int {
	slice = append([]int{}, slice...)
	left, right := 0, len(slice)-1

	for left <= right {
		mid := left + (right-left)/2

		if slice[mid] == target {
			return mid
		}

		if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
