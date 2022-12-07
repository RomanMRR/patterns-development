package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func SortWord(word string) string {
	wordRune := []rune(word)
	sort.Slice(wordRune, func(i, j int) bool { return wordRune[i] < wordRune[j] })
	return string(wordRune)
}

func GetMapAnagram(words []string) map[string][]string {
	result := make(map[string][]string)
	uniqWord := make(map[string]struct{})
	outputResult := make(map[string][]string)

	for _, word := range words {
		word = strings.ToLower(word)
		anagram := SortWord(word)
		if _, ok := uniqWord[word]; !ok {
			result[anagram] = append(result[anagram], word)
			uniqWord[word] = struct{}{}
		}
	}

	for _, value := range result {
		if len(value) > 1 {
			outputResult[value[0]] = value
		}
	}

	return outputResult
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	result := GetMapAnagram(words)

	for key, value := range result {
		print(key, ": ")
		for _, word := range value {
			print(word, ", ")
		}
		println()
	}
	// fmt.Println(result)
}
