package main

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

import (
	"bufio"
	"flag"
	"fmt"

	// "fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/exp/slices"
)

// Объявляем ключи
var (
	after      = flag.Int("A", 0, "after")
	before     = flag.Int("B", 0, "before")
	context    = flag.Int("C", 0, "context")
	count      = flag.Bool("c", false, "count")
	ignoreCase = flag.Bool("i", false, "ignoreCase")
	invert     = flag.Bool("v", false, "invert")
	fixed      = flag.Bool("F", false, "fixed")
	lineNum    = flag.Bool("n", false, "lineNum")
)

var root, query string //где ищем, что ищем

var wg sync.WaitGroup

// Читаем файл и ищем
func readFile(wg *sync.WaitGroup, path string) {
	var pattern *regexp.Regexp
	defer wg.Done()
	fileStrings := make([]string, 0) //Здесь храним все строки файла
	fileIndexes := make([]int, 0)    //Здесь храним номера нужных строка

	if *ignoreCase {
		pattern, _ = regexp.Compile("(?i)" + query) //Регулярное выражение для игнорирования регистра
	} else {
		pattern, _ = regexp.Compile(query)
	}

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return
	}

	//Читаем весь файл и сохраняем прочитанное
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		fileStrings = append(fileStrings, scanner.Text())

	}

	for i := 0; i < len(fileStrings); i++ {
		if *fixed { //Ищем только то, что написано искать
			if strings.Contains(fileStrings[i], query) {
				fileIndexes = append(fileIndexes, i)
			}
		} else if pattern.MatchString(fileStrings[i]) { //Ищем по регулярному выражению
			fileIndexes = append(fileIndexes, i)
		}
	}

	switch {
	case *count:
		println(len(fileIndexes)) //Выводим только количество найденных строк
	case *invert: //Выводим только те строки, которые нам не подходят
		for i, v := range fileStrings {
			if !slices.Contains(fileIndexes, i) {
				if *lineNum { //Выводим с номерами
					println(fmt.Sprintf("%d:%s", i+1, v))
				} else { //Без номеров
					println(v)
				}
			}
		}
	default:
		for i, string := range fileStrings {
			for _, indexMatch := range fileIndexes {
				flag := printBeforeOrAfter(i, indexMatch)

				if *lineNum && flag {
					println(fmt.Sprintf("%d:%s", i+1, string))
				} else if flag {
					println(fmt.Sprintf(string))
				}
			}
		}
	}
}

// Проверяет, нужно ли печатать до или после нужной строки ещё строки
func printBeforeOrAfter(index int, indexMatch int) bool {
	//index - номер строки, которая не подходит
	//indexMatch - подходящая строка
	if *after > 0 {
		return index-indexMatch <= *after && index-indexMatch >= 0
	} else if *before > 0 {
		return indexMatch-index <= *before && indexMatch-index >= 0
	} else if *context > 0 {
		return index-indexMatch <= *context && index-indexMatch >= 0 ||
			indexMatch-index <= *context && indexMatch-index >= 0
	} else if index == indexMatch {
		return true
	}

	return false
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		println("Usage: program [OPTION[... PATTERS[FILE]...")
		return
	}

	query = flag.Arg(0)
	root = flag.Arg(1)

	filepath.Walk(root, func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			wg.Add(1)
			go readFile(&wg, path)
		}
		return nil
	})
	wg.Wait()
}
