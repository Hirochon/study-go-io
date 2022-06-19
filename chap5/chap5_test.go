package chap5

import (
	"bufio"
	"os"
	"testing"
)

const (
	X = 100
)

func BenchmarkReadChapX(b *testing.B) {
	f, _ := os.Open("./chap5/sample.txt")
	f.Close()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ReadOS(f, X)
	}
}

//// bufio用
func BenchmarkBufferChapX(b *testing.B) {
	f, _ := os.Open("./chap5/sample.txt")
	f.Close()

	bf := bufio.NewReader(f)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ReadOS(bf, X)
	}
}
