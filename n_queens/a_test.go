package n_queens

import (
	"testing"
	"time"

	ffmt "gopkg.in/ffmt.v1"
)

func TestA88(t *testing.T) {
	ffmt.Mark(c88(8))
}

func TestANN(t *testing.T) {
	ffmt.Mark(cNN(8))
}

func TestA1_N(t *testing.T) {
	for i := uint8(1); i != 15; i++ {
		begin := time.Now()
		ffmt.Mark("No.", i, ", Sum:", cc(i), ", ", time.Now().Sub(begin))
	}
}

func BenchmarkA8(b *testing.B) {
	for i := 0; i != b.N; i++ {
		c88(8)
	}
}

func BenchmarkEightQueen(b *testing.B) {
	for i := 0; i != b.N; i++ {
		cNN(16)
	}
}
