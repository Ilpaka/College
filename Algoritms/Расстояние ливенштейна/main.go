package main

import "fmt"

const string_1 = "лабрадор"
const string_2 = "гибралтар"

func main() {
	fmt.Println(Levenshtein(string_1, string_2))
}

func min(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func Levenshtein(str1, str2 string) int {
	runes1 := []rune(str1)
	runes2 := []rune(str2)

	n := len(runes1)
	m := len(runes2)

	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		matrix[i][0] = i
	}

	for j := 0; j <= m; j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			var cost int
			if runes1[i-1] == runes2[j-1] {
				cost = 0
			} else {
				cost = 1
			}

			// Применяем формулу:
			// D(i,j) = min{
			//     D(i,j-1) + 1,           // вставка
			//     D(i-1,j) + 1,           // удаление
			//     D(i-1,j-1) + m(S1[i],S2[j])  // замена
			// }
			matrix[i][j] = min(
				matrix[i][j-1]+1,      // вставка
				matrix[i-1][j]+1,      // удаление
				matrix[i-1][j-1]+cost, // замена с m()
			)
		}
	}

	fmt.Println("Матрица D(i,j):")
	for i := 0; i <= n; i++ {
		fmt.Println(matrix[i])
	}

	return matrix[n][m]
}
