Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Функция test() возвращает структуру customError, которая удовлетворяет интерфейсу error (реализован метод Error()). 
Err - это переменная интерфейса, она состоит из двух указателей: на данные и на тип этих данных, а потому переменная
err не равна nil, данные внутри неё равны nil, а вот тип данных не равен nil. А потому будет выведено:
error

```
