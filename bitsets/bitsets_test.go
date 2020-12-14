package bitsets_test

import (
	"math/rand"
	"testing"

	"github.com/misterikkit/bloom-county/bitsets"
)

const (
	BitFieldSize    = 1000 // number of bits
	BitFieldDensity = 0.1  // proportion of set 1s
)

func BenchmarkBitSet_IsSubset(b *testing.B) {

	sets := make([]bitsets.BitSet, b.N)
	numSetBits := int(float32(BitFieldSize) * BitFieldDensity)
	for i := range sets {
		sets[i] = bitsets.NewDense(BitFieldSize)
		populateRandom(sets[i], numSetBits)
	}
	input := bitsets.NewDense(BitFieldSize)
	populateRandom(input, numSetBits)
	b.ResetTimer()
	for _, bs := range sets {
		if input.IsSuperSet(bs) {
			b.Log("Match")
		} else {
			b.Log("No Match")
		}
	}
}

func populateRandom(bs bitsets.BitSet, numSetBits int) {
	for i := 0; i < numSetBits; i++ {
		bs.Set(uint(rand.Uint64() % BitFieldSize))
	}
}
