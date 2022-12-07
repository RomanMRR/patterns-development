package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type caller struct{}

func (c caller) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	println("-- connect --")

	var buffer [1]byte
	p := buffer[:]

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			oi.LongWrite(w, scanner.Bytes())
			oi.LongWrite(w, []byte("\n"))
		}
	}()

	for {
		n, err := r.Read(p)
		if n > 0 {
			bytes := p[:n]
			print(string(bytes))
		}

		if nil != err {
			break
		}

	}

	println("-- disconnect --")
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var timeout time.Duration
	var isStopped bool
	flag.DurationVar(&timeout, "timeout", time.Second*10, "timeout")
	flag.Parse()
	args := flag.Args() //Получаем IP-адресс и порт
	if len(args) < 2 {
		panic("A host and port are required")
	}
	host := args[0]
	port := args[1]
	fmt.Println(timeout)

	//Подключаемся к нужному серверу

	var timeExceed bool

	go func() {
		<-time.After(timeout)
		timeExceed = true
	}()

	go func() {
		<-sigChan
		println("The client closes the connection")
		os.Exit(0)

	}()
ForLoop:
	for {
		switch {
		case isStopped:
			break ForLoop
		case timeExceed:
			println("Connection time exceeded")
			break ForLoop
		default:
			err := telnet.DialToAndCall(fmt.Sprintf("%s:%s", host, port), caller{})
			if err != nil {
				continue
			}
		}

	}
}
