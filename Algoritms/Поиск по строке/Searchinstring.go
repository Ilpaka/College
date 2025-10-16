package main

import "fmt"

func main() {
	text := "ababcabcabc"
	pattern := "abcab"
	pos := KMP(pattern, text)
	fmt.Println("Позиция:", pos)

	text = "Пример строки для поиска подстроки"
	pattern = "поиска"
	pos = RabinKarp(text, pattern)
	if pos != -1 {
		fmt.Printf("Подстрока '%s' найдена в позиции %d\n", pattern, pos)
	} else {
		fmt.Printf("Подстрока '%s' не найдена\n", pattern)
	}
}

func KMP(pattern, text string) int {
	n := len(text)
	m := len(pattern)

	if m == 0 {
		return 0
	}

	table := make([]int, m)
	table[0] = 0
	shift := 0

	for q := 1; q < m; q++ {
		for shift > 0 && pattern[shift] != pattern[q] {
			shift = table[shift-1]
		}
		if pattern[shift] == pattern[q] {
			shift++
		}
		table[q] = shift
	}

	shift = 0
	for i := 0; i < n; i++ {
		for shift > 0 && pattern[shift] != text[i] {
			shift = table[shift-1]
		}
		if pattern[shift] == text[i] {
			shift++
		}
		if shift == m {
			return i - m + 1
		}
	}
	return -1
}

func RabinKarp(text, pattern string) int {
	n := len(text)
	m := len(pattern)
	if m > n || m == 0 {
		return -1
	}

	var q int64 = 9973
	var h int64 = 1
	for i := 1; i < m; i++ {
		h = (h << 8) % q
	}

	var p int64 = 0
	var t int64 = 0
	for k := 0; k < m; k++ {
		p = ((p << 8) + int64(pattern[k])) % q
		t = ((t << 8) + int64(text[k])) % q
	}

	for i := 0; i <= n-m; i++ {
		if p == t {
			found := true
			for j := 0; j < m; j++ {
				if text[i+j] != pattern[j] {
					found = false
					break
				}
			}
			if found {
				return i
			}
		}
		if i < n-m {
			t = (t - int64(text[i])*h) % q
			if t < 0 {
				t += q
			}
			t = ((t << 8) + int64(text[i+m])) % q
		}
	}

	return -1
}

func BoyerMoore(text, pattern string) int {
	n := len(text)
	m := len(pattern)
	if m == 0 || m > n {
		return -1
	}
	shLen := 256

	hash := func(c byte) int { return int(c & 0xFF) }

	shifts := make([]int, shLen)
	for i := 0; i < shLen; i++ {
		shifts[i] = m
	}
	for i := 0; i < m-1; i++ {
		shifts[hash(pattern[i])] = m - i - 1
	}

	i := m - 1
	for i < n {
		j := m - 1
		for j >= 0 && text[i] == pattern[j] {
			if j == 0 {
				return i
			}
			i--
			j--
		}
		if i < n-1 {
			i += shifts[hash(text[i+1])]
		} else {
			break
		}
	}
	return -1
}
