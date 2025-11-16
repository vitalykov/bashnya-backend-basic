//go:build amd64 || arm64 || mips64 || mips64le || ppc64 || ppc64le || riscv64 || s390x || wasm

package bits

const mask = 0x7FFFFFFFFFFFFFFF

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
