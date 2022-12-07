package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	k = flag.Int("k", 0, "указание колонки для сортировки")
	n = flag.Bool("n", false, "")
	r = flag.Bool("r", false, "сортировать в обратном порядке")
	u = flag.Bool("u", false, "не выводить повторяющиеся строки")
	M = flag.Bool("M", false, "сортировать по названию месяца")
	b = flag.Bool("b", false, "игнорировать хвостовые пробелы")
	c = flag.Bool("c", false, "проверять отсортированы ли данные")
	h = flag.Bool("h", false, "сортировать по числовому значению с учетом суффиксов")
)

// Функция для чтения файла
func readFile(fileName string) (result [][]string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var words []string
		words = strings.Split(scanner.Text(), " ")
		result = append(result, words)
	}
	file.Close()
	return result
}

// Функция для сортировки данных
func Sort(fileName string) [][]string {
	var data [][]string
	var sortFunc func(i, j int) bool //Функция по которой будут сортироваться элементы
	if *k < 0 {
		println(*k)
		panic(fmt.Sprintf("Invalid number at %v", *k))
	} else {
		if *k != 0 {
			*k--
		}
	}
	data = readFile(fileName) //Получаем файл
	switch {
	case *n:
		sortFunc = func(i, j int) bool {
			//Получаем слова из файла
			a, _ := strconv.ParseFloat(getDataElem(data, i, *k), 64) //Преобразуем в числа
			b, _ := strconv.ParseFloat(getDataElem(data, j, *k), 64)
			//И сравниваем их
			if *r {
				return a > b
			}
			return a < b
		}
	default:
		sortFunc = func(i, j int) bool {
			if *r {
				return getDataElem(data, i, *k) > getDataElem(data, j, *k)
			}
			return getDataElem(data, i, *k) < getDataElem(data, j, *k)
		}
	}
	//Смотрим, отсортирован ли массив
	if *c {
		println("Данные отсортированы:", sort.SliceIsSorted(data, sortFunc))
		os.Exit(0)
	}

	sort.Slice(data, sortFunc) //Сортируем данные по созданной функции

	return data
}

func getDataElem(data [][]string, i, k int) string {
	if k < len(data[i]) {
		return data[i][k]
	}
	return ""
}

// Вывод результата
func printResult(data [][]string) {
	RepeatedStrings := make(map[string]bool) //Для сохранения повторяющихся строк
	for _, row := range data {
		value := strings.Join(row, " ")

		if *u {
			if !RepeatedStrings[value] {
				println(value)
			}
			RepeatedStrings[value] = true
		} else {
			println(value)
		}

	}
}

func main() {
	flag.Parse()
	var file string

	file = flag.Arg(0) //Считываем название файла
	result := Sort(file)
	printResult(result)
}
