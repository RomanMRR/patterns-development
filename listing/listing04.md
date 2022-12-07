Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

Мы сначала получаем числа из канала, но потом получаем ошибку, так как канал не закрыт, а мы пытаемся всё из него читать, а так как мы больше в канал ничего писать не будем, то программа никогда не завершится, а потому и получаем ошибку. 

```