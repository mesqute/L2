package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов (кол-во символов)

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var n, r, u, m, b, c, h bool
	var k int

	// считываем ключи вида сортировки
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&m, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&h, "h", false, "сортировать по числовому значению с учётом суффиксов")
	flag.BoolVar(&c, "c", false, "проверять отсортированы ли данные")

	// считываем ключи параметров сортировки
	flag.IntVar(&k, "k", 0, "указание индекса колонки для сортировки")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&b, "b", false, "игнорировать хвостовые пробелы")

	flag.Parse()

	// считываем путь к сортируемому файлу
	input := flag.Arg(0)
	output := flag.Arg(1)

	if input == "" || output == "" {
		panic("Нужно указать файл для чтения и файл для записи")
	}

	if k < 1 {
		k = 0
	} else {
		k--
	}

	var data [][]string
	// получаем содержимое файла
	data = readFileToStringsLines(input)

	var sortFunc func(i, j int) bool
	// выбираем метод сортировки
	switch true {
	case n:
		// числовое значение
		sortFunc = func(i, j int) bool {
			a, _ := strconv.ParseFloat(getDataElem(data, i, k), 64)
			b, _ := strconv.ParseFloat(getDataElem(data, j, k), 64)
			if r {
				return a > b
			}
			return a < b
		}
	case m:
		// месяцы
		sortFunc = func(i, j int) bool {
			if r {
				return getMonth(getDataElem(data, j, k)).Before(getMonth(getDataElem(data, i, k)))
			}
			return getMonth(getDataElem(data, i, k)).Before(getMonth(getDataElem(data, j, k)))
		}
	case h:
		// кол-во символов в строке
		sortFunc = func(i, j int) bool {
			if r {
				return getLen(data[i][k:]) > getLen(data[j][k:])
			}
			return getLen(getDataElemSlice(data, i, k)) < getLen(getDataElemSlice(data, j, k))
		}
	default:
		// обычная сортировка (строки)
		sortFunc = func(i, j int) bool {
			if r {
				return getDataElem(data, i, k) > getDataElem(data, j, k)
			}
			return getDataElem(data, i, k) < getDataElem(data, j, k)
		}
	}

	// если введен ключ проверки сортировки, то проверяем файл на упорядоченность
	// согласно указанному правилу (ключу)
	if c {
		isSorted := sort.SliceIsSorted(data, sortFunc)
		fmt.Println("Sorted?", isSorted)
		return
	}

	// сортируем
	sort.Slice(data, sortFunc)

	// записываем отсортированные данные в файл
	writeToFile(data, output)
}

// readFileToStringsLines считывает данные из файла и преобразует в двумерный слайс (строки, колонки)
func readFileToStringsLines(dir string) (result [][]string) {
	// открываем файл для чтения
	file, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// создаем сканер
	scanner := bufio.NewScanner(file)
	// сканируем строки
	for scanner.Scan() {
		var words []string
		// разделяем строки на слова (колонки)
		words = strings.Split(scanner.Text(), " ")
		// записываем строки в возвращаемый слайс
		result = append(result, words)
	}

	return result
}

// writeToFile записывает данные в файл
func writeToFile(data [][]string, dir string) {
	// создаем (или перезаписываем) файл и открываем
	file, err := os.Create(dir)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// преобразуем данные в текст
	lines := make([]string, len(data))
	for i, datum := range data {
		str := strings.Join(datum, " ")
		lines[i] = str
	}
	// записываем текст в файл
	_, err = file.WriteString(strings.Join(lines, "\n"))
	if err != nil {
		panic(err)
	}
}

func getMonth(month string) time.Time {
	if t, err := time.Parse("Jan", month); err == nil {
		return t
	}
	if t, err := time.Parse("January", month); err == nil {
		return t
	}
	if t, err := time.Parse("1", month); err == nil {
		return t
	}
	if t, err := time.Parse("01", month); err == nil {
		return t
	}
	return time.Time{}
}

func getLen(str []string) int {
	var result int
	result = len(str) - 1
	for _, s := range str {
		result += len(s)
	}
	return result
}

// getDataElem возвращает элемент с индексом i,k если он есть,
// если его нет, возвращает пустую строку
func getDataElem(data [][]string, i, k int) string {
	if k < len(data[i]) {
		return data[i][k]
	}
	return ""
}

// getDataElemSlice возвращает элементы строки i начиная с индекса k если они есть,
// если их нет, возвращает пустой слайс
func getDataElemSlice(data [][]string, i, k int) []string {
	if k < len(data[i]) {
		return data[i][k:]
	}
	return []string{}
}
