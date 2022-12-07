package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
)

func server() error {
	fmt.Printf("Telnet server start on port:%d\n", 5555)
	return telnet.ListenAndServe(fmt.Sprintf(":%d", 5555), handler{})
}

type handler struct{}

// func (h handler) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for scanner.Scan() {
// 		oi.LongWrite(w, scanner.Bytes())
// 		oi.LongWrite(w, []byte("\n"))
// 	}
// }

func (h handler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
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
	err := server()
	if nil != err {
		//@TODO: Handle this error better.
		panic(err)
	}
}
