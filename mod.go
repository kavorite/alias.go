package alias

import (
	"math/rand"
	"time"
)

type wkq []int

func (q wkq) pop() (i int) {
	l := len(q) - 1
	i = q[l]
	q = q[:l]
	return
}

func (q wkq) push(i int) {
	q = append(q, i)
}

// T is a Vose alias table for sampling from a discrete probability
// distribution.
type T struct {
	*rand.Rand
	alias []int
	prob  []float64
}

func (vose T) Roll() int {
	i := vose.Int() % len(vose.prob)
	if vose.Float64() >= vose.prob[i] {
		return i
	}
	return vose.alias[i]
}

// New generates a new alias vose for sampling the given discrete probability
// distribution D.
func New(D []float64, src rand.Source) T {
	k := len(D)
	n := float64(k)
	// μ := 1.0 / n
	S := make(wkq, 0, k)
	L := make(wkq, 0, k)
	sigma := float64(0)
	for _, x := range D {
		sigma += x
	}
	for i := range D {
		D[i] /= sigma
	}
	push := func(i int, x float64) {
		if x < 1 {
			S.push(i)
		} else {
			L.push(i)
		}
	}
	for i, x := range D {
		push(i, x)
	}
	P := make([]float64, k)
	A := make([]int, k)
	for len(S) > 0 && len(L) > 0 {
		l := S.pop()
		g := L.pop()
		P[l] = n * D[l]
		A[l] = g
		P[g] = P[g] + P[l] - 1
		push(g, P[g])
	}
	for len(L) > 0 {
		P[L.pop()] = 1
	}
	for len(S) > 0 {
		P[S.pop()] = 1
	}
	if src == nil {
		src = rand.NewSource(time.Now().UnixNano())
	}

	return T{
		Rand:  rand.New(src),
		alias: A,
		prob:  P,
	}
}
