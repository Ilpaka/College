package main

import (
	"context"
	"math/rand"
	"time"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GenerateSlice(size int) []int {
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

func (a *App) BubbleSort(slice []int) []int {
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

func (a *App) QuickSort(slice []int) []int {
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

	return append(append(a.QuickSort(less), pivot), a.QuickSort(more)...)
}

func (a *App) InsertionSort(slice []int) []int {
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

func (a *App) ShellSort(slice []int) []int {
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
