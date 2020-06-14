/*
* Модифицируйте сервер reverb2 так, чтобы он использовад по
* одному объекту sync.WaitGroup для каждого соединения для посчета
* количества активных go-подпрограмм echo. Когда он обнуляется, закрывайте
* пишущую половину TCP-соедиения, как описано в упражнении 8.3. Убедитесь, что
* вы изменили клиентскую программу netcat3 из этого упражнения так, чтобы она
* ожидала последние ответы от параллельных go-подпрограмм сервера даже после
* закрытия стандартного ввода.
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
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
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			echo(c, input.Text(), 1*time.Second)
		}()
	}

	wg.Wait()
	c.(*net.TCPConn).CloseWrite()
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
