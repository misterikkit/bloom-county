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

func BenchmarkBitSet_IsSuperSet(b *testing.B) {
	b.Logf("%v: N=%v", b.Name(), b.N)

	bsTypes := map[string]func(uint) bitsets.BitSet{
		"dense":  bitsets.NewDense,
		"sparse": bitsets.NewSparse,
	}
	for name, bs := range bsTypes {
		b.Run(name, func(b *testing.B) {
			b.Logf("%v: N=%v", b.Name(), b.N)

			sets := make([]bitsets.BitSet, b.N)
			numSetBits := int(float32(BitFieldSize) * BitFieldDensity)
			for i := range sets {
				sets[i] = bs(BitFieldSize)
				populateRandom(sets[i], numSetBits)
			}
			input := bs(BitFieldSize)
			populateRandom(input, numSetBits*3) // set more bits for a chance of matching
			b.ResetTimer()
			hit, miss := 0, 0
			for _, bs := range sets {
				if input.IsSuperSet(bs) {
					hit++
				} else {
					miss++
				}
			}
			b.Logf("Match ratio is %d/%d hit/miss", hit, miss)
		})
	}
}

func populateRandom(bs bitsets.BitSet, numSetBits int) {
	for i := 0; i < numSetBits; i++ {
		bs.Set(uint(rand.Uint64() % BitFieldSize))
	}
}
