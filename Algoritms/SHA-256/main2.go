package main

import (
	"fmt"
	"strconv"
	"strings"
)

// структура для создания перемнной
type cryptographer struct {
	word string
}

// метод который вызывает все функции
func (c *cryptographer) sha256(word string) string {
	first_step := convert_to_byte(word)
	second_step := plus_one(first_step)
	third_step := add_zero(second_step, len(word)*8)

	blocks := split_into_blocks(third_step)

	hashState := H0

	for i, block_bits := range blocks {
		fourth_step := parseBlock(block_bits)

		W := ExpansionBlock(fourth_step)

		compressed := Compress64(W, hashState)

		hashState = add_hash_states(hashState, compressed)

		fmt.Printf("Блок %d обработан\n", i+1)
	}

	result := convert_to_16(hashState)

	return result
}

// костанты
var K = [64]uint32{
	0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5,
	0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3,
	0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc,
	0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7,
	0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13,
	0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3,
	0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5,
	0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208,
	0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
}

// для в кубе
var H0 = [8]uint32{
	0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a,
	0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19,
}

// базовый вызов
func main() {
	one_String := cryptographer{word: "hello world"}
	hash1 := one_String.sha256(one_String.word)

	twoo_String := cryptographer{word: "more symbols than can be contained in block"}
	hash2 := twoo_String.sha256(twoo_String.word)

	fmt.Println("Получен:", hash1)
	fmt.Println("Получен для второй строки:", hash2)
}

// перефод строки в двоичную
func convert_to_byte(word string) string {
	eight_byte_String := "" //пустая стркоа

	for _, b := range []byte(word) {
		temp := fmt.Sprintf("%08b", b)
		eight_byte_String += temp //перебираем и добавляем
	}
	return eight_byte_String
}

// +1
func plus_one(bit_string string) string {
	bit_string += "1"
	return bit_string
}

// добавялем нули
func add_zero(bit_String string, originalLengthBits int) string {
	current_len := len(bit_String)
	//(448, потому что ещё 64 бита займёт длина сообщения
	// +512 в формуле гарантирует положительный результат
	k := (448 - current_len%512 + 512) % 512 //выщитываем в зависимости от отстатка, чтобы большие строки не ломались

	bit_String += strings.Repeat("0", k) //добавялем нули

	lengthBits := fmt.Sprintf("%064b", originalLengthBits) //переводим
	bit_String += lengthBits

	return bit_String
}

// разбиваем на блоки по 512
func split_into_blocks(bitString string) []string {
	var blocks []string
	for i := 0; i < len(bitString); i += 512 {
		end := i + 512
		if end > len(bitString) {
			end = len(bitString)
		}
		blocks = append(blocks, bitString[i:end])
	}
	return blocks
}

// Берём по одному блоку и сохранем в массив uint32 получиться 16 слов по 32
func parseBlock(block string) []uint32 {
	words := make([]uint32, 16)

	for i := 0; i < 16; i++ {
		start := i * 32
		end := start + 32

		wordBits := block[start:end]
		val, _ := strconv.ParseUint(wordBits, 2, 32)
		words[i] = uint32(val)
	}

	return words
}

// снова объединяем в одну стркоу
func add_hash_states(h1, h2 [8]uint32) [8]uint32 {
	var result [8]uint32
	for i := 0; i < 8; i++ {
		result[i] = h1[i] + h2[i]
	}
	return result
}

func RightRotate32(n uint32, d uint32) uint32 {
	return (n >> d) | (n << (32 - d))
}

func C0(x uint32) uint32 {
	return RightRotate32(x, 7) ^ RightRotate32(x, 18) ^ (x >> 3)
}

func C1(x uint32) uint32 {
	return RightRotate32(x, 17) ^ RightRotate32(x, 19) ^ (x >> 10)
}

func S0(x uint32) uint32 {
	return RightRotate32(x, 2) ^ RightRotate32(x, 13) ^ RightRotate32(x, 22)
}

func S1(x uint32) uint32 {
	return RightRotate32(x, 6) ^ RightRotate32(x, 11) ^ RightRotate32(x, 25)
}

func ch(e, f, g uint32) uint32 {
	return (e & f) ^ ((^e) & g)
}

func maj(a, b, c uint32) uint32 {
	return (a & b) ^ (a & c) ^ (b & c)
}

// расширение с 16 слов до 64(эти 64 высчитываются по формуле из paint)
func ExpansionBlock(block []uint32) [64]uint32 {
	var W [64]uint32
	for i := 0; i < 16; i++ {
		W[i] = block[i]
	}
	for i := 16; i < 64; i++ {
		W[i] = C1(W[i-2]) + W[i-7] + C0(W[i-15]) + W[i-16]
	}
	return W
}

// производим сжатие по алгоритму из Paint
func Compress64(W [64]uint32, hashState [8]uint32) [8]uint32 {
	a, b, c, d, e, f, g, h := hashState[0], hashState[1], hashState[2], hashState[3],
		hashState[4], hashState[5], hashState[6], hashState[7]

	for i := 0; i < 64; i++ {
		S1_val := S1(e)
		ch_val := ch(e, f, g)
		temp1 := h + S1_val + ch_val + K[i] + W[i]

		S0_val := S0(a)
		maj_val := maj(a, b, c)
		temp2 := S0_val + maj_val

		h = g
		g = f
		f = e
		e = d + temp1
		d = c
		c = b
		b = a
		a = temp1 + temp2
	}

	return [8]uint32{a, b, c, d, e, f, g, h}
}

// конвертируем получившиеся 2-ичную систему в в 16 ричную(Как Вы сказали)
func convert_to_16(h [8]uint32) string {
	fmt.Println("Конвертация из двоичной в шестнадцатеричную:")
	fmt.Println()

	hex := fmt.Sprintf("%08x%08x%08x%08x%08x%08x%08x%08x",
		h[0], h[1], h[2], h[3], h[4], h[5], h[6], h[7])

	fmt.Println(hex)

	return hex
}
