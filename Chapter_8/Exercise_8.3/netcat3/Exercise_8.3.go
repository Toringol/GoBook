/*
* В программе netcat3 значение интерфейса conn имеет
* конкретный тип *net.TCPConn, который предствляет TCP-
* соедиение. TCP-соединение состоит из двух половин, которые
* могут быть закрыты независимо с использование методов CloseRead и
* CloseWrite. Измените главную go-подпрограмму netcat3 так, чтобы
* она закрывала только записывающую половину соединения, так, чтобы
* программа продолжала выводить последние эхо от сервера reverb1 даже
* после того, как стандартный ввод будет закрыт.
 */

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.(*net.TCPConn).CloseWrite()
	<-done // wait for background goroutine to finish
	conn.(*net.TCPConn).CloseRead()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
}
