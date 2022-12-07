package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

	ПЛЮСЫ
	 Уменьшает зависимость между клиентом и обработчиками.
 	Реализует принцип единственной обязанности.
	 Реализует принцип открытости/закрытости.

	МИНУСЫ
	Запрос может остаться никем не обработанным.

	 Применимость
	 Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
	  Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
	   Когда набор объектов, способных обработать запрос, должен задаваться динамически.

	Пример на практике
	К примеру, программа, если идёт разработка система приёма онлайн-заказов. Нужно ограничить к ней доступ так, чтобы только авторизованные пользователи могли создавать
	заказы.Кроме того, определённые пользователи, владеющие правами администратора, должны иметь полный доступ к заказам. Также нужно проверять данные, передаваемые в запросе перед тем, как вносить их в систему.
	Блокировать массовые отправки формы с одним и тем же логином, чтобы предотвратить подбор паролей ботами. Форму заказа неплохо бы доставать из кеша, если она уже была однажды показана.
	С помощью данного паттерна можно превратить отельные поведения в объекты, каждую проверку перенести в отдельный класс с единственным методом выполнения. Данные запроса, над которым происходит проверка, будут передаваться в метод как аргументы.



*/

// Интерфейс обработчика
type Department interface {
	execute(*Patient)
	setNext(Department)
}

// Реализация приёмного отделения
type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

// Реализация доктора
type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

// Реализая медикаментов
type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

// Реализация кассира
type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

// Пациент
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func main() {

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
