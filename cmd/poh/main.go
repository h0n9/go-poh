package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/h0n9/go-poh/poh"
)

func main() {
	var wait sync.WaitGroup

	p := poh.NewPoH([]byte("hello world"))

	// wait.Add(1)
	go p.Tick(1000)

	wait.Add(1)
	go func() {
		stop := false
		reader := bufio.NewReader(os.Stdin)
		for !stop {
			fmt.Printf("> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				continue
			}
			input = strings.TrimRight(input, "\r\n")
			history := p.GetLatestHash()
			msg, err := poh.NewMsg(history, input)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(msg)
			count := p.GetCount(history)
			fmt.Printf("%x - %d\n", history, count)
		}
		wait.Done()
	}()

	wait.Wait()
}
