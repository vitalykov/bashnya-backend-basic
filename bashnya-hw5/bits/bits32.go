//go:build 386 || arm || mips || mipsle

package bits

const mask = 0x7FFFFFFF

func Set(n int, pos int) int {
	n |= 1 << pos
	return n
}

func Reset(n int, pos int) int {
	n &= mask ^ (1 << pos)
	return n
}

func Flip(n int, pos int) int {
	n ^= 1 << pos
	return n
}
