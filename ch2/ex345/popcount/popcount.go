package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//  returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Loop returns the population count using a loop instead of an unrolled expression.
func Loop(x uint64) int {
	var result int
	for i := uint64(0); i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

// Shift64 returns the population count using a shift and check method
func Shift64(x uint64) int {
	var result uint64
	for i := 0; i < 64; i++ {
		result += x & 1
		x >>= 1
	}
	return int(result)
}

// ShiftOff returns the population count using the "shift off rightmost set bit and check" method
func ShiftOff(x uint64) int {
	var result uint64
	for x&(x-1) != x {
		result++
		x = x & (x - 1)
	}
	return int(result)
}
