package main

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCountBook(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountEx1(x uint64) int {
	var result int
	var i uint
	for i = 0; i < 64; i++ {
		result += int(byte(x>>i) & 1)
	}
	return result
}

func PopCountEx2(x uint64) int {
	var result int
	for x != 0 {
		x = x & (x - 1)
		result++
	}
	return result
}

func main() {

}
