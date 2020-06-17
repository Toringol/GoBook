/*
* Упражнение 8.12. Заставьте широковещательюну go-подпрограмму
* сообщать текущее множество клиентов вновь подключенному клиенту.
* Для этого требуется, чтобы множество clients и каналы entering
* и leaving записывали также имена клиентов.
*
* Упражнение 8.13. Заставьте сервер отключать простаивающих клиентов,
* которые не прислали ни одного сообщения за последние 5 минут.
* Указание: вызов conn.Close() в другой go-подпрограмме деблокирует
* активный вызов Read, такой, как выполняемый вызовом input.Scan().
*
* Упражнение 8.14. Измените сетевой протокол чат-сервера так, чтобы каждый
* клиент предоставлял при подключении свое имя. Используйте это имя вместо
* сетевого адреса в префиксе сообщения.
*
* Упражнение 8.15. Ошибка своевременного чтения данных любой клиентской
* программой в конечном итоге вызывает сбой всех клиентских программ. Измените
* широковещатель таким образом, чтобы в случае, когда go-подпрограмма передачи
* сообщения клиентам не готова принять сообщение, он вместо ожидания пропускал
* сообщение. Кроме того, добавьте буферизацию каждого канала исходящих сообщений
* клиента, чтобы большинство сообщений не удалялись. Широковещатель должен использовать
* неблокирующее отправление в этот канал.
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

const timeout = 5 * time.Minute
const buffer = 20

type client struct {
	name string
	ch   chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				select {
				case cli.ch <- msg:
				default:
				}
			}

		case cli := <-entering:
			var currentClients []string
			for cl := range clients {
				currentClients = append(currentClients, cl.name)
			}

			clients[cli] = true

			outputClients := strings.Join(currentClients, ", ")

			if outputClients == "" {
				outputClients = "noone"
			}

			cli.ch <- fmt.Sprintf("Clients Online %d : %s", len(currentClients), outputClients)

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	done := make(chan struct{})
	stillTalk := make(chan struct{})

	inputCh := make(chan string)
	outputCh := make(chan string, buffer)

	var who string

	_, err := io.WriteString(conn, "Enter your username: ")
	if err != nil {
		log.Fatal(err)
	}

	input := bufio.NewScanner(conn)

	for input.Scan() {
		who = input.Text()
		break
	}

	outputCh <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{who, outputCh}

	ticker := time.NewTicker(timeout)

	go clientReader(conn, inputCh, input, ticker, stillTalk, done)
	go clientWriter(conn, outputCh)

	go func() {
		for {
			select {
			case <-stillTalk:
				messages <- who + ": " + input.Text()
			case <-ticker.C:
				conn.Close()
			}
		}
	}()

	<-done
	leaving <- client{who, outputCh}
	messages <- who + " has left"
}

func clientReader(conn net.Conn, ch chan<- string, input *bufio.Scanner, ticker *time.Ticker,
	stillTalk chan struct{}, done chan struct{}) {

	for {
		if input.Scan() {
			ticker = time.NewTicker(timeout)
			stillTalk <- struct{}{}
		} else {
			done <- struct{}{}
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
