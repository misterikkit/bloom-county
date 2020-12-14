package bitsets_test

import (
	"math/rand"
	"testing"

	"github.com/misterikkit/bloom-county/bitsets"
)

const (
	BitFieldSize    = 1000 // number of bits
	BitFieldDensity = 0.1  // approximate proportion of set 1s
)

func BenchmarkBitSet_IsSuperSet(b *testing.B) {
	b.Logf("Running benchmark with BitFieldSize=%v and BitFieldDensity=%v", BitFieldSize, BitFieldDensity)
	bsTypes := map[string]func(uint) bitsets.BitSet{
		"dense":  bitsets.NewDense,
		"sparse": bitsets.NewSparse,
	}
	for name, bs := range bsTypes {
		b.Run(name, func(b *testing.B) {
			// Create all the bit sets before resetting the benchmark timer
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

func TestConvertType(t *testing.T) {
	// TODO: Test reverse direction
	dbs := bitsets.NewDense(BitFieldSize)
	populateRandom(dbs, 100)
	sbs := bitsets.NewSparse(BitFieldSize)
	bitsets.BitCopy(dbs, sbs)
	for i := uint(0); i < BitFieldSize; i++ {
		if d, s := dbs.Test(i), sbs.Test(i); d != s {
			t.Errorf("Bit %d differs: dense=%v, sparse=%v", i, d, s)
		}
	}
}
