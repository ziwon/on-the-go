package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

// the difficulty at which the block was mined.
// currently no need to target adjusting algorithm like bitcoin
// For examples)
// The requirements like first 20 bits of a hash must be zeros
// ensures block generating in every 10 minutes.
const targetBits = 24

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block *Block
	target *big.Int
}

// NewProofOfWork builds and returns a ProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)

	/*
	Lsh sets z = target << (256 - 24) and returns z.

	You can think of the target as an upper boundary of the validation of proof.
	0fac49161af82ed938add1d8725835cc123a1a87b1b196488360e58d4bfb51e3: Invalid
	0000010000000000000000000000000000000000000000000000000000000000: Boundary
	0000008b0f41ec78bab747864db66bcb9fb89920ee75f43fdaaeb5544f7f76ca: Valid
	*/
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Run perform a proof-of-work
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	// nonce is the counter in HashCash algorithm.
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

// Validate validates block's PoW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
