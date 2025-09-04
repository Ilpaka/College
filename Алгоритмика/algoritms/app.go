package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// func generate_slice() []int {
// 	var slice []int
// 	rand.Seed(time.Now().UnixNano())

// 	for i := 0; i < 100; i++ {
// 		slice = append(slice, rand.Intn(1000))
// 	}
// 	return slice
// }

// func bubble_Sort(slice []int) []int {
// 	for i := 0; i < len(slice); i++ {
// 		for j := 0; j < len(slice)-i-1; j++ {
// 			if slice[j] > slice[j+1] {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 			}
// 		}
// 	}
// 	return slice
// }

// ===== ПАРАЛЛЕЛЬНАЯ СОРТИРОВКА БЭТЧЕРА =====

const (
	SORT_ASC  = true
	SORT_DESC = false
)

// Сравнение и обмен элементов
func compareAndSwap(arr []int, i, j int, dir bool) {
	if (arr[i] > arr[j]) == dir {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Битоническое слияние
func bitonicMerge(arr []int, lo, n int, dir bool, wg *sync.WaitGroup) {
	defer wg.Done()

	if n > 1 {
		m := n / 2

		// Сравниваем и обмениваем элементы
		for i := lo; i < lo+m; i++ {
			compareAndSwap(arr, i, i+m, dir)
		}

		// Параллельно обрабатываем половины
		var wg1, wg2 sync.WaitGroup
		wg1.Add(1)
		wg2.Add(1)

		go bitonicMerge(arr, lo, m, dir, &wg1)
		go bitonicMerge(arr, lo+m, m, dir, &wg2)

		wg1.Wait()
		wg2.Wait()
	}
}

// Основная функция битонической сортировки
func bitonicSort(arr []int, lo, n int, dir bool, wg *sync.WaitGroup) {
	defer wg.Done()

	if n > 1 {
		m := n / 2
		var wg1, wg2 sync.WaitGroup

		wg1.Add(1)
		wg2.Add(1)

		// Сортируем половины в разных направлениях
		go bitonicSort(arr, lo, m, SORT_ASC, &wg1)
		go bitonicSort(arr, lo+m, m, SORT_DESC, &wg2)

		wg1.Wait()
		wg2.Wait()

		// Сливаем битоническую последовательность
		var wgMerge sync.WaitGroup
		wgMerge.Add(1)
		go bitonicMerge(arr, lo, n, dir, &wgMerge)
		wgMerge.Wait()
	}
}

// Подготовка массива к сортировке (размер должен быть 2^n)
func prepareArrayForBatcher(arr []int) ([]int, int) {
	originalLen := len(arr)

	// Находим ближайшую степень двойки
	size := 1
	for size < len(arr) {
		size <<= 1
	}

	// Если размер не равен степени двойки, добавляем максимальные значения
	if size > len(arr) {
		maxVal := 0
		for _, val := range arr {
			if val > maxVal {
				maxVal = val
			}
		}

		// Добавляем максимальные значения, чтобы они оказались в конце
		for len(arr) < size {
			arr = append(arr, maxVal+1)
		}
	}

	return arr, originalLen
}

// Публичная функция для вызова из фронтенда
func (a *App) ParallelBatcherSort(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}

	// Подготавливаем массив
	prepared, originalLen := prepareArrayForBatcher(append([]int{}, slice...))

	// Запускаем сортировку
	var wg sync.WaitGroup
	wg.Add(1)
	go bitonicSort(prepared, 0, len(prepared), SORT_ASC, &wg)
	wg.Wait()

	// Возвращаем только оригинальную часть
	return prepared[:originalLen]
}

// Генерация среза (публичная функция для фронтенда)
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

// Пузырьковая сортировка (публичная функция для фронтенда)
func (a *App) BubbleSort(slice []int) []int {
	result := append([]int{}, slice...)

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

func quickSort(arr []int, low, high int) {
	if low < high {
		// Разделяем массив и получаем позицию опорного элемента
		pivotIndex := partition(arr, low, high)

		// Рекурсивно сортируем левую и правую части
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// Выбираем последний элемент как опорный
	pivot := arr[high]
	i := low - 1 // Индекс меньшего элемента

	for j := low; j < high; j++ {
		// Если текущий элемент меньше или равен опорному
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Ставим опорный элемент на правильное место
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

type TournamentNode struct {
	value     int
	index     int
	processed bool
}

type TournamentTree struct {
	nodes []TournamentNode
	size  int
}

func NewTournamentTree(arr []int) *TournamentTree {
	n := len(arr)
	size := 1
	for size < n {
		size <<= 1
	}

	nodes := make([]TournamentNode, 2*size)

	// Инициализируем листья
	for i := 0; i < size; i++ {
		if i < n {
			nodes[size+i] = TournamentNode{
				value:     arr[i],
				index:     i,
				processed: false,
			}
		} else {
			nodes[size+i] = TournamentNode{
				value:     math.MaxInt32,
				index:     -1,
				processed: true,
			}
		}
	}

	// Строим внутренние узлы
	for i := size - 1; i > 0; i-- {
		left := &nodes[2*i]
		right := &nodes[2*i+1]

		if !left.processed && (right.processed || left.value <= right.value) {
			nodes[i] = *left
		} else if !right.processed {
			nodes[i] = *right
		} else {
			nodes[i] = TournamentNode{processed: true}
		}
	}

	return &TournamentTree{nodes: nodes, size: size}
}

func (t *TournamentTree) ExtractMin() (int, bool) {
	if t.nodes[1].processed {
		return 0, false
	}

	minVal := t.nodes[1].value

	// Находим и исключаем минимальный элемент
	pos := 1
	for pos < t.size {
		if !t.nodes[2*pos].processed && t.nodes[2*pos].value == minVal {
			pos = 2 * pos
		} else {
			pos = 2*pos + 1
		}
	}

	// Исключаем элемент
	t.nodes[pos].processed = true

	// Перестраиваем путь к корню
	pos /= 2
	for pos > 0 {
		left := &t.nodes[2*pos]
		right := &t.nodes[2*pos+1]

		if !left.processed && (right.processed || left.value <= right.value) {
			t.nodes[pos] = *left
		} else if !right.processed {
			t.nodes[pos] = *right
		} else {
			t.nodes[pos] = TournamentNode{processed: true}
		}
		pos /= 2
	}

	return minVal, true
}

func TournamentSortAdvanced(arr []int) []int {
	if len(arr) <= 1 {
		return append([]int{}, arr...)
	}

	tree := NewTournamentTree(arr)
	result := make([]int, 0, len(arr))

	for {
		if val, ok := tree.ExtractMin(); ok {
			result = append(result, val)
		} else {
			break
		}
	}

	return result
}

// binaryInsertionSortOptimized с ранним выходом для уже отсортированных элементов
func binaryInsertionSortOptimized(arr []int) {
	for i := 1; i < len(arr); i++ {
		val := arr[i]

		// Если элемент уже на правильном месте, пропускаем
		if val >= arr[i-1] {
			continue
		}

		pos := binarySearch(arr, val, 0, i)

		// Используем copy для более эффективного сдвига
		copy(arr[pos+1:i+1], arr[pos:i])
		arr[pos] = val
	}
}

func binarySearch(arr []int, val int, start int, end int) int {
	for start < end {
		mid := start + (end-start)/2
		if arr[mid] < val {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

// QuickSort - публичная функция для фронтенда
func (a *App) QuickSort(slice []int) []int {
	if len(slice) <= 1 {
		return append([]int{}, slice...)
	}

	result := make([]int, len(slice))
	copy(result, slice)

	quickSort(result, 0, len(result)-1)
	return result
}

// TournamentSort - публичная функция для фронтенда
func (a *App) TournamentSort(slice []int) []int {
	return TournamentSortAdvanced(slice)
}

// BinaryInsertionSort - публичная функция для фронтенда
func (a *App) BinaryInsertionSort(slice []int) []int {
	if len(slice) <= 1 {
		return append([]int{}, slice...)
	}

	result := make([]int, len(slice))
	copy(result, slice)

	binaryInsertionSortOptimized(result)
	return result
}
