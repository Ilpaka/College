package main

import (
	"runtime"
	"testing"
	"time"
)

// ============================================
// БЕНЧМАРКИ ДЛЯ ОТДЕЛЬНЫХ ФУНКЦИЙ
// ============================================

// Бенчмарк для функции create_population
func BenchmarkCreatePopulation(b *testing.B) {
	individuals := []string{"hfksfbnrugdkfnt ", "lfjgntcgasdtrcjd", "pqwlrbc dhfyriem"}

	b.ReportAllocs() // Включаем отчет по аллокациям памяти
	b.ResetTimer()   // Сбрасываем таймер после инициализации

	for i := 0; i < b.N; i++ {
		_ = create_population(individuals)
	}
}

// Бенчмарк для функции distance_haminga
func BenchmarkDistanceHaminga(b *testing.B) {
	population := []string{"hfksfbnrugdkfnt ", "lfjgntcgasdtrcjd", "pqwlrbc dhfyriem"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = distance_haminga(BASE_STRING, population)
	}
}

// Бенчмарк для функции sort_the_best_individs
func BenchmarkSortTheBestIndivids(b *testing.B) {
	// Создаём тестовую map
	testMap := map[int]string{
		5:    "test string one ",
		10:   "test string two ",
		3:    "test string thr ",
		1003: "duplicate dist 3",
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = sort_the_best_individs(testMap)
	}
}

// Бенчмарк для функции mutation
func BenchmarkMutation(b *testing.B) {
	individual := "the rule is good"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = mutation(individual, CHARSET)
	}
}

// ============================================
// БЕНЧМАРК ДЛЯ ВСЕГО ГЕНЕТИЧЕСКОГО АЛГОРИТМА
// ============================================

// Бенчмарк полного цикла алгоритма (малое количество поколений)
func BenchmarkGeneticAlgorithmSmall(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		runGeneticAlgorithm(10) // 10 поколений
	}
}

// Бенчмарк полного цикла алгоритма (среднее количество поколений)
func BenchmarkGeneticAlgorithmMedium(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		runGeneticAlgorithm(100) // 100 поколений
	}
}

// Бенчмарк полного цикла алгоритма (большое количество поколений)
func BenchmarkGeneticAlgorithmLarge(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		runGeneticAlgorithm(1000) // 1000 поколений
	}
}

// ============================================
// ТЕСТЫ ПАМЯТИ С ДЕТАЛЬНОЙ СТАТИСТИКОЙ
// ============================================

// Тест памяти для всего алгоритма
func TestMemoryUsage(t *testing.T) {
	// Принудительно запускаем GC перед тестом
	runtime.GC()

	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// Запускаем алгоритм
	runGeneticAlgorithm(100)

	runtime.ReadMemStats(&m2)

	// Статистика
	allocatedMemory := m2.TotalAlloc - m1.TotalAlloc
	currentMemory := m2.Alloc - m1.Alloc
	numGC := m2.NumGC - m1.NumGC

	t.Logf("📊 Статистика памяти для 100 поколений:")
	t.Logf("   💾 Всего выделено памяти: %d байт (%.2f MB)", allocatedMemory, float64(allocatedMemory)/1024/1024)
	t.Logf("   🔥 Текущая память: %d байт (%.2f KB)", currentMemory, float64(currentMemory)/1024)
	t.Logf("   🗑️  Количество GC циклов: %d", numGC)
	t.Logf("   📦 Heap объектов: %d", m2.HeapObjects)
}

// Тест памяти для функции mutation
func TestMutationMemory(t *testing.T) {
	runtime.GC()

	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	individual := "the rule is good"
	iterations := 1000

	for i := 0; i < iterations; i++ {
		_ = mutation(individual, CHARSET)
	}

	runtime.ReadMemStats(&m2)

	allocatedMemory := m2.TotalAlloc - m1.TotalAlloc

	t.Logf("📊 Статистика памяти для mutation (%d итераций):", iterations)
	t.Logf("   💾 Всего выделено: %d байт (%.2f MB)", allocatedMemory, float64(allocatedMemory)/1024/1024)
	t.Logf("   📈 На одну итерацию: %.2f KB", float64(allocatedMemory)/float64(iterations)/1024)
}

// ============================================
// ТЕСТ ПРОИЗВОДИТЕЛЬНОСТИ С ТАЙМЕРОМ
// ============================================

// Тест скорости выполнения алгоритма
func TestPerformanceSpeed(t *testing.T) {
	tests := []struct {
		name        string
		generations int
	}{
		{"10 поколений", 10},
		{"100 поколений", 100},
		{"1000 поколений", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()

			runGeneticAlgorithm(tt.generations)

			duration := time.Since(start)

			t.Logf("⏱️  %s выполнено за: %v", tt.name, duration)
			t.Logf("   📊 Среднее время на поколение: %v", duration/time.Duration(tt.generations))
		})
	}
}

// ============================================
// ВСПОМОГАТЕЛЬНАЯ ФУНКЦИЯ ДЛЯ ЗАПУСКА АЛГОРИТМА
// ============================================

// Функция запуска генетического алгоритма (без вывода в консоль)
func runGeneticAlgorithm(maxGenerations int) (string, int) {
	individuals := []string{"hfksfbnrugdkfnt ", "lfjgntcgasdtrcjd", "pqwlrbc dhfyriem"}
	population := create_population(individuals)

	generation := 0
	bestDistance := 999999
	bestIndivid := ""

	for generation < maxGenerations {
		generation++

		distanceMap := distance_haminga(BASE_STRING, population)
		currentBests, currentDistance := sort_the_best_individs(distanceMap)
		currentBest := currentBests[0]

		if currentDistance < bestDistance {
			bestDistance = currentDistance
			bestIndivid = currentBest
		}

		if currentDistance == 0 {
			return currentBest, generation
		}

		population = mutation(currentBest, CHARSET)
	}

	return bestIndivid, generation
}

// ============================================
// ПРОДВИНУТЫЙ БЕНЧМАРК С ПРОФИЛИРОВАНИЕМ
// ============================================

// Бенчмарк с детальной информацией о CPU
func BenchmarkWithCPUProfile(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		individuals := []string{"hfksfbnrugdkfnt ", "lfjgntcgasdtrcjd", "pqwlrbc dhfyriem"}
		population := create_population(individuals)
		b.StartTimer()

		for gen := 0; gen < 100; gen++ {
			distanceMap := distance_haminga(BASE_STRING, population)
			currentBests, _ := sort_the_best_individs(distanceMap)
			population = mutation(currentBests[0], CHARSET)
		}
	}
}
