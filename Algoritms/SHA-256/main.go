// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// type cryptographer struct {
// 	word string
// }

// func (c *cryptographer) sha256(word string) {
// 	first_step := convert_to_byte(word)
// 	second_step := plus_one(first_step)
// 	fird_step := add_zero(second_step)
// 	fourth_step := add_second_format(stringa, fird_step)
// 	fifth_Step := thirty_two_bit_as_string(fourth_step)
// 	sixth_Step := round_add_zeros_to_64(fifth_Step)
// 	seventh := parseWords32(sixth_Step)

// 	W := ExpansionBlock(seventh)
// 	eight := Compress64(W)
// 	ninth := joinWords32Bin(eight)

// 	convert_to_16(ninth)
// }

// const stringa = "hello world"

// var K = [64]uint32{
// 	0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5,
// 	0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
// 	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3,
// 	0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
// 	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc,
// 	0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
// 	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7,
// 	0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
// 	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13,
// 	0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
// 	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3,
// 	0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
// 	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5,
// 	0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
// 	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208,
// 	0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
// }

// // Начальные значения SHA‑256
// var H0 = [8]uint32{
// 	0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a,
// 	0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19,
// }

// func main() {
// 	one_String := cryptographer{word: "hello world"}
// 	one_String.sha256(one_String.word)

// 	twoo_String := cryptographer{word: "more symbols than can be contained in block"}
// 	one_String.sha256(twoo_String.word)

// 	// first_step := convert_to_byte(stringa)
// 	// second_step := plus_one(first_step)
// 	// fird_step := add_zero(second_step)
// 	// fourth_step := add_second_format(stringa, fird_step)
// 	// fifth_Step := thirty_two_bit_as_string(fourth_step)
// 	// sixth_Step := round_add_zeros_to_64(fifth_Step)
// 	// seventh := parseWords32(sixth_Step)

// 	// W := ExpansionBlock(seventh)
// 	// eight := Compress64(W)
// 	// ninth := joinWords32Bin(eight)

// 	// convert_to_16(ninth)

// }

// func convert_to_byte(word string) string {
// 	eight_byte_String := ""
// 	for _, b := range []byte(word) {
// 		temp := fmt.Sprintf("%08b", b)
// 		eight_byte_String += temp
// 	}

// 	return eight_byte_String
// }

// func plus_one(bit_string string) string {
// 	bit_string += "1"
// 	fmt.Println(bit_string)
// 	return bit_string
// }

// func add_zero(bit_String string) string {
// 	for range 359 {
// 		bit_String += "0"
// 	}

// 	fmt.Println(bit_String)
// 	return bit_String
// }

// func add_second_format(word, bite_String string) string {
// 	fmt.Println("")

// 	length_in_byte := len(word) * 8

// 	n := int64(length_in_byte)
// 	lenght_string := strconv.FormatInt(n, 2)

// 	quantity := 512 - len(bite_String) - len(lenght_string)

// 	for range quantity {
// 		bite_String += "0"
// 	}

// 	bite_String += lenght_string

// 	fmt.Println(bite_String)

// 	fmt.Println(len(bite_String))

// 	return bite_String

// }

// func thirty_two_bit_as_string(word string) string {
// 	out := ""
// 	for i := 0; i < len(word); i += 32 {

// 		end := i + 32
// 		if end > len(word) {
// 			end = len(word)
// 		}
// 		chunk := word[i:end]
// 		if i > 0 {
// 			out += " "
// 		}
// 		out += chunk
// 	}
// 	fmt.Println(out)
// 	return out
// }

// func round_add_zeros_to_64(word string) string {
// 	fmt.Println("")
// 	count := 0
// 	for i := 0; i < len(word); i++ {
// 		if word[i] == ' ' {
// 			count++
// 		}
// 	}

// 	for range count {
// 		word += " 00000000000000000000000000000000"
// 	}

// 	fmt.Println(word)

// 	return word
// }

// func parseWords32(line string) []uint32 {
// 	parts := strings.Fields(line)
// 	out := make([]uint32, 0, len(parts))
// 	for _, p := range parts {
// 		v, _ := strconv.ParseUint(p, 2, 32)
// 		out = append(out, uint32(v))
// 	}

// 	if len(out) > 16 {
// 		out = out[:16]
// 	}
// 	return out
// }

// func RightRotate32(n uint32, d uint32) uint32 {
// 	return (n >> d) | (n << (32 - d))
// }

// func C0(x uint32) uint32 {
// 	return RightRotate32(x, 7) ^ RightRotate32(x, 18) ^ (x >> 3)
// }

// func C1(x uint32) uint32 {
// 	return RightRotate32(x, 17) ^ RightRotate32(x, 19) ^ (x >> 10)
// }

// func S0(x uint32) uint32 {
// 	return RightRotate32(x, 2) ^ RightRotate32(x, 13) ^ RightRotate32(x, 22)
// }

// func S1(x uint32) uint32 {
// 	return RightRotate32(x, 6) ^ RightRotate32(x, 11) ^ RightRotate32(x, 25)
// }

// func ch(e, f, g uint32) uint32 {
// 	return (e & f) ^ ((^e) & g)
// }

// func maj(a, b, c uint32) uint32 {
// 	return (a & b) ^ (a & c) ^ (b & c)
// }

// func ExpansionBlock(block []uint32) [64]uint32 {
// 	var W [64]uint32
// 	for i := 0; i < 16; i++ {
// 		W[i] = block[i]
// 	}
// 	for i := 16; i < 64; i++ {
// 		W[i] = W[i-16] + S0(W[i-15]) + W[i-7] + S1(W[i-2])
// 	}
// 	return W
// }

// func Compress64(W [64]uint32) [8]uint32 {
// 	a, b, c, d, e, f, g, h := H0[0], H0[1], H0[2], H0[3], H0[4], H0[5], H0[6], H0[7]

// 	for i := 0; i < 64; i++ {
// 		S1 := S1(e)
// 		ch := ch(e, f, g)
// 		temp1 := h + S1 + ch + K[i] + W[i]

// 		S0 := S0(a)
// 		maj := maj(a, b, c)
// 		temp2 := S0 + maj

// 		h = g
// 		g = f
// 		f = e
// 		e = d + temp1
// 		d = c
// 		c = b
// 		b = a
// 		a = temp1 + temp2
// 	}

// 	return [8]uint32{
// 		H0[0] + a, H0[1] + b, H0[2] + c, H0[3] + d,
// 		H0[4] + e, H0[5] + f, H0[6] + g, H0[7] + h,
// 	}
// }

// func joinWords32Bin(h [8]uint32) string {
// 	parts := make([]string, 0, 8)
// 	for i := 0; i < 8; i++ {
// 		parts = append(parts, fmt.Sprintf("%032b", h[i]))
// 	}
// 	return strings.Join(parts, " ")
// }

// func convert_to_16(word string) string {
// 	binaryNumbers := strings.Fields(word)

// 	fmt.Println("Конвертация из двоичной в шестнадцатеричную:")
// 	fmt.Println()

// 	var hex string

// 	for _, binary := range binaryNumbers {

// 		decimal, err := strconv.ParseUint(binary, 2, 32)
// 		if err != nil {
// 			fmt.Printf("Ошибка при конвертации %s: %v\n", binary, err)
// 			continue
// 		}

// 		hex += fmt.Sprintf("%08x", decimal)

// 	}

// 	fmt.Println(hex)

// 	return hex
// }
