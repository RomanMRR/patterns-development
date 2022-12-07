package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin) //Создаём объект для чтения данных
	for {
		fmt.Print("$ ")                           //Приглашение на ввод
		cmdString, err := reader.ReadString('\n') //Читаем введённую команду
		if err != nil {
			fmt.Fprintln(os.Stderr, err) //Если пполучили ошибку, то выводим ошибку в поток stderr
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// Для обработки введённых команд
func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n") //Так как пользователь нажимает Enter в конце, то удаляем этот символ
	arrCommandStr := strings.Fields(commandStr)       //Разделям строку на слова и заносим из в массив строк
	switch arrCommandStr[0] {                         //Выбираем нужную команду
	case `quit`: //Просто выходим
		os.Exit(0)
	case "cd":
		if len(arrCommandStr) < 2 { //Если ввели только команду без аргументов, то ничего не делам
			return nil
		} else if len(arrCommandStr) > 2 { //Если аргументов больше двух, то это ошибка
			return errors.New("-shell: cd: too many arguments")
		}
		os.Chdir(arrCommandStr[1]) //Меняем директорию
		return nil
	case "fork":
		syscall.Syscall(syscall.SYS_FORK, 0, 0, 0) //Применям системный вызов fork(), чтобы создать дочерний процесс
		return nil
	case "exec":
		cmd := exec.Command(arrCommandStr[1], arrCommandStr[2:]...) //Создаю команду, получунную из аргументов
		//Делаем так, чтобы команда выводила результаты и ошибки в стандартные потоки выводу и ошибок
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return cmd.Run() //Запускаем команду

	}
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...) //Если нет какой-то команду выше, то пробуем запустить её стандартной функцией
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
