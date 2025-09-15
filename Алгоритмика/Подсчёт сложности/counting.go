package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Три набора данных: лучший, средний, худший
	sorted := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	random := GenerateSlice(10)
	reversed := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	fmt.Println("sorted_slice:", sorted)
	fmt.Println("unsorted_slice:", random)
	fmt.Println("bad_slice:", reversed)
	fmt.Println("==========================================")

	// Bubble
	fmt.Println("ПУЗЫРЬКОВАЯ СОРТИРОВКА (Bubble):")
	fmt.Print("Best: ")
	BubbleSort(sorted)
	fmt.Print("Avg:  ")
	BubbleSort(random)
	fmt.Print("Worst:")
	BubbleSort(reversed)
	fmt.Println()

	// Insertion
	fmt.Println("ВСТАВКАМИ (Insertion):")
	fmt.Print("Best: ")
	InsertionSort(sorted)
	fmt.Print("Avg:  ")
	InsertionSort(random)
	fmt.Print("Worst:")
	InsertionSort(reversed)
	fmt.Println()

	// Selection
	fmt.Println("ВЫБОРОМ (Selection):")
	fmt.Print("Best: ")
	SelectionSort(sorted)
	fmt.Print("Avg:  ")
	SelectionSort(random)
	fmt.Print("Worst:")
	SelectionSort(reversed)
	fmt.Println()

	// Quick
	fmt.Println("БЫСТРАЯ (Quick):")
	fmt.Print("Best: ")
	QuickSortWithCounter(sorted)
	fmt.Print("Avg:  ")
	QuickSortWithCounter(random)
	fmt.Print("Worst:")
	QuickSortWithCounter(reversed)
	fmt.Println()

	// Merge
	fmt.Println("СЛИЯНИЕМ (Merge):")
	fmt.Print("Best: ")
	MergeSortWithCounter(sorted)
	fmt.Print("Avg:  ")
	MergeSortWithCounter(random)
	fmt.Print("Worst:")
	MergeSortWithCounter(reversed)
	fmt.Println()

	// Heap
	fmt.Println("ПИРАМИДАЛЬНАЯ (Heap):")
	fmt.Print("Best: ")
	HeapSortWithCounter(sorted)
	fmt.Print("Avg:  ")
	HeapSortWithCounter(random)
	fmt.Print("Worst:")
	HeapSortWithCounter(reversed)
	fmt.Println()

	// Bucket (целые 0..999, 10 бакетов, внутри insertion)
	fmt.Println("БЛОЧНАЯ (Bucket):")
	fmt.Print("Best: ")
	BucketSortWithCounter(sorted)
	fmt.Print("Avg:  ")
	BucketSortWithCounter(random)
	fmt.Print("Worst:")
	BucketSortWithCounter(reversed)
	fmt.Println()

	// Radix (LSD base 10)
	fmt.Println("ПОРАЗРЯДНАЯ (Radix):")
	fmt.Print("Best: ")
	RadixSortWithCounter(sorted)
	fmt.Print("Avg:  ")
	RadixSortWithCounter(random)
	fmt.Print("Worst:")
	RadixSortWithCounter(reversed)
}

// Генерация случайного массива в диапазоне [0,1000)
func GenerateSlice(size int) []int {
	if size <= 0 {
		size = 100
	}
	slice := make([]int, 0, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice = append(slice, rand.Intn(1000))
	}
	return slice
}

// ----------------- Bubble Sort -----------------
// BUBBLE SORT:
// T (лучший)  = N * 5
// T (средний) = N^2 * 5
// T (худший)  = N^2 * 5
//
// В лучшем случае O(n) - массив уже отсортирован, выход по swapped=false
// В худшем случае O(n^2) - массив отсортирован в обратном порядке
// В среднем O(n^2) - массив не отсортирован
func BubbleSort(slice []int) []int {
	var counter int
	a := append([]int{}, slice...)
	counter++ // copy

	n := len(a)
	counter++
	for i := 0; i < n; i++ {
		counter++ // for i compare
		swapped := false
		counter++ // assign
		for j := 0; j < n-i-1; j++ {
			counter++ // for j compare
			if a[j] > a[j+1] {
				counter++ // comparison true
				a[j], a[j+1] = a[j+1], a[j]
				counter += 3 // swap
				swapped = true
				counter++
			} else {
				counter++ // comparison false path
			}
			counter++ // j++
		}
		if !swapped {
			counter++ // break condition
			break
		}
		counter++ // i++
	}
	counter++ // end
	fmt.Println(counter)
	return a
}

// ----------------- Insertion Sort -----------------
// INSERTION SORT:
// T (лучший)  = N * 3
// T (средний) = N^2 * 4
// T (худший)  = N^2 * 4
//
// В лучшем случае O(n) - массив уже отсортирован, минимум сдвигов
// В худшем случае O(n^2) - массив отсортирован в обратном порядке, максимум сдвигов
// В среднем O(n^2) - массив не отсортирован
func InsertionSort(slice []int) []int {
	var counter int
	a := append([]int{}, slice...)
	counter++

	for i := 1; i < len(a); i++ {
		counter++
		key := a[i]
		counter++
		j := i - 1
		counter++
		for j >= 0 && a[j] > key {
			counter += 2
			a[j+1] = a[j]
			counter++
			j--
			counter++
		}
		if j >= 0 {
			counter++
		}
		a[j+1] = key
		counter++
		counter++ // i++
	}
	counter++
	fmt.Println(counter)
	return a
}

// ----------------- Selection Sort -----------------
// SELECTION SORT:
// T (лучший)  = N^2 * 3
// T (средний) = N^2 * 3
// T (худший)  = N^2 * 3
//
// Во всех случаях O(n^2) - всегда ищет минимум в оставшейся части
// В лучшем случае O(n^2) - не зависит от входных данных
// В худшем случае O(n^2) - не зависит от входных данных
// В среднем O(n^2) - не зависит от входных данных
func SelectionSort(slice []int) []int {
	var counter int
	a := append([]int{}, slice...)
	counter++

	n := len(a)
	counter++
	for i := 0; i < n-1; i++ {
		counter++
		minIdx := i
		counter++
		for j := i + 1; j < n; j++ {
			counter++
			if a[j] < a[minIdx] {
				counter++
				minIdx = j
				counter++
			} else {
				counter++
			}
			counter++
		}
		if minIdx != i {
			counter++
			a[i], a[minIdx] = a[minIdx], a[i]
			counter += 3
		} else {
			counter++
		}
		counter++
	}
	counter++
	fmt.Println(counter)
	return a
}

// ----------------- Quick Sort -----------------
// QUICK SORT:
// T (лучший)  = N * log(N) * 10
// T (средний) = N * log(N) * 10
// T (худший)  = N^2 * 8
//
// В лучшем случае O(n log n) - удачное разбиение пополам
// В худшем случае O(n^2) - неудачный выбор опорного элемента
// В среднем O(n log n) - хорошее разбиение в среднем
func QuickSortWithCounter(slice []int) []int {
	a := append([]int{}, slice...)
	counter := 0
	a, counter = quickRec(a, &counter)
	fmt.Println(counter)
	return a
}

func quickRec(a []int, counter *int) ([]int, int) {
	*counter++
	if len(a) <= 1 {
		*counter++
		return a, *counter
	}
	pivot := a[len(a)/2]
	*counter++
	less := make([]int, 0, len(a))
	more := make([]int, 0, len(a))
	eq := make([]int, 0, 1)
	*counter += 3

	for i := 0; i < len(a); i++ {
		*counter++
		if a[i] < pivot {
			*counter++
			less = append(less, a[i])
			*counter++
		} else if a[i] > pivot {
			*counter += 2
			more = append(more, a[i])
			*counter++
		} else {
			*counter += 2
			eq = append(eq, a[i])
			*counter++
		}
		*counter++
	}

	var c1, c2 int
	less, c1 = quickRec(less, counter)
	*counter += c1
	more, c2 = quickRec(more, counter)
	*counter += c2

	out := append(append(less, eq...), more...)
	*counter++
	return out, *counter
}

// ----------------- Merge Sort -----------------
// MERGE SORT:
// T (лучший)  = N * log(N) * 12
// T (средний) = N * log(N) * 12
// T (худший)  = N * log(N) * 12
//
// Во всех случаях O(n log n) - стабильная сложность
// В лучшем случае O(n log n) - всегда делит пополам и сливает
// В худшем случае O(n log n) - всегда делит пополам и сливает
// В среднем O(n log n) - всегда делит пополам и сливает
func MergeSortWithCounter(slice []int) []int {
	a := append([]int{}, slice...)
	counter := 0
	a, counter = mergeRec(a, &counter)
	fmt.Println(counter)
	return a
}

func mergeRec(a []int, counter *int) ([]int, int) {
	*counter++
	if len(a) <= 1 {
		*counter++
		return a, *counter
	}
	mid := len(a) / 2
	*counter++
	left, _ := mergeRec(a[:mid], counter)
	right, _ := mergeRec(a[mid:], counter)

	out := make([]int, 0, len(a))
	*counter++
	i, j := 0, 0
	*counter += 2
	for i < len(left) && j < len(right) {
		*counter++
		if left[i] <= right[j] {
			*counter++
			out = append(out, left[i])
			*counter++
			i++
			*counter++
		} else {
			*counter++
			out = append(out, right[j])
			*counter++
			j++
			*counter++
		}
		*counter++
	}
	*counter++
	for i < len(left) {
		*counter++
		out = append(out, left[i])
		*counter++
		i++
		*counter++
	}
	*counter++
	for j < len(right) {
		*counter++
		out = append(out, right[j])
		*counter++
		j++
		*counter++
	}
	*counter++
	return out, *counter
}

// ----------------- Heap Sort -----------------
// HEAP SORT:
// T (лучший)  = N * log(N) * 8
// T (средний) = N * log(N) * 8
// T (худший)  = N * log(N) * 8
//
// Во всех случаях O(n log n) - стабильная сложность
// В лучшем случае O(n log n) - построение кучи O(n) + n извлечений O(log n)
// В худшем случае O(n log n) - построение кучи O(n) + n извлечений O(log n)
// В среднем O(n log n) - построение кучи O(n) + n извлечений O(log n)
func HeapSortWithCounter(slice []int) []int {
	a := append([]int{}, slice...)
	counter := 0

	n := len(a)
	counter++
	// Build heap
	for i := n/2 - 1; i >= 0; i-- {
		counter++
		counter = heapify(a, n, i, counter)
	}
	counter++

	// Extract elements
	for i := n - 1; i > 0; i-- {
		counter++
		a[0], a[i] = a[i], a[0] // Исправлено!
		counter += 3
		counter = heapify(a, i, 0, counter)
	}
	counter++
	fmt.Println(counter)
	return a
}

func heapify(a []int, n, i, counter int) int {
	largest := i
	counter++
	l := 2*i + 1
	r := 2*i + 2
	counter += 2

	if l < n && a[l] > a[largest] {
		counter += 2
		largest = l
		counter++
	} else {
		counter++
	}
	if r < n && a[r] > a[largest] {
		counter += 2
		largest = r
		counter++
	} else {
		counter++
	}
	if largest != i {
		counter++
		a[i], a[largest] = a[largest], a[i]
		counter += 3
		counter = heapify(a, n, largest, counter)
	} else {
		counter++
	}
	return counter
}

// ----------------- Bucket Sort (для int, base=10) -----------------
// BUCKET SORT:
// T (лучший)  = N + K * 6
// T (средний) = N + K * 6
// T (худший)  = N^2 * 4
//
// В лучшем случае O(n+k) - равномерное распределение по бакетам
// В худшем случае O(n^2) - все элементы попадают в один бакет
// В среднем O(n+k) - хорошее распределение по бакетам
func BucketSortWithCounter(slice []int) []int {
	a := append([]int{}, slice...)
	counter := 0

	if len(a) == 0 {
		counter++
		fmt.Println(counter)
		return a
	}
	counter++
	// max
	max := a[0] // Исправлено!
	counter++
	for i := 1; i < len(a); i++ {
		counter++
		if a[i] > max {
			counter++
			max = a[i]
			counter++
		} else {
			counter++
		}
		counter++
	}
	counter++

	bucketsCount := 10
	counter++
	buckets := make([][]int, bucketsCount)
	counter++

	// Раскидываем по бакетам
	for _, v := range a {
		counter++
		idx := (v * (bucketsCount - 1)) / (max + 1)
		counter++
		buckets[idx] = append(buckets[idx], v)
		counter++
	}

	// Сортируем каждый бакет вставками (с тем же счётчиком)
	out := make([]int, 0, len(a))
	counter++
	for i := 0; i < bucketsCount; i++ {
		counter++
		if len(buckets[i]) > 1 {
			counter++
			buckets[i], counter = insertionCount(buckets[i], counter)
		} else {
			counter++
		}
		out = append(out, buckets[i]...)
		counter++
	}
	counter++
	fmt.Println(counter)
	return out
}

func insertionCount(a []int, counter int) ([]int, int) {
	for i := 1; i < len(a); i++ {
		counter++
		key := a[i]
		counter++
		j := i - 1
		counter++
		for j >= 0 && a[j] > key {
			counter += 2
			a[j+1] = a[j]
			counter++
			j--
			counter++
		}
		if j >= 0 {
			counter++
		}
		a[j+1] = key
		counter++
		counter++
	}
	counter++
	return a, counter
}

// ----------------- Radix Sort (LSD base 10) -----------------
// RADIX SORT:
// T (лучший)  = N * K * 15
// T (средний) = N * K * 15
// T (худший)  = N * K * 15
//
// Во всех случаях O(n*k) где k - количество разрядов
// В лучшем случае O(n*k) - не зависит от порядка элементов
// В худшем случае O(n*k) - не зависит от порядка элементов
// В среднем O(n*k) - не зависит от порядка элементов
func RadixSortWithCounter(slice []int) []int {
	a := append([]int{}, slice...)
	counter := 0

	if len(a) == 0 {
		counter++
		fmt.Println(counter)
		return a
	}
	counter++
	// max
	max := a[0] // Исправлено!
	counter++
	for i := 1; i < len(a); i++ {
		counter++
		if a[i] > max {
			counter++
			max = a[i]
			counter++
		} else {
			counter++
		}
		counter++
	}
	counter++

	for exp := 1; max/exp > 0; exp *= 10 {
		counter += 2
		a, counter = countingByDigit(a, exp, counter)
	}
	counter++
	fmt.Println(counter)
	return a
}

func countingByDigit(a []int, exp int, counter int) ([]int, int) {
	n := len(a)
	counter++
	out := make([]int, n)
	count := make([]int, 10)
	counter += 2

	// count freq
	for i := 0; i < n; i++ {
		counter++
		d := (a[i] / exp) % 10
		counter++
		count[d]++
		counter++
	}
	counter++

	// prefix sums
	for i := 1; i < 10; i++ {
		counter++
		count[i] += count[i-1]
		counter++
	}
	counter++

	// stable placement (reverse)
	for i := n - 1; i >= 0; i-- {
		counter++
		d := (a[i] / exp) % 10
		counter++
		out[count[d]-1] = a[i]
		counter++
		count[d]--
		counter++
	}
	counter++

	// copy back
	for i := 0; i < n; i++ {
		counter++
		a[i] = out[i]
		counter++
	}
	counter++
	return a, counter
}
