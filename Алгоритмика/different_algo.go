package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generate_slice(10)
	fmt.Println("Не отсортирвоанный", slice)
	sorted_slice := bubble_sort(slice)
	sorted_insert := insertion_sort(slice)
	quick_sorted := quick_sort(slice)
	shell_sorted := shell_sort(slice)
	fmt.Println("Сортированный", sorted_slice)
	fmt.Println("Сортированный вставкой", sorted_insert)
	fmt.Println("Сортированный быстрой", quick_sorted)
	fmt.Println("Сортированный шелла", shell_sorted)

}

func generate_slice(size int) []int {
	if size <= 0 {
		size = 100
	}

	var slice []int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		slice = append(slice, rand.Intn(100))
	}
	return slice
}

// ? Это рекурсивынй алгоритм, который работает по прринципу разделяй и властвуй.
// ? Суть в том чтобы взять опорный элемент и потом создать 2 массива с меньшим и большим элементами.
// ? проделать это рекурсивно и потом собратьвместе
// ? Один из самых быстрых алгоритмов
func quick_sort(slice []int) []int {
	slice = append([]int{}, slice...)

	if len(slice) <= 1 {
		return slice
	}

	pivot := slice[0]
	less := []int{}
	more := []int{}

	for i := 1; i < len(slice); i++ {
		if slice[i] < pivot {
			less = append(less, slice[i])
		} else {
			more = append(more, slice[i])
		}
	}

	newslice := append(append(quick_sort(less), pivot), quick_sort(more)...)

	return newslice
}

// ? СОртровка шелла рабоытате по принципу пар. Мы делим кол-во элементов попалам
// ?И потом с этим шагом меняем местами элементы. Потом проделываем это пока не станет 1
// ? И потом проходимся сонртирвкой вставками. Работает быстрее так как мы берём уе почти отсортирвоанный массив
func shell_sort(slice []int) []int {
	slice = append([]int{}, slice...)
	step := len(slice) / 2

	for step > 0 {
		for i := step; i < len(slice); i++ {
			temp := slice[i]
			fmt.Println(temp)
			j := i
			fmt.Println("i=", i, "j=", j)
			fmt.Println("slice[j-step]", slice[j-step])
			for j >= step && slice[j-step] > temp {
				slice[j] = slice[j-step]
				j -= step
			}

			slice[j] = temp
			fmt.Println("slice[j]", slice[j])
		}

		step /= 2
	}

	return slice
}

//? Итог:
//? Преимущества Шелла
//? Отсутствие стека: не нужна память под рекурсию
//? Стабильная производительность: нет деградации до O(n²) при плохих данных
//? Простота реализации: без рекурсии
//? Преимущества быстрой сортировки
//? Быстрее в среднем: обычно работает быстрее Шелла
//? Лучшая асимптотика: O(n log n) против O(n log n)²
//? Вывод: Быстрая сортировка обычно быстрее, но Шелла надежнее и проще.

//! ДОП

func bubble_sort(slice []int) []int {
	slice = append([]int{}, slice...)
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}

func insertion_sort(slice []int) []int {
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
