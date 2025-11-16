package bits

const mask64 = 0x7FFFFFFFFFFFFFFF

func Set64(n int64, pos int) int64 {
	n |= 1 << pos
	return n
}

func Reset64(n int64, pos int) int64 {
	n &= mask64 ^ (1 << pos)
	return n
}

func Flip64(n int64, pos int) int64 {
	n ^= 1 << pos
	return n
}
