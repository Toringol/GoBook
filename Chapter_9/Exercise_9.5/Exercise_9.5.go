/*
* Напишите программу с двумя go-подпрограммами, которая
* отправляет сообщения назад и вперед по двум небуферизованным
* каналам наподобие мячика для пинг-понга. Какое количество
* сообщений в секунду может поддерживать эта программа?
* Не измерял, но для 100 выполняется мгновенно, для моего решения
* будет предел не по скорости, а блок по выводу в стандартный вывод
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	firstMsg := make(chan string)
	secondMsg := make(chan string)
	cancel := make(chan struct{})

	pingPongCounter := 100

	go func() {
		for i := 0; i < pingPongCounter; i++ {
			msg := "I`m Sergey"
			firstMsg <- msg
			fmt.Println(msg)
			<-secondMsg
		}
	}()

	go func() {
		for i := 0; i < pingPongCounter; i++ {
			<-firstMsg
			msg := "I`m Katya"
			secondMsg <- msg
			fmt.Println(msg)
		}
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
	}()

	<-cancel
}
