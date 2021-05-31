package fpe

import (
	"math/big"
)

// Encrypt function masks subject into unique random number less than modulus.
// Key should be kept secure, and tweak could be altered to generate more uniqueness.
// Subject argument will take the absolute value, so if it's -1 then it will be processed as 1.
func Encrypt(modulus, subject int64, key, tweak string, rounds int64) (*big.Int, error) {
	mod := big.NewInt(modulus)
	cipher := newEncryptor([]byte(key), []byte(tweak), mod)
	first, second, err := factor(mod)
	if err != nil {
		return nil, err
	}

	x := big.NewInt(subject)
	x.Abs(x)

	// Number of rounds is hard-coded to 3 rounds
	// higher means better encryption
	// lower means better performance
	for i := big.NewInt(0); i.Cmp(big.NewInt(rounds)) == -1; i = i.Add(i, big.NewInt(1)) {
		// formula:
		// right = x % second
		// x = first * right + (cipher.Format(i,right) + x / second) % first

		right := new(big.Int)
		right.Mod(x, second)
		// a
		a := new(big.Int)
		a.Mul(first, right)
		// b
		b := new(big.Int)
		bx := new(big.Int)
		bx.Div(x, second)
		b.Add(cipher.format(i, right), bx)
		// b%c
		bc := new(big.Int)
		bc.Mod(b, first)

		abc := new(big.Int)
		x = abc.Add(a, bc)
	}
	return x, nil
}
