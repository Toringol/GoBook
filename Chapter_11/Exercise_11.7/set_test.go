/*
* Напишите функции производительности для Add, UnionWith
* и других методов *IntSet (раздел 6.5) с использованием больших
* псевдослучайных входных данных. Насколько быстрыми вы сможете
* сделать эти методы? Как влияет на производительность выбор размера
* слова? Насколько быстро работает IntSet по сравнению с реализацией множества
* на основе отображений?
 */

package main

import (
	"math/rand"
	"testing"
	"time"
)

const (
	max  = 1 << 24
	size = 1 << 23
)

var set1, set2 []int

func init() {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < size; i++ {
		set1 = append(set1, rng.Intn(max+1))
		set2 = append(set2, rng.Intn(max+1))
	}
}

func BenchmarkMapSetAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intset := NewMapSet()
		for _, n := range set1 {
			intset.Add(n)
		}
	}
}

func BenchmarkBitSet64Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intset := NewIntSet()
		for _, n := range set1 {
			intset.Add(n)
		}
	}
}

func BenchmarkMapSetUnionWith(b *testing.B) {
	intset1 := NewMapSet()
	intset2 := NewMapSet()
	for _, n := range set1 {
		intset1.Add(n)
	}
	for _, n := range set2 {
		intset2.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intset1.UnionWith(intset2)
	}
}

func BenchmarkBitSet64UnionWith(b *testing.B) {
	intset1 := NewIntSet()
	intset2 := NewIntSet()
	for _, n := range set1 {
		intset1.Add(n)
	}
	for _, n := range set2 {
		intset2.Add(n)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intset1.UnionWith(intset2)
	}
}
