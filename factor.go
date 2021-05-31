package fpe

import (
	"math/big"
)

type primeGenerator struct {
	seq               int
	markedNotPrimeMap map[int][]int
}

func newPrimeGenerator() *primeGenerator {
	return &primeGenerator{
		seq:               1,
		markedNotPrimeMap: make(map[int][]int),
	}
}

func (p *primeGenerator) next() int {
	for {
		p.seq++
		if _, ok := p.markedNotPrimeMap[p.seq]; !ok {
			p.markedNotPrimeMap[p.seq*p.seq] = []int{p.seq}
			return p.seq
		}
		primes := p.markedNotPrimeMap[p.seq]
		for _, v := range primes {
			nextMultipleOfPrime := v + p.seq
			if p.markedNotPrimeMap[nextMultipleOfPrime] != nil {
				p.markedNotPrimeMap[nextMultipleOfPrime] = append(p.markedNotPrimeMap[nextMultipleOfPrime], v)
			} else {
				p.markedNotPrimeMap[nextMultipleOfPrime] = []int{v}
			}
		}
		p.markedNotPrimeMap[p.seq] = nil
	}
}

func factor(n *big.Int) (a, b *big.Int, err error) {
	if n.Cmp(big.NewInt(0)) == -1 {
		return a, b, ErrNegativeArgs
	}
	primes := newPrimeGenerator()
	a = big.NewInt(1)
	b = big.NewInt(1)
	var p *big.Int

	for k := 0; n.Int64() > 1; k++ {
		p = big.NewInt(int64(primes.next()))
		if n.Int64()%p.Int64() == 0 {
			for n.Int64()%p.Int64() == 0 {
				b.Mul(b, p)
				if a.Cmp(b) == -1 {
					b, a = a, b
				}
				n.Div(n, p)
			}
		}
	}

	return a, b, nil
}
