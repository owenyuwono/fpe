package fpe

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"hash"
	"math/big"

	"golang.org/x/text/encoding/unicode"
)

type encryptor struct {
	keyByte    []byte
	macNT      []byte
	macFactory hash.Hash
}

func newEncryptor(key, tweak []byte, modulus *big.Int) encryptor {
	// encode strings into byte array UTF16LE
	enc := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	tweakByte, err := enc.Bytes(tweak)
	if err != nil {
		panic(err)
	}
	keyenc := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	keyByte, err := keyenc.Bytes(key)
	if err != nil {
		panic(err)
	}

	// Create BigEndian binary representation
	nBin := make([]byte, 8)
	binary.BigEndian.PutUint64(nBin, modulus.Uint64())
	nBinLength := make([]byte, 8)
	binary.BigEndian.PutUint64(nBinLength, uint64(len(nBin)))
	tweakLength := make([]byte, 8)
	binary.BigEndian.PutUint64(tweakLength, uint64(len(tweakByte)))

	// Create the resulting Encryption struct
	e := encryptor{
		keyByte:    key,
		macFactory: hmac.New(sha256.New, keyByte),
	}

	// Do hash
	mac := e.macFactory
	written, err := mac.Write(nBinLength)
	if err != nil && written == 0 {
		panic(err)
	}
	written, err = mac.Write(nBin)
	if err != nil && written == 0 {
		panic(err)
	}
	written, err = mac.Write(tweakLength)
	if err != nil && written == 0 {
		panic(err)
	}
	written, err = mac.Write(tweakByte)
	if err != nil && written == 0 {
		panic(err)
	}
	e.macNT = mac.Sum(nil)

	return e
}

func (e encryptor) format(roundNumber, r *big.Int) *big.Int {
	// Initialize fixed-sized BigEndian byte arrays
	round := make([]byte, 8)
	binary.BigEndian.PutUint64(round, roundNumber.Uint64())

	rBinLength := make([]byte, 8)
	binary.BigEndian.PutUint64(rBinLength, uint64(len(round)))

	rBin := make([]byte, 8)
	binary.BigEndian.PutUint64(rBin, r.Uint64())

	// Do the hashing
	mac := e.macFactory
	mac.Reset()

	written, err := mac.Write(e.macNT)
	if err != nil && written == 0 {
		panic(err)
	}
	written, err = mac.Write(round)
	if err != nil && written == 0 {
		panic(err)
	}
	written, err = mac.Write(rBinLength)
	if err != nil && written == 0 {
		panic(err)
	}
	written, err = mac.Write(rBin)
	if err != nil && written == 0 {
		panic(err)
	}

	digest := mac.Sum(nil)
	i := new(big.Int)
	i.SetBytes(digest)
	return i
}
