package alias

import (
	"math"
	"testing"
)

func BenchmarkSampling(b *testing.B) {
	D := []float64{1.0 / 3.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0}
	F := make([]float64, len(D))
	A := New(D, nil)
	n := float64(0)
	for i := 0; i < b.N; i++ {
		F[A.Roll()]++
		n++
	}
	mse := float64(0)
	for i := range F {
		r := F[i]/n - D[i]
		mse += r * r / float64(len(F))
	}
	if b.N != 1 {
		b.Logf("sampled vs. original distribution %% error over %d samples = %5.2f%%\n", b.N, math.Sqrt(mse)*100)
	}
}

func BenchmarkConstructor(b *testing.B) {
	D := []float64{1.0 / 3.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0}
	for i := 0; i < b.N; i++ {
		New(D, nil)
	}
}
