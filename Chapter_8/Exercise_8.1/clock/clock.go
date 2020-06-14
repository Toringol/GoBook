/*
* Измените программу clock2 таким образом, чтобы она
* принимала номер порта, и напишите программу clockwall,
* которая действует в качестве клиента нескольких серверов
* одновременно, считывая время из каждого и выводя результаты
* в виде таблицы, сродни настенным часам, которые можно увидеть
* в некоторых офисах. Если у вас есть доступ к географиски
* разнесенным компьютерам, запустите экземпляры серверов удаленно;
* в противном случае запустите локальные экземпляры на разных портах
* с поддельными часовыми поясами.
 */

package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

var (
	port     = flag.Int("port", 8080, "portNumber of server")
	timezone = flag.Int("timezone", 1, "timezone of server")
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		loc := time.FixedZone("UTC-8", (*timezone)*60*60)
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}
