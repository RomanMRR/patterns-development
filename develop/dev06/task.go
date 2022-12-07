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
	flag.StringVar(&f, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&d, "d", "	", "использовать другой разделитель")
	flag.BoolVar(&s, "s", false, "только строки с разделителями")

	flag.Parse()

	//Здесь храним номера колонок строки
	var colons []int
	if f == "" {
		println(" you must specify a list of bytes, characters, or fields")
		return
	}

	colons = MakeColons(f)
	for {
		var str string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			str = scanner.Text()
		}

		words := strings.Split(str, d)
		for _, col := range colons {
			if col < len(words) {
				if s && strings.Contains(words[col-1], d) {
					fmt.Println(words[col-1])
				} else {
					fmt.Println("Result:", words[col-1])
				}
			}
		}
	}
}

func MakeColons(flagF string) []int {
	fields := strings.Split(flagF, ",")

	var result []int
	for _, field := range fields {
		rangeStr := strings.Split(field, "-")
		var ranges [2]int
		for index, valueRange := range rangeStr {
			ranges[index], _ = strconv.Atoi(valueRange)
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
