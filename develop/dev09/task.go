package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	bufSize = 1024 * 8
)

// Функция записывает результаты запроса Get в файл.
// В качестве имени файла используется последний фрагмент URL-адреса. Например: http://foo/baz.jar => baz.jar
func Wget(url string) {
	resp := getResponse(url)
	urlSplit := strings.Split(url, "/")
	fileName := urlSplit[len(urlSplit)-1]
	writeToFile(fileName, resp)
}

// Функуция делает GET-запрос и возвращает ответ
func getResponse(url string) *http.Response {
	tr := new(http.Transport)
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	errorChecker(err)
	return resp
}

// Пишем ответ GET-запроса в файл
func writeToFile(fileName string, resp *http.Response) {
	// Credit for this implementation should go to github user billnapier
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	errorChecker(err)
	defer file.Close()
	bufferedWriter := bufio.NewWriterSize(file, bufSize)
	errorChecker(err)
	_, err = io.Copy(bufferedWriter, resp.Body)
	errorChecker(err)
}

// Проверка наличия ошибки после последнего вызова функции
func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin) //Создаём объект для чтения данных
	for {
		fmt.Print("$ ")                           //Приглашение на ввод
		cmdString, err := reader.ReadString('\n') //Читаем введённую команду
		if err != nil {
			fmt.Fprintln(os.Stderr, err) //Если пполучили ошибку, то выводим ошибку в поток stderr
		}
		commands := strings.Fields(cmdString)        //Разделям строку на слова и заносим из в массив строк
		commands = append(commands, commands[1:]...) //Убираем из массива комманду wget
		Wget(commands[1])
	}
}
