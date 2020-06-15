/*
* Используя инструкцию select, добавьте к эхо серверу из
* раздела 8.3 тайм-аут, чтобы он отключал любого клиента,
* который ничего не передает в течение 10 секунд.
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	textCh := make(chan string)
	ticker := time.NewTicker(10 * time.Second)
	done := make(chan bool)

	go func() {
		for input.Scan() {
			textCh <- input.Text()
			ticker = time.NewTicker(10 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case text := <-textCh:
				echo(c, text, 1*time.Second)
			case <-ticker.C:
				done <- true
			}
		}
	}()

	<-done
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
