package bits

const mask32 = 0x7FFFFFFF

func Set32(n int32, pos int) int32 {
	n |= 1 << pos
	return n
}

func Reset32(n int32, pos int) int32 {
	n &= mask32 ^ (1 << pos)
	return n
}

func Flip32(n int32, pos int) int32 {
	n ^= 1 << pos
	return n
}
