/*
* Ex.1
* Реализуйте следующие дополнительные методы:
* func (*IntSet) Len() int 		// Возвращает количество элементов
* func (*IntSet) Remove(x int)  // Удаляет x из множества
* func (*IntSet) Clear() 		// Удаляет все элементы множества
* func (*IntSet) Copy() *IntSet // Возвращает копию множества
*
* Ex.2
* Определите вариативный метод (*IntSet).AddAll(...int),
* который позволяет добавлять список значений, например s.Add(1, 2, 3).
*
* Ex.3
* (*IntSet).UnionWith вычисляет объединение двух множеств с помощью оператора |,
* побитового оператора ИЛИ. Реализуйте методы IntersectWith, DifferenceWith и
* SymmetricDifference для соответсвующих операций над множествами. (Симметричная разность
* двух множеств содержит элементы, имеющиеся в одном из множеств, но не в обоих одновременно.)
*
* Ex.4
* Добавьте метод Elems, который возвращает срез, содержащий элементы множества
* и годящийся для итерирования с использованием цикла по диапазону range.
*
* Ex.5
* Типом каждого слова, используемного в IntSet, является uint64, но 64-разрядная
* арифметика может быть неэффективной на 32-разрядных платформах. Измените программу так,
* чтобы она использовала тип uint, который представляет собой наиболее эффективный беззнаковый
* целочисленный тип для данной платформы. Вместо деления на 64 определите константу, в которой
* хранится эффективный размер uint в битах, 32 или 64. Для этого воспользуйтесь, возможно, слишком
* умным выражением 32<<(^uint(0)>>63).
 */

package main

import (
	"bytes"
	"fmt"
)

const uintSize = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/uintSize, uint(x%uintSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Len return counter of words in IntSet
func (s *IntSet) Len() int {
	counter := 0

	for _, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				counter++
			}
		}
	}

	return counter
}

// Elems return slice of elems from set
func (s *IntSet) Elems() (result []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				result = append(result, uintSize*i+j)
			}
		}
	}

	return result
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/uintSize, uint(x%uintSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds all non-negative values to the set
func (s *IntSet) AddAll(x ...int) {
	for _, elem := range x {
		s.Add(elem)
	}
}

// Remove remomes value from set
func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/uintSize, uint(x%uintSize)
		s.words[word] &^= 1 << bit
	}
}

// Clear remove all values from set
func (s *IntSet) Clear() {
	s.words = []uint{}
}

// Copy return copy set
func (s *IntSet) Copy() *IntSet {
	copySet := &IntSet{}

	for _, word := range s.words {
		copySet.words = append(copySet.words, word)
	}

	return copySet
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersect of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// DifferenceWith sets s to the dif of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// SymmetricDifference sets s to the xor dif of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x IntSet
	x.AddAll(1, 2, 3)
	fmt.Println(x.String())
	fmt.Println(x.Len())

	y := x.Copy()
	fmt.Println(y.String())

	res := y.Elems()
	fmt.Println(res)
	for _, elem := range res {
		fmt.Println(elem)
	}
}
