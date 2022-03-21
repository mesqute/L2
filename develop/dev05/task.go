package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

	var a, b, c int
	var count, ignore, invert, fixed, lnum bool

	// считываем ключи
	flag.IntVar(&a, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&b, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&c, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&count, "c", false, "количество строк")
	flag.BoolVar(&lnum, "n", false, "печатать номер строки")

	flag.BoolVar(&fixed, "F", false, "точное совпадение со строкой")
	flag.BoolVar(&invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&ignore, "i", false, "игнорировать регистр")

	flag.Parse()

	// считываем искомую строку
	str := flag.Arg(0)
	if len(str) == 0 {
		panic("некорректный ввод искомой строки!")
	}
	// считываем файл для чтения
	input := flag.Arg(1)
	if len(str) == 0 {
		panic("Нужно указать файл для чтения")
	}

	// считываем данные из файла
	data := readFileToStrings(input)

	// функции для обработки ключей F, v, i

	// возвращает неизмененные значения
	prepareStrings := func(str, substr string) (string, string) {
		return str, substr
	}
	// проверяет, содержит ли строка подстроку
	checkContain := func(str, substr string) bool {
		return strings.Contains(str, substr)
	}
	// возвращает неизмененное значение
	handleCheck := func(check bool) bool {
		return check
	}

	// если активен ключ i
	if ignore {
		// приводит обе строки к нижнему регистру
		prepareStrings = func(str, substr string) (string, string) {
			return strings.ToLower(str), strings.ToLower(substr)
		}
	}
	// если активен ключ F
	if fixed {
		// проводит прямое сравнение строк
		checkContain = func(str, substr string) bool {
			if str == substr {
				return true
			}
			return false
		}
	}
	// если активен ключ v
	if invert {
		// инвертирует значение
		handleCheck = func(check bool) bool {
			return !check
		}
	}

	// инициализируем слайс, который будет хранить значения индексов строк
	// удовлетворяющих условиям
	var conIdx []int

	// обходим все строки
	for i, line := range data {
		// если строка удовлетворяет всем условиям добавляем ее индекс в слайс
		if handleCheck(checkContain(prepareStrings(line, str))) {
			conIdx = append(conIdx, i)
		}
	}

	switch true {
	// обработка ключа A
	case a > 0:
		// обходим все сохраненные индексы
		for _, idx := range conIdx {
			// выводим в STDOUT A строк стоящих перед строкой с сохраненным индексом
			for i := 0; i < a; i++ {
				if idx-a+i >= 0 {
					fmt.Println(data[idx-a+i])
				}
			}
		}
	// обработка ключа B
	case b > 0:
		// обходим все сохраненные индексы
		for _, idx := range conIdx {
			// выводим в STDOUT B строк стоящих после строки с сохраненным индексом
			for i := 0; i < b; i++ {
				if idx+i < len(data) {
					fmt.Println(data[idx+i])
				}
			}
		}
	// обработка ключа C
	case c > 0:
		// обходим все сохраненные индексы
		for _, idx := range conIdx {
			// выводим в STDOUT C строк стоящих перед строкой и после строки с сохраненным индексом
			for i := 0; i < c; i++ {
				if idx-c+i >= 0 {
					fmt.Println(data[idx-c+i])
				}
				if idx+i < len(data) {
					fmt.Println(data[idx+i])
				}
			}
		}
	// обработка ключа c
	case count:
		// выводим кол-во сохраненных индексов (кол-во строк удовлетворяющих условиям)
		fmt.Println(len(conIdx))
	// обработка ключа n
	case lnum:
		// выводим все сохраненные индексы + 1 (номера строк)
		for _, idx := range conIdx {
			fmt.Println(idx + 1)
		}
	default:
		// выводим все строки удовлетворяющие условиям
		for _, idx := range conIdx {
			fmt.Println(data[idx])
		}
	}

}

// readFileToStringsLines считывает данные из файла и преобразует в слайс
func readFileToStrings(dir string) (result []string) {
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
		// записываем строки в возвращаемый слайс
		result = append(result, scanner.Text())
	}

	return result
}
