Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false
Интерфейс состоит из двух указателей на метаданные типа и сами данные.
При возврате из функции ошибки Go оборачивает nil-указатель *os.PathError в не nil-интерфейс error, а потому указатель на метаданные в интерфейсе error будет не пустым, там как там будут содержаться данные об os.PathError, а значит сам интерфейс не будет равен nil. Но вот данные в интерфейсе равны nil, поэтому на экран и выводится nil.

```
