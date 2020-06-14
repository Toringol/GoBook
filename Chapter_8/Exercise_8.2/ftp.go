/*
* Реализуйте параллельный FTP-сервер. Сервер должен
* интерпретировать команды от каждого клиента, такие как
* cd для именения каталога, ls для вывода списка файлов в
* каталоге, get для отправки содержимого файла и close для
* закрытия соединения. В качестве клиента можно использовать
* стандартную команду ftp или написать собственную программу.
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)

	for input.Scan() {
		cmd := input.Text()
		if strings.HasPrefix(cmd, "cd") {
			changeDir(cmd, c)
		} else if strings.HasPrefix(cmd, "ls") {
			list(cmd, c)
		} else if strings.HasPrefix(cmd, "get") {
			fileInfo(cmd, c)
		} else if cmd == "close" {
			return
		} else {
			_, err := io.WriteString(c, "Command not allowed!\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if input.Err() != nil {
		log.Printf("%s", input.Err())
	}
}

func changeDir(cmd string, c net.Conn) {
	path := strings.Split(cmd, " ")[1]

	err := os.Chdir(path)
	if err != nil {
		log.Fatal(err)
	}
}

func list(cmd string, c net.Conn) {
	path := strings.Split(cmd, " ")[1]

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(c, "---------------------------\n")

	for _, f := range files {
		_, err := io.WriteString(c, f.Name()+" Dir: "+strconv.FormatBool(f.IsDir())+"\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Fprintf(c, "---------------------------\n")
}

func fileInfo(cmd string, c net.Conn) {
	fileName := strings.Split(cmd, " ")[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(c, "---------------------------\n")

	if _, err := io.Copy(c, file); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(c, "---------------------------\n")
}

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
