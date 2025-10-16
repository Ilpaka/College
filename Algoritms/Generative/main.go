package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	BASE_STRING     = "the rule is good"
	CHARSET         = "abcdefghijklmnopqrstuvwxyz "
	MAX_GENERATIONS = 100
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// individuals := []string{"hfksfbnrugdkfnt ", "lfjgntcgasdtrcjd", "pqwlrbc dhfyriem"}

	individuals := generate_rand_individs(16, CHARSET, 100)

	fmt.Printf("🎯 Целевая строка: '%s'\n", BASE_STRING)
	fmt.Printf("📏 Длина: %d символов\n\n", len(BASE_STRING))

	population := create_population(individuals)

	fmt.Println("🧬 Начальная популяция:")
	for i, ind := range population {
		fmt.Printf("  %d: '%s'\n", i+1, ind)
	}
	fmt.Println()

	generation := 0
	bestDistance := math.MaxInt
	bestIndivid := ""

	for generation < MAX_GENERATIONS {
		generation++

		distanceMap := distance_haminga(BASE_STRING, population)

		currentBests, currentDistance := sort_the_best_individs(distanceMap)

		currentBest := currentBests[0]

		if currentDistance < bestDistance {
			bestDistance = currentDistance
			bestIndivid = currentBest
		}

		if generation%100 == 0 {
			fmt.Printf("Поколение %4d | Лучшая: '%s' | Расстояние: %2d\n",
				generation, currentBest, currentDistance)
		}

		if currentDistance == 0 {
			fmt.Printf("\n🎉 ЭВРИКА! Решение найдено за %d поколений!\n", generation)
			fmt.Printf("🏆 Результат: '%s'\n", currentBest)
			fmt.Printf("✅ Расстояние Хэмминга: %d (идеально!)\n", currentDistance)
			return
		}

		population = mutation(currentBest, CHARSET)
	}

	fmt.Printf("\n⏰ Достигнуто максимальное количество поколений: %d\n", MAX_GENERATIONS)
	fmt.Printf("🥈 Лучший результат: '%s'\n", bestIndivid)
	fmt.Printf("📊 Расстояние Хэмминга: %d\n", bestDistance)
}

func generate_rand_individs(length int, charset string, quantity int) []string {
	individs := make([]string, quantity)

	for i := 0; i < quantity; i++ {
		individ := make([]byte, length)

		for j := 0; j < length; j++ {
			individ[j] = charset[rand.Intn(len(charset))]
		}

		individs[i] = string(individ)
	}

	return individs
}

func create_population(individuals []string) []string {
	population := make([]string, len(individuals))

	for i := range individuals {
		population[i] = strings.ToLower(individuals[i])
	}

	return population
}

func distance_haminga(base_string string, s2 []string) map[int]string {
	table_of_distance := make(map[int]string)

	for i := 0; i < len(s2); i++ {
		distance := 0
		individ := s2[i]

		for j := 0; j < len(base_string); j++ {
			if individ[j] != base_string[j] {
				distance++
			}
		}

		key := distance
		_, exists := table_of_distance[key]
		for exists {
			key += 1000
			_, exists = table_of_distance[key]
		}

		table_of_distance[key] = individ
	}

	return table_of_distance
}

func sort_the_best_individs(table_of_distance map[int]string) ([]string, int) {
	best_individs := make([]string, 2)
	minKey := math.MaxInt
	secondMinKey := math.MaxInt

	for key := range table_of_distance {
		realKey := key % 1000
		if realKey < minKey {
			minKey = realKey
		}
	}

	for key := range table_of_distance {
		realKey := key % 1000
		if realKey == minKey {
			best_individs[0] = table_of_distance[key]
		} else if realKey < secondMinKey {
			secondMinKey = realKey
		}
	}

	for key := range table_of_distance {
		realKey := key % 1000
		if realKey == secondMinKey {
			best_individs[1] = table_of_distance[key]
			break
		}
	}

	return best_individs, minKey
}

func mutation(individual string, charset string) []string {
	mutants := make([]string, 1000)

	for i := 0; i < 1000; i++ {
		mutant := []byte(individual)

		position := rand.Intn(len(mutant))

		mutant[position] = charset[rand.Intn(len(charset))]

		mutants[i] = string(mutant)
	}

	return mutants
}

// package main

// import (
// 	"fmt"
// 	"math"
// 	"math/rand"
// 	"strings"
// 	"sync"
// 	"time"
// )

// const (
// 	BASE_STRING     = "the rule is good"
// 	CHARSET         = "abcdefghijklmnopqrstuvwxyz "
// 	MAX_GENERATIONS = 100
// 	NUM_WORKERS     = 8 // Количество параллельных воркеров
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	individuals := generate_rand_individs(16, CHARSET, 100)

// 	fmt.Printf("🎯 Целевая строка: '%s'\n", BASE_STRING)
// 	fmt.Printf("📏 Длина: %d символов\n\n", len(BASE_STRING))

// 	population := create_population(individuals)

// 	fmt.Println("🧬 Начальная популяция (первые 10):")
// 	for i := 0; i < 10 && i < len(population); i++ {
// 		fmt.Printf("  %d: '%s'\n", i+1, population[i])
// 	}
// 	fmt.Println()

// 	generation := 0
// 	bestDistance := math.MaxInt
// 	bestIndivid := ""

// 	for generation < MAX_GENERATIONS {
// 		generation++

// 		// 🔥 ПАРАЛЛЕЛЬНОЕ вычисление расстояния Хэмминга
// 		distanceMap := distance_haminga_parallel(BASE_STRING, population)

// 		currentBests, currentDistance := sort_the_best_individs(distanceMap)

// 		currentBest := currentBests[0]

// 		if currentDistance < bestDistance {
// 			bestDistance = currentDistance
// 			bestIndivid = currentBest
// 		}

// 		if generation%10 == 0 {
// 			fmt.Printf("Поколение %4d | Лучшая: '%s' | Расстояние: %2d\n",
// 				generation, currentBest, currentDistance)
// 		}

// 		if currentDistance == 0 {
// 			fmt.Printf("\n🎉 ЭВРИКА! Решение найдено за %d поколений!\n", generation)
// 			fmt.Printf("🏆 Результат: '%s'\n", currentBest)
// 			fmt.Printf("✅ Расстояние Хэмминга: %d (идеально!)\n", currentDistance)
// 			return
// 		}

// 		// 🔥 ПАРАЛЛЕЛЬНАЯ мутация
// 		population = mutation_parallel(currentBest, CHARSET, 1000)
// 	}

// 	fmt.Printf("\n⏰ Достигнуто максимальное количество поколений: %d\n", MAX_GENERATIONS)
// 	fmt.Printf("🥈 Лучший результат: '%s'\n", bestIndivid)
// 	fmt.Printf("📊 Расстояние Хэмминга: %d\n", bestDistance)
// }

// // 🔥 ПАРАЛЛЕЛЬНАЯ генерация случайных индивидов
// func generate_rand_individs(length int, charset string, quantity int) []string {
// 	individs := make([]string, quantity)
// 	var wg sync.WaitGroup

// 	// Разбиваем на chunks для параллельной обработки
// 	chunkSize := quantity / NUM_WORKERS
// 	if chunkSize == 0 {
// 		chunkSize = 1
// 	}

// 	for i := 0; i < quantity; i += chunkSize {
// 		wg.Add(1)

// 		start := i
// 		end := i + chunkSize
// 		if end > quantity {
// 			end = quantity
// 		}

// 		// Запускаем горутину для каждого chunk
// 		go func(start, end int) {
// 			defer wg.Done()

// 			// Создаём локальный рандом для каждой горутины (thread-safe)
// 			localRand := rand.New(rand.NewSource(time.Now().UnixNano() + int64(start)))

// 			for j := start; j < end; j++ {
// 				individ := make([]byte, length)
// 				for k := 0; k < length; k++ {
// 					individ[k] = charset[localRand.Intn(len(charset))]
// 				}
// 				individs[j] = string(individ)
// 			}
// 		}(start, end)
// 	}

// 	wg.Wait() // Ждём завершения всех горутин
// 	return individs
// }

// func create_population(individuals []string) []string {
// 	population := make([]string, len(individuals))

// 	for i := range individuals {
// 		population[i] = strings.ToLower(individuals[i])
// 	}

// 	return population
// }

// // 🔥 ПАРАЛЛЕЛЬНОЕ вычисление расстояния Хэмминга
// func distance_haminga_parallel(base_string string, s2 []string) map[int]string {
// 	table_of_distance := make(map[int]string)
// 	var mu sync.Mutex // Мьютекс для защиты map от одновременного доступа
// 	var wg sync.WaitGroup

// 	// Канал для передачи результатов
// 	type result struct {
// 		distance int
// 		individ  string
// 	}
// 	resultChan := make(chan result, len(s2))

// 	// Разбиваем популяцию на chunks
// 	chunkSize := len(s2) / NUM_WORKERS
// 	if chunkSize == 0 {
// 		chunkSize = 1
// 	}

// 	for i := 0; i < len(s2); i += chunkSize {
// 		wg.Add(1)

// 		start := i
// 		end := i + chunkSize
// 		if end > len(s2) {
// 			end = len(s2)
// 		}

// 		// Горутина для обработки chunk популяции
// 		go func(start, end int) {
// 			defer wg.Done()

// 			for j := start; j < end; j++ {
// 				distance := 0
// 				individ := s2[j]

// 				// Вычисляем расстояние Хэмминга
// 				for k := 0; k < len(base_string); k++ {
// 					if individ[k] != base_string[k] {
// 						distance++
// 					}
// 				}

// 				// Отправляем результат в канал
// 				resultChan <- result{distance: distance, individ: individ}
// 			}
// 		}(start, end)
// 	}

// 	// Горутина для закрытия канала после завершения всех воркеров
// 	go func() {
// 		wg.Wait()
// 		close(resultChan)
// 	}()

// 	// Собираем результаты из канала
// 	for res := range resultChan {
// 		mu.Lock() // Лочим мьютекс перед записью в map

// 		key := res.distance
// 		_, exists := table_of_distance[key]
// 		for exists {
// 			key += 1000
// 			_, exists = table_of_distance[key]
// 		}
// 		table_of_distance[key] = res.individ

// 		mu.Unlock() // Разлочим мьютекс
// 	}

// 	return table_of_distance
// }

// func sort_the_best_individs(table_of_distance map[int]string) ([]string, int) {
// 	best_individs := make([]string, 2)
// 	minKey := math.MaxInt
// 	secondMinKey := math.MaxInt

// 	for key := range table_of_distance {
// 		realKey := key % 1000
// 		if realKey < minKey {
// 			minKey = realKey
// 		}
// 	}

// 	for key := range table_of_distance {
// 		realKey := key % 1000
// 		if realKey == minKey {
// 			best_individs[0] = table_of_distance[key]
// 		} else if realKey < secondMinKey {
// 			secondMinKey = realKey
// 		}
// 	}

// 	for key := range table_of_distance {
// 		realKey := key % 1000
// 		if realKey == secondMinKey {
// 			best_individs[1] = table_of_distance[key]
// 			break
// 		}
// 	}

// 	return best_individs, minKey
// }

// // 🔥 ПАРАЛЛЕЛЬНАЯ мутация
// func mutation_parallel(individual string, charset string, count int) []string {
// 	mutants := make([]string, count)
// 	var wg sync.WaitGroup

// 	// Разбиваем на chunks
// 	chunkSize := count / NUM_WORKERS
// 	if chunkSize == 0 {
// 		chunkSize = 1
// 	}

// 	for i := 0; i < count; i += chunkSize {
// 		wg.Add(1)

// 		start := i
// 		end := i + chunkSize
// 		if end > count {
// 			end = count
// 		}

// 		// Горутина для генерации мутантов
// 		go func(start, end int) {
// 			defer wg.Done()

// 			// Локальный рандом для thread-safety
// 			localRand := rand.New(rand.NewSource(time.Now().UnixNano() + int64(start)))

// 			for j := start; j < end; j++ {
// 				mutant := []byte(individual)
// 				position := localRand.Intn(len(mutant))
// 				mutant[position] = charset[localRand.Intn(len(charset))]
// 				mutants[j] = string(mutant)
// 			}
// 		}(start, end)
// 	}

// 	wg.Wait() // Ждём завершения всех горутин
// 	return mutants
// }
