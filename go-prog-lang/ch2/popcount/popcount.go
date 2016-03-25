package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

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

// ex 2.3
func PopCountLoop(x uint64) int {
	var sum byte
	for i := 0; i < 8; i++ {
		sum += pc[byte(x>>uint64(i*8))]
	}
	return int(sum)
}

// ex 2.4
func PopCountShift64(x uint64) int {
	var sum int
	for ; x != 0; x = x >> 1 {
		if (x & 1) == 1 {
			sum++
		}
	}
	return sum
}

// ex 2.5
func PopCountClearRightmostBit(x uint64) int {
	var sum int
	for x != 0 {
		sum++
		x = x & (x - 1)
	}
	return sum
}
