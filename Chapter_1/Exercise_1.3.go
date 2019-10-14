package main

/*
* Поэкспериментируйте с измерением разницы времени выполнения потенциально
* неэффективных версий и версий с применением strings.Join.
* (В разделе 1.6 демонстрируется часть пакета time, а в разделе 11.4 - как
* написать тест производительности для ее систематической оценки)
*/

// Tested for < 50 elements

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	// ----------------------------------
	// Time - 996.7µs echo v1
	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)
	// ----------------------------------

	// ----------------------------------
	// Time - 998.4µs echo v2
	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)
	// ----------------------------------


	// ---------------------------------- 
	// Time - 987.3µs echo v3
	// fmt.Println(strings.Join(os.Args[1:], " "))
	// ----------------------------------

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
