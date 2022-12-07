package main

import (
	"log"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unpack(s []rune) string {
	var lastSymbol, lastLetter rune
	var result, num strings.Builder
	var escape bool
	lastSymbol = 0 //Здесь будем сохранять и буквы и цифры
	lastLetter = 0 //Здесь будем сохранять только буквы
	//Если в начале число, то сразу выходим
	if len(s) == 0 || unicode.IsDigit(s[0]) {
		return ""
	}
	for i := 0; i < len(s); i++ {
		curSymbol := s[i]
		if unicode.IsLetter(curSymbol) {
			// Если предыдущий символ был числом
			if unicode.IsDigit(lastSymbol) {
				numRunes, err := strconv.Atoi(num.String())
				if err != nil {
					log.Fatal(err)
				}
				for j := 0; j < numRunes-1; j++ {
					result.WriteRune(lastLetter) //Записываем в результат букву перед этим числом
				}
				num.Reset()
			}
			//Букву записываем в результат
			result.WriteRune(curSymbol)
			//И сохраняем символ
			lastLetter = curSymbol
			lastSymbol = curSymbol
		}
		if unicode.IsDigit(curSymbol) {
			// Пропускаем цифры
			if escape { //если было экранирование
				result.WriteRune(curSymbol) //То просто записываем цифры в результат
				lastLetter = curSymbol
				lastSymbol = curSymbol
				escape = false
			} else {
				num.WriteRune(curSymbol) //Добавляем в числовую строку ещё одну цифру
				lastSymbol = curSymbol   //Сохраняем последнюю цифру
				// Обработка последней цифры в строке
				if i == utf8.RuneCountInString(string(s))-1 {
					numRunes, err := strconv.Atoi(num.String())
					if err != nil {
						log.Fatal(err)
					}
					for j := 0; j < numRunes-1; j++ {
						result.WriteRune(lastLetter)
					}
				}
			}

		}
		if curSymbol == '\\' {
			if lastSymbol == '\\' { //Если предыдущий символ был "/", то он работает как экранирование, а значит текущий символ нужно сохранить
				result.WriteRune(curSymbol)
				lastLetter = curSymbol
				lastSymbol = curSymbol
				escape = false //Символ "/" уже произвёл экранирование, а значит пропускать цифры сы не будем

			} else {
				escape = true //Возможно придётся пропустить цифры как есть
				lastSymbol = curSymbol
			}
		}
	}

	return result.String()
}

func main() {
	packedString := []rune(`qwe\4\5`)
	println(Unpack(packedString))
}
