// main.go — полноценный пример априорных метрик с построчными комментариями //

// Указываем исполняемый пакет                                                          //
package main // точка входа для программы                                              //

// Импорты стандартной библиотеки                                                       //
import ( // начало списка импортов                                                      //
	"fmt"     // форматированный вывод                                                  //
	"math"    // математические функции (для бесконечности)                             //
	"sort"    // сортировка срезов                                                      //
	"strings" // работа со строками                                                    //
) // конец списка импортов                                                              //

// Описание структуры транзакции (чека)                                                 //
type Transaction struct { // объявление структуры                      //
	ID    int             // уникальный идентификатор чека                           //
	Items map[string]bool // множество товаров в чеке (set через map[string]bool)    //
} // конец структуры                                                                    //

// Набор всех транзакций                                                                //
type Dataset []Transaction // тип-алиас для удобства                                    //

// Проверка, содержит ли чек все товары из заданного набора                             //
func containsAll(items map[string]bool, itemset []string) bool { // сигнатура функции   //
	for _, it := range itemset { // перебираем товары из набора                         //
		if !items[it] { // если товар не найден в чеке                                  //
			return false // набор не содержится полностью                               //
		} // конец if                                                                   //
	} // конец цикла                                                                    //
	return true // все товары из набора присутствуют                                     //
} // конец функции                                                                      //

// Подсчитать количество транзакций, содержащих заданный набор товаров                  //
func countWithItemset(dataset Dataset, itemset []string) int { // объявление функции    //
	cnt := 0                     // счётчик подходящих транзакций                                           //
	for _, tx := range dataset { // перебор всех транзакций                             //
		if containsAll(tx.Items, itemset) { // проверяем вхождение набора               //
			cnt++ // увеличиваем счётчик                                               //
		} // конец if                                                                   //
	} // конец цикла                                                                    //
	return cnt // возвращаем число транзакций с itemset                                  //
} // конец функции                                                                      //

// Расчёт поддержки: доля транзакций, содержащих набор                                  //
func support(dataset Dataset, itemset []string) float64 { // объявление функции         //
	if len(dataset) == 0 { // защита от деления на ноль                                 //
		return 0 // если данных нет — поддержка нулевая                                 //
	} // конец if                                                                       //
	return float64(countWithItemset(dataset, itemset)) / float64(len(dataset)) // формула//
} // конец функции                                                                      //

// Объединение двух наборов товаров (с устранением дублей и сортировкой)                //
func union(a, b []string) []string { // объявление функции                              //
	seen := make(map[string]bool)           // вспомогательное множество уже встреченных          //
	res := make([]string, 0, len(a)+len(b)) // результирующий срез                      //
	for _, x := range a {                   // добавляем элементы из a                                    //
		if !seen[x] { // если элемента ещё не было                                      //
			seen[x] = true       // помечаем как встреченный                                  //
			res = append(res, x) // добавляем в результат                                //
		} // конец if                                                                   //
	} // конец цикла                                                                    //
	for _, y := range b { // добавляем элементы из b                                    //
		if !seen[y] { // проверяем дубликаты                                            //
			seen[y] = true       // помечаем                                                  //
			res = append(res, y) // добавляем                                           //
		} // конец if                                                                   //
	} // конец цикла                                                                    //
	sort.Strings(res) // сортируем для стабильного отображения                           //
	return res        // возвращаем объединённый набор                                          //
} // конец функции                                                                      //

// Расчёт достоверности: Confidence(X→Y) = Support(X∪Y) / Support(X)                    //
func confidence(dataset Dataset, X, Y []string) float64 { // объявление функции         //
	supX := support(dataset, X) // поддержка антецедента X                              //
	if supX == 0 {              // защита от деления на ноль                                         //
		return 0 // если X не встречается — достоверность 0                             //
	} // конец if                                                                       //
	supXY := support(dataset, union(X, Y)) // поддержка объединения X∪Y                 //
	return supXY / supX                    // возвращаем значение confidence                               //
} // конец функции                                                                      //

// Расчёт лифта: Lift(X→Y) = Support(X∪Y) / (Support(X)*Support(Y))                     //
func lift(dataset Dataset, X, Y []string) float64 { // объявление функции               //
	supX := support(dataset, X) // поддержка X                                          //
	supY := support(dataset, Y) // поддержка Y                                          //
	if supX == 0 || supY == 0 { // защита от нулевого знаменателя                       //
		return 0 // если одна из поддержек нулевая — лифт 0                             //
	} // конец if                                                                       //
	supXY := support(dataset, union(X, Y)) // поддержка X∪Y                             //
	return supXY / (supX * supY)           // возвращаем значение лифта                           //
} // конец функции                                                                      //

// Расчёт conviction: Conv = (1 - Support(Y)) / (1 - Confidence(X→Y))                    //
func conviction(dataset Dataset, X, Y []string) float64 { // объявление функции         //
	conf := confidence(dataset, X, Y) // вычисляем достоверность                        //
	supY := support(dataset, Y)       // поддержка следствия Y                                //
	denom := 1 - conf                 // знаменатель формулы                                            //
	if denom == 0 {                   // если Confidence == 1                                             //
		return math.Inf(1) // по определению возвращаем +Inf                            //
	} // конец if                                                                       //
	return (1 - supY) / denom // возвращаем значение conviction                         //
} // конец функции                                                                      //

// Красивое форматирование набора товаров                                               //
func formatItemset(items []string) string { // объявление функции                       //
	cp := append([]string(nil), items...)     // делаем копию, чтобы не портить исходный    //
	sort.Strings(cp)                          // сортируем для предсказуемого порядка                             //
	return "{" + strings.Join(cp, ", ") + "}" // возвращаем строку вида {a, b, c}       //
} // конец функции                                                                      //

// Точка входа                                                                          //
func main() { // начало main                                                            //
	// Пример небольшого датасета транзакций (классика Market Basket Analysis)         //
	dataset := Dataset{ // инициализируем набор транзакций                              //
		{ID: 1, Items: map[string]bool{"Хлеб": true, "Молоко": true}},                                   // транзакция 1  //
		{ID: 2, Items: map[string]bool{"Хлеб": true, "Подгузники": true, "Пиво": true, "Яйца": true}},   // 2 //
		{ID: 3, Items: map[string]bool{"Молоко": true, "Подгузники": true, "Пиво": true, "Кола": true}}, // 3 //
		{ID: 4, Items: map[string]bool{"Хлеб": true, "Молоко": true, "Подгузники": true, "Пиво": true}}, // 4 //
		{ID: 5, Items: map[string]bool{"Хлеб": true, "Молоко": true, "Подгузники": true, "Кола": true}}, // 5 //
	} // конец инициализации                                                             //

	// Набор правил X -> Y, для которых посчитаем метрики                               //
	rules := []struct { // объявляем срез правил                                        //
		X []string // антецедент (левая часть)                                          //
		Y []string // консеквент (правая часть)                                         //
	}{ // начало литерала                                                                //
		{X: []string{"Молоко", "Подгузники"}, Y: []string{"Пиво"}}, // правило 1        //
		{X: []string{"Хлеб"}, Y: []string{"Молоко"}},               // правило 2                       //
		{X: []string{"Хлеб", "Молоко"}, Y: []string{"Подгузники"}}, // правило 3        //
		{X: []string{"Пиво"}, Y: []string{"Подгузники"}},           // правило 4                   //
	} // конец литерала                                                                  //

	// Заголовки и форматированная таблица                                               //
	fmt.Println("Оценка ассоциативных правил (Support, Confidence, Lift, Conviction)") // вывод заголовка  //
	fmt.Println(strings.Repeat("=", 78))                                               // разделитель                                   //

	header := fmt.Sprintf("%-35s | %-10s | %-10s | %-10s | %-10s", "Правило X -> Y", "Support", "Conf.", "Lift", "Conv.") // шапка таблицы //
	fmt.Println(header)                                                                                                   // печатаем шапку                                                //
	fmt.Println(strings.Repeat("-", len(header)))                                                                         // разделительная линия                //

	// Основной цикл расчётов по каждому правилу                                         //
	for _, r := range rules { // перебираем правила                                      //
		Xs := formatItemset(r.X)                                                                                 // строка для X                                         //
		Ys := formatItemset(r.Y)                                                                                 // строка для Y                                         //
		supXY := support(dataset, union(r.X, r.Y))                                                               // поддержка объединённого набора     //
		conf := confidence(dataset, r.X, r.Y)                                                                    // достоверность правила                   //
		lf := lift(dataset, r.X, r.Y)                                                                            // лифт правила                                    //
		conv := conviction(dataset, r.X, r.Y)                                                                    // conviction правила                      //
		row := fmt.Sprintf("%-35s | %-10.4f | %-10.4f | %-10.4f | %-10.4f", Xs+" -> "+Ys, supXY, conf, lf, conv) // строка //
		fmt.Println(row)                                                                                         // печать строки                                                //
	} // конец цикла                                                                     //

	fmt.Println(strings.Repeat("=", 78)) // финальный разделитель                        //

	// Для наглядности выведем поддержку отдельных наборов                               //
	samples := [][]string{ // список наборов для демонстрации поддержки                  //
		{"Хлеб"},                 // набор 1                                              //
		{"Молоко"},               // набор 2                                              //
		{"Подгузники"},           // набор 3                                              //
		{"Пиво"},                 // набор 4                                              //
		{"Молоко", "Подгузники"}, // набор 5                                             //
	} // конец списка                                                                     //

	fmt.Println("Поддержка некоторых наборов:") // подзаголовок                           //
	for _, s := range samples {                 // цикл по наборам                                       //
		fmt.Printf("supp(%s) = %.4f\n", formatItemset(s), support(dataset, s)) // печать //
	} // конец цикла                                                                      //
} // конец main                                                                          //
