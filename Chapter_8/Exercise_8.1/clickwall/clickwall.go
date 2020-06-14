package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type portsArray []string

func (i *portsArray) String() string {
	return "ports"
}

func (i *portsArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var ports portsArray

func main() {
	flag.Var(&ports, "ports", "Ports to connect to tcp servers")
	flag.Parse()

	for i, port := range ports {
		conn, err := net.Dial("tcp", "localhost:"+port)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn, i)
	}

	for {
		time.Sleep(time.Minute)
	}
}

func mustCopy(dst io.Writer, src io.Reader, tabCounter int) {
	s := bufio.NewScanner(src)
	for s.Scan() {
		fmt.Fprintf(dst, strings.Repeat("\t", tabCounter*2)+s.Text()+"\n")
	}

	if s.Err() != nil {
		log.Printf("%s", s.Err())
	}
}
