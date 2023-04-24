package bloom

import (
	"fmt"
	"math"
)

type BloomFilter struct {
	bitarray []bool
	arrSize  int
}

func (b *BloomFilter) hash1(s string) int {
	hash := 0

	for _, r := range s {
		hash = (hash + int(r))
		hash = hash % b.arrSize
	}

	return hash
}

func (b *BloomFilter) hash2(s string) int {
	hash := 1

	for i, r := range s {
		hash = hash + int(math.Pow(19, float64(i)))*int(r)
		hash = hash % b.arrSize
	}

	return hash % b.arrSize
}

func (b *BloomFilter) hash3(s string) int {
	hash := 7

	for _, r := range s {
		hash = (hash*31 + int(r)) % b.arrSize
	}

	return hash % b.arrSize
}

func (b *BloomFilter) hash4(s string) int {
	hash := 0
	p := 7

	for i := range s {
		hash += hash*7 + int(s[0])*int(math.Pow(float64(p), float64(i)))
		hash = hash % b.arrSize
	}

	return hash
}

func (b *BloomFilter) lookup(s string) bool {
	aResult := b.hash1(s)
	bResult := b.hash2(s)
	cResult := b.hash3(s)
	dResult := b.hash4(s)

	if b.bitarray[aResult] && b.bitarray[bResult] && b.bitarray[cResult] && b.bitarray[dResult] {
		return true
	} else {
		return false
	}
}

func (b *BloomFilter) Insert(s string) {
	if b.lookup(s) {
		fmt.Println(s, "is probably already present")
	} else {
		aResult := b.hash1(s)
		bResult := b.hash2(s)
		cResult := b.hash3(s)
		dResult := b.hash4(s)

		b.bitarray[aResult] = true
		b.bitarray[bResult] = true
		b.bitarray[cResult] = true
		b.bitarray[dResult] = true

		fmt.Println(s, "inserted")
	}
}

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		bitarray: make([]bool, size),
		arrSize:  size - 1,
	}
}
