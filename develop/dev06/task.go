package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

	var f, d string
	var s bool

	// считываем ключи
	flag.StringVar(&f, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&d, "d", " ", "использовать другой разделитель")
	flag.BoolVar(&s, "s", false, "только строки с разделителем")

	flag.Parse()

	// инициализируем переменную, сохраняющую индексы нужных колонок
	var cols []int
	// если флаг f активирован
	if f != "" {
		// считываем номера нужных колонок
		cols = convertStrToIntArr(f)
	}

	// бесконечный цикл обработки входящих строк
	for {
		fmt.Print("Введите строку: ")
		// инициализируем переменную, хранящую полученную из STDIN строку
		var str string
		// инициализируем сканер STDIN
		scanner := bufio.NewScanner(os.Stdin)
		// ждем пока в STDIN не будет написана строка
		if scanner.Scan() {
			// считываем строку
			str = scanner.Text()
		}

		// делим полученную строку
		words := strings.Split(str, d)

		// если есть определенные нужные колонки (активирован ключ f), то выводим их содержимое
		if len(cols) > 0 {
			for _, col := range cols {
				if col < len(words) {
					fmt.Println(words[col])
				}
			}
		}
		// если ключ f не активен, выводим все полученные отрезки
		for _, word := range words {
			fmt.Println(word)
		}
	}
}

// convertStrToIntArr конвертирует текстовое представление числовых отрезков в численное
// ("1-3 5 7-8" -> 1,2,3,5,7,8)
func convertStrToIntArr(str string) []int {
	// делим полученную строку на области
	fields := strings.Split(str, " ")

	// инициализируем переменную для вывода
	var result []int
	// обходим каждую область
	for _, field := range fields {
		// считываем промежутки
		rangesStr := strings.Split(field, "-")
		// инициализируем переменную хранящую промежутки
		var ranges [2]int
		// считываем промежутки
		for i, s := range rangesStr {
			ranges[i], _ = strconv.Atoi(s)
		}
		// если записана только одна граница промежутка (область - одно число),
		// то копируем границу
		if ranges[1] == 0 {
			ranges[1] = ranges[0]
		}
		// записываем в возвращаемую переменную числа входящие в промежуток
		for i := ranges[0]; i <= ranges[1]; i++ {
			result = append(result, i)
		}
	}
	return result
}
