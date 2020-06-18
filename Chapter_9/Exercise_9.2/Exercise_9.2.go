/*
* Перепишите пример PopCount из раздела 2.6.2 так, чтобы он
* инициализировал таблицу поиска с использованием sync.Once
* при первом к ней обращении. (В реальности стоимость сихронизации
* для таких малых и высокооптимизированных функций, как PopCount,
* является чрезмерно высокой.)
 */

package main

import "sync"

// pc[i] is the population count of i.
var pc [256]byte
var loadArrayOnce sync.Once

func loadArray() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	loadArrayOnce.Do(loadArray)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {

}
