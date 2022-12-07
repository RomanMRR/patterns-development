package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

	ПЛЮСЫ
	 Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	Позволяет реализовать простую отмену и повтор операций.
	Позволяет реализовать отложенный запуск операций.
	Позволяет собирать сложные команды из простых.
	Реализует принцип открытости/закрытости.

	МИНУСЫ
	Усложняет код программы из-за введения множества дополнительных классов.

	 Применимость
	 Когда вы хотите параметризовать объекты выполняемым действием.
	 Когда вы хотите ставить операции в очередь, выполнять их по расписанию или передавать по сети.
	  Когда вам нужна операция отмены.

	Пример на практике
	К примеру, программа, текстового редактора. Вместо создания множество классов кнопок, которые выполняют различные команды, можно создать одну кнопку,
	и несколько классов команд, которые вызывают нужно исполнителя.Также и контекстное меню или горячие клавиши, которые также будут обращаться к тем же
	командам, что избавит код от дублирования.


*/

// Реализуем того, что вызывает команду
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

// Интерфейс для всех команд
type Command interface {
	execute()
}

// Реализация самой команды включения
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

// Реализация команды выключения
type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// Интерфейс того, кто будет исполнять команды
type Device interface {
	on()
	off()
}

// Тот кто будет исполнять команды
type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
