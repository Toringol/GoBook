/*
* Напишите функцию производительности для сравнения
* реализации PopCount из раздела 2.6.2 с вашими
* решениями упражнений 2.4 и 2.5. В какой момент
* не срабатывает даже табличное тестирование?
 */

package main

import "testing"

var (
	value = uint64(378923689236823867)
)

func BenchmarkPopCountBook(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountBook(value)
	}
}

func BenchmarkPopCountEx1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountEx1(value)
	}
}

func BenchmarkPopCountEx2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountEx2(value)
	}
}
