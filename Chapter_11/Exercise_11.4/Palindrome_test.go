/*
* Измените функцию randomPalindrome, так чтобы
* она проверяла, как функция IsPalindrome
* обрабатывает пробельные символы и знаки
* пунктуации.
 */

package main

import (
	"math/rand"
	"testing"
	"time"
)

var points = []rune{
	' ', ',', '.', '!', '?', '　', '、', '。', '！', '？',
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}

	if n == 0 {
		return string(runes)
	}

	pos := rng.Intn(n)

	runes = append(runes[:pos], append([]rune{points[rng.Intn(len(points))]}, runes[pos:]...)...)

	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
