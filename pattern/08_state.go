package main

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

	ПЛЮСЫ
	 Избавляет от множества больших условных операторов машины состояний.
	 Концентрирует в одном месте код, связанный с определённым состоянием.
 	Упрощает код контекста.

	МИНУСЫ
	Может неоправданно усложнить код, если состояний мало и они редко меняются.

	 Применимость
	Когда у вас есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния, причём типов состояний много, и их код часто меняется.
	Когда код класса содержит множество больших, похожих друг на друга, условных операторов, которые выбирают поведения в зависимости от текущих значений полей класса.
	Когда вы сознательно используете табличную машину состояний, построенную на условных операторах, но вынуждены мириться с дублированием кода для похожих состояний и переходов.

	Пример на практике
	К примеру, создаётся программное обеспечение для управление музыкальным плеером. В зависимости от состояния плеера, кнопки будут реализовывать разный функционал.
	К примеру, в заблокированном состоянии доступна только кнопка блокировки, которая разблокирует плеер, остальные  ничего делать не будут. В состоянии проигрывания музыки
	кнопка блокировки будет блокировать плеер, кнопки увеличения/уменьшения громкости увеличивать/уменьшать громкость, кнопки следующий/предыдущий переключает треки
	вперёд/назад. Эти состояния выносятся в отдельные классы и главный класс контекста вызывает методы нужного класса в зависимости от состояния.
*/

//Реализация торгового автомата

// Торговый автомат
type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State

	currentState State

	itemCount int
	itemPrice int
}

// Создаём новый тоговый автомат
func newVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

// Запрашиваем товар из автомата
func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

// Добавляем товар в автомат
func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

// Добавляем деньги в автомат
func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

// Выдаём товар из автомата
func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

// Устанавливаем нужное состояние
func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

// Увеличиваем количество товаров в автомате
func (v *VendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

// Все состояние удовлетворяют одному интерфейсу
type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

// Состояние когда товара нет в автомате
type NoItemState struct {
	vendingMachine *VendingMachine
}

// Возвращаем ошибку: товара нет в наличии
func (i *NoItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

// Добавляем товар
func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

// Вносим деньги
func (i *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

// Получаем товар
func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}

// Сосояние того, что товар есть в автомате
type HasItemState struct {
	vendingMachine *VendingMachine
}

// Запрашиваем товар
func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

// Добавляем довар
func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

// Вносим деньги
func (i *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

// Получаем товар
func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}

// Состояние запроса товара
type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

// Запрашиваем товар
func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("Item already requested")
}

// Добавляем  товар
func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

// Вносим деньги
func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

// Получаем товар
func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}

// Состояние наличия денег
type HasMoneyState struct {
	vendingMachine *VendingMachine
}

// Запрашиваем товар
func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

// Добавляем  товар
func (i *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

// Вносим деньги
func (i *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

// Получаем товар
func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
