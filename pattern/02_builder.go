package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern


    Применимость	
	 Когда вы хотите избавиться от большого конструктора.
	Когда ваш код должен создавать разные представления какого-то объекта. Например, деревянные и железобетонные дома.
    Когда вам нужно собирать сложные составные объекты, например, деревья.

    ПЛЮСЫ
    Позволяет создавать продукты пошагово.
    Позволяет использовать один и тот же код для создания различных продуктов.
    Изолирует сложный код сборки продукта от его основной бизнес-логики.

    МИНУСЫ
    Усложняет код программы из-за введения дополнительных классов.
    Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.
*/



//Определяем общий интерфейс для строителей
type IBuilder interface {
    setWindowType()
    setDoorType()
    setNumFloor()
    getHouse() House
}

//Определяем, какого строителя будем использовать
func getBuilder(builderType string) IBuilder {
    if builderType == "normal" { //Для нормальных домов
        return newNormalBuilder()
    }

    if builderType == "igloo" { //Для снежных домов
        return newIglooBuilder()
    }
    return nil
}

//Строитель нормальных домов
type NormalBuilder struct {
    windowType string
    doorType   string
    floor      int
}

func newNormalBuilder() *NormalBuilder {
    return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
    b.windowType = "Wooden Window"
}

func (b *NormalBuilder) setDoorType() {
    b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
    b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
    return House{
        doorType:   b.doorType,
        windowType: b.windowType,
        floor:      b.floor,
    }
}

//Строитель снежных домов
type IglooBuilder struct {
    windowType string
    doorType   string
    floor      int
}

func newIglooBuilder() *IglooBuilder {
    return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
    b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
    b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
    b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
    return House{
        doorType:   b.doorType,
        windowType: b.windowType,
        floor:      b.floor,
    }
}

//Получаемый результат: дом
type House struct {
    windowType string
    doorType   string
    floor      int
}

//Определяем директора, он будет управлять строителями, говорить им, что делать

type Director struct {
    builder IBuilder
}

func newDirector(b IBuilder) *Director {
    return &Director{
        builder: b,
    }
}

func (d *Director) setBuilder(b IBuilder) {
    d.builder = b
}

func (d *Director) buildHouse() House {
    d.builder.setDoorType()
    d.builder.setWindowType()
    d.builder.setNumFloor()
    return d.builder.getHouse()
}

func main() {
    normalBuilder := getBuilder("normal")
    iglooBuilder := getBuilder("igloo")

	//Строим обычный дом
    director := newDirector(normalBuilder) //Директор получает строителя обычных домов
    normalHouse := director.buildHouse() //И говорит ему что делать

	//Дом готов
    fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
    fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
    fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

    director.setBuilder(iglooBuilder) //Затем директор управляет строителем снежных домов
    iglooHouse := director.buildHouse() //И говорит ему что делать

	//Снежный дом готов
    fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
    fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
    fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}