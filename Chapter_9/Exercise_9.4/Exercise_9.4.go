/*
* Постройте конвейер, соединяющий произвольное количество
* go-подпрограмм каналами. Каково максимальное количество
* этапов конвейера, который можно создать без исчерпания
* памяти? Сколько времени длится транзит значения через
* весь конвейер?
* 1000000 - еще обрабатывает на компьютере с 8гб ОЗУ
* на 10 000 000 уже падает ОС
* примитивного значения в пределах 2 секунд на миллионе
* на 4 ядерном процессоре
 */

package main

import "fmt"

func main() {
	inPipe, outPipe := pipeline(1000000)
	inPipe <- 1
	fmt.Println(<-outPipe)
}

func pipeline(numberPipes int) (chan<- interface{}, <-chan interface{}) {
	if numberPipes < 1 {
		return nil, nil
	}
	inPipe := make(chan interface{})
	outPipe := inPipe
	for i := 1; i < numberPipes; i++ {
		prev := outPipe
		next := make(chan interface{})
		go func() {
			for val := range prev {
				next <- val
			}
			close(next)
		}()
		outPipe = next
	}

	return inPipe, outPipe
}
