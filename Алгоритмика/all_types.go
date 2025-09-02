package main

import (
	"fmt"
	"math"
	"time"
)

// Суть обработать все типы данных и структуры, которые есть в Go

const text string = "Выбери, что хочешь изучить:\n1.Nill-переменные\n2.Числовые-переменные\n3.Перемнные - истинности\n4.Константы\n5.Строковые-переменные\n6.Массивы и Слайсы\n7.Структуры\n8.Выход из программы"

var choice int

func main() {
	fmt.Println(text)

	fmt.Println("Введите номер:")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		nillValuables()
	case 2:
		digits()
	case 3:
		boolean()
	case 4:
		constants()
	case 5:
		strings()
	case 6:
		arrays_slices()
	case 7:
		strcuture()
	case 8:
		fmt.Println("Выход из программы...")
		return
	default:
		fmt.Println("У нас нет такой команды")
		time.Sleep(time.Second * 2)
		main()
	}

}

func nillValuables() {
	//Когда мы инициализируем переменную без значения - она будет иметь значение по умолчанию для типа
	var intNill int //создали 0 переменную а
	fmt.Println("intNill =", intNill)

	var strNill string
	var floatNill float64
	var bollNill bool

	fmt.Printf("strNill: Value %#v\n, пустая строка и тд", strNill)
	fmt.Printf("floatNill: Value %#v\n, пустая строка и тд", floatNill)
	fmt.Printf("boolNill: Value %#v\n, пустая строка и тд", bollNill)

	intNill = 5 //Присвоим 5
	fmt.Println("intNill =", intNill)

	time.Sleep(time.Second * 2)
	main()
}

func digits() {
	fmt.Println("Создём все виды числовых значений")
	fmt.Println("")
	fmt.Println("Создаём все типы int (они различаются тем насколько большое число может хранить)")

	var valuableInt int
	var valuableInt8 int8
	var valuableInt32 int32
	var valuableInt64 int64

	// Присваиваем максимальные значения для signed int типов
	valuableInt = math.MaxInt
	valuableInt8 = math.MaxInt8   // 127
	valuableInt32 = math.MaxInt32 // 2147483647
	valuableInt64 = math.MaxInt64 // 9223372036854775807

	fmt.Printf("valuableInt: Value %d, Int\n", valuableInt)
	fmt.Printf("valuableInt8: Value %d, Int8\n", valuableInt8)
	fmt.Printf("valuableInt32: Value %d, Int32\n", valuableInt32)
	fmt.Printf("valuableInt64: Value %d, Int64\n", valuableInt64)

	fmt.Println("")
	fmt.Println("Создаём все типы uint (беззнаковые целые числа)")

	var valuableUint uint
	var valuableUint8 uint8
	var valuableUint32 uint32
	var valuableUint64 uint64

	// Присваиваем максимальные значения для unsigned int типов
	valuableUint = math.MaxUint     // зависит от архитектуры
	valuableUint8 = math.MaxUint8   // 255
	valuableUint32 = math.MaxUint32 // 4294967295
	valuableUint64 = math.MaxUint64 // 18446744073709551615

	fmt.Printf("valuableUint: Value %d, Uint\n", valuableUint)
	fmt.Printf("valuableUint8: Value %d, Uint8\n", valuableUint8)
	fmt.Printf("valuableUint32: Value %d, Uint32\n", valuableUint32)
	fmt.Printf("valuableUint64: Value %d, Uint64\n", valuableUint64)

	fmt.Println("")
	fmt.Println("Создаём типы с плавающей точкой")

	var valuableFloat32 float32
	var valuableFloat64 float64

	// Присваиваем максимальные значения для float типов
	valuableFloat32 = math.MaxFloat32 // примерно 3.4028235e+38
	valuableFloat64 = math.MaxFloat64 // примерно 1.7976931348623157e+308

	fmt.Printf("valuableFloat32: Value %v, float32\n", valuableFloat32)
	fmt.Printf("valuableFloat64: Value %v, float64\n", valuableFloat64)

	time.Sleep(time.Second * 2)
	main()
}

func boolean() {
	fmt.Println("Bool значения")

	True := true
	False := false
	fmt.Println("True:", True, "False:", False)

	True = False
	False = true
	fmt.Println("True:", True, "False:", False)

	time.Sleep(time.Second * 2)
	main()
}

func constants() {
	fmt.Println("Константы")
	fmt.Println("Константы - это пеемнные, которые мы создали и больше не меняем. Это нужно для чистоты и читабельности кода")

	const s string = "Создание константы"

	fmt.Printf("Константа: %#v", s)
	// s = "Изменяем" //Это не сработает

	time.Sleep(time.Second * 2)
	main()
}

func strings() {
	fmt.Println("")
	fmt.Println("Работа со строками")
	// Создание строк разными способами
	var emptyString string
	var simpleString string = "Привет, мир!"
	shortString := "Короткое объявление"

	fmt.Printf("emptyString: %#v (длина: %d)\n", emptyString, len(emptyString))
	fmt.Printf("simpleString: %#v (длина: %d)\n", simpleString, len(simpleString))
	fmt.Printf("shortString: %#v (длина: %d)\n", shortString, len(shortString))

	// Многострочные строки
	multilineString := `Это многострочная строка.
Она может содержать несколько строк.
И символы "кавычки" без экранирования.`

	fmt.Printf("multilineString:\n%s\n", multilineString)

	// Экранирование символов
	escapedString := "Строка с \"кавычками\" и \nпереносом строки и \tтабуляцией"
	fmt.Printf("escapedString: %s\n", escapedString)

	// Конкатенация строк
	firstName := "Иван"
	lastName := "Петров"
	fullName := firstName + " " + lastName
	fmt.Printf("Полное имя: %s\n", fullName)

	// Работа с Unicode и рунами
	unicodeString := "Привет 🚀 мир!"
	fmt.Printf("Unicode строка: %s\n", unicodeString)
	fmt.Printf("Длина в байтах: %d\n", len(unicodeString))
	fmt.Printf("Длина в рунах: %d\n", len([]rune(unicodeString)))

	// Перебор строки как байтов
	fmt.Print("Байты: ")
	for i := 0; i < len(unicodeString); i++ {
		fmt.Printf("%d ", unicodeString[i])
	}
	fmt.Println()

	// Перебор строки как рун
	fmt.Print("Руны: ")
	for _, r := range unicodeString {
		fmt.Printf("%c ", r)
	}
	fmt.Println()

	// Конвертация типов
	numberStr := "42"
	fmt.Printf("Строка числа: %#v\n", numberStr)

	// Строки неизменяемы в Go
	originalString := "Неизменная строка"
	fmt.Printf("Оригинальная строка: %s\n", originalString)

	// Создание новой строки на основе существующей
	newString := originalString + " была изменена"
	fmt.Printf("Новая строка: %s\n", newString)

	// Срезы строк (substring)
	testString := "Тестовая строка для среза"
	fmt.Printf("Исходная строка: %s\n", testString)
	fmt.Printf("Срез [0:8]: %s\n", testString[0:8])
	fmt.Printf("Срез [9:]: %s\n", testString[9:])

	// Сравнение строк
	str1 := "apple"
	str2 := "banana"
	str3 := "apple"

	fmt.Printf("'%s' == '%s': %t\n", str1, str2, str1 == str2)
	fmt.Printf("'%s' == '%s': %t\n", str1, str3, str1 == str3)
	fmt.Printf("'%s' < '%s': %t (лексикографическое сравнение)\n", str1, str2, str1 < str2)

	// Raw strings с обратными кавычками
	rawString := `C:\Users\Name\Documents\file.txt`
	fmt.Printf("Raw string (путь): %s\n", rawString)

	// Пустая строка vs nil
	var nilString string
	emptyStr := ""
	fmt.Printf("nil строка: %#v (пустая: %t)\n", nilString, nilString == "")
	fmt.Printf("пустая строка: %#v (пустая: %t)\n", emptyStr, emptyStr == "")
	fmt.Printf("nil == empty: %t\n", nilString == emptyStr)

	fmt.Println("=== Конец работы со строками ===")

	time.Sleep(time.Second * 2)
	main()
}

func arrays_slices() {
	fmt.Println("")
	fmt.Println("=== Массивы и Слайсы ===")
	fmt.Println("Массив - фиксированный размер")
	fmt.Println("Слайс - может изменяться")
	fmt.Println("")

	// ===== МАССИВЫ =====
	fmt.Println("--- МАССИВЫ ---")

	// Создание массива (размер фиксированный!)
	var numbers [3]int // [0 0 0]
	fmt.Printf("Пустой массив: %v\n", numbers)

	// Заполняем массив
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	fmt.Printf("Заполненный массив: %v\n", numbers)

	// Создание массива сразу со значениями
	fruits := [3]string{"яблоко", "банан", "вишня"}
	fmt.Printf("Массив фруктов: %v\n", fruits)
	fmt.Printf("Длина массива: %d\n", len(fruits))

	fmt.Println("")

	// ===== СЛАЙСЫ =====
	fmt.Println("--- СЛАЙСЫ ---")

	// Создание пустого слайса
	var colors []string
	fmt.Printf("Пустой слайс: %v (nil: %t)\n", colors, colors == nil)

	// Добавляем элементы в слайс
	colors = append(colors, "красный")
	colors = append(colors, "синий", "зеленый")
	fmt.Printf("Слайс с цветами: %v\n", colors)
	fmt.Printf("Длина: %d, Вместимость: %d\n", len(colors), cap(colors))

	// Создание слайса сразу со значениями
	animals := []string{"кот", "собака", "птица"}
	fmt.Printf("Слайс животных: %v\n", animals)

	// Создание слайса с make
	grades := make([]int, 3) // создает слайс длиной 3
	grades[0] = 85
	grades[1] = 92
	grades[2] = 78
	fmt.Printf("Оценки: %v\n", grades)

	fmt.Println("")

	// ===== ОСНОВНЫЕ ОПЕРАЦИИ =====
	fmt.Println("--- ОСНОВНЫЕ ОПЕРАЦИИ ---")

	// Получение части слайса
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Printf("Все буквы: %v\n", letters)
	fmt.Printf("Первые 3: %v\n", letters[:3])       // [A B C]
	fmt.Printf("С 2-й до конца: %v\n", letters[2:]) // [C D E]
	fmt.Printf("Средние: %v\n", letters[1:4])       // [B C D]

	// Изменение элементов
	letters[0] = "Z"
	fmt.Printf("После изменения: %v\n", letters)

	// Перебор элементов
	fmt.Print("Перебор слайса: ")
	for i, letter := range letters {
		fmt.Printf("%d:%s ", i, letter)
	}
	fmt.Println()

	fmt.Println("")

	// ===== ГЛАВНЫЕ РАЗЛИЧИЯ =====
	fmt.Println("--- РАЗЛИЧИЯ ---")

	// Массив: размер фиксированный
	array := [2]int{1, 2}
	fmt.Printf("Массив: %v (размер нельзя изменить)\n", array)

	// Слайс: размер можно изменять
	slice := []int{1, 2}
	slice = append(slice, 3, 4, 5)
	fmt.Printf("Слайс: %v (размер изменился!)\n", slice)

	// Длина и вместимость
	fmt.Printf("У слайса len=%d, cap=%d\n", len(slice), cap(slice))
	fmt.Printf("У массива только len=%d\n", len(array))

	fmt.Println("=== Конец ===")

	time.Sleep(time.Second * 2)
	main()
}

type Person struct {
	Name string
	Age  int
}

func strcuture() {
	fmt.Println("Структуры")
	fmt.Println("Структуры - это набор полей, которые мы объединяем в одну переменную")
	fmt.Println("Структуры это упорядоченные коллекции данных, которые могут хранить различные типы данных")

	// type Person struct {
	// 	Name string
	// 	Age  int
	// }

	Ilya := Person{Name: "Илья", Age: 18}

	fmt.Println(Ilya)
	fmt.Printf("Можем обращаться отдельно Name: %s, Age: %d\n", Ilya.Name, Ilya.Age)

	Ilya.PrintName() //Метод который пользуется структурой(Только эта структура может этим воспользоватьсяы)
	time.Sleep(time.Second * 2)
	main()
}

func (p Person) PrintName() {
	fmt.Println(p.Name)
}
