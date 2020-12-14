package bitsets_test

import (
	"math/rand"
	"testing"

	"github.com/misterikkit/bloom-county/bitsets"
)

const (
	BitFieldSize    = 1000 // number of bits
	BitFieldDensity = 0.1  // approximate proportion of set 1s
	numSetBits      = int(float32(BitFieldSize) * BitFieldDensity)
)

func BenchmarkBitSet_IsSuperSet(b *testing.B) {
	b.Logf("Running benchmark with BitFieldSize=%v and BitFieldDensity=%v", BitFieldSize, BitFieldDensity)

	// Use the exact same random bits in each benchmark. Even though the number of
	// bit sets may differ between benchmarks, I saw odd discrepancies when using
	// different random sets for each test.
	data := struct {
		sets  []bitsets.BitSet
		input bitsets.BitSet
	}{}
	data.input = bitsets.NewSparse(BitFieldSize)
	populateRandom(data.input, numSetBits*3) // set more bits for a chance of matching
	growData := func(n int) {
		for len(data.sets) < n {
			bs := bitsets.NewSparse(BitFieldSize) // arbitrary type since it will be copied for each test
			populateRandom(bs, numSetBits)
			data.sets = append(data.sets, bs)
		}
	}
	bsTypes := map[string]func(uint) bitsets.BitSet{
		"dense":  bitsets.NewDense,
		"sparse": bitsets.NewSparse,
	}
	for name, bs := range bsTypes {
		b.Run(name, func(b *testing.B) {
			// Create all the bit sets before resetting the benchmark timer
			sets := make([]bitsets.BitSet, b.N)
			growData(b.N)
			for i := range sets {
				sets[i] = bs(BitFieldSize)
				bitsets.BitCopy(data.sets[i], sets[i])
			}
			input := bs(BitFieldSize)
			bitsets.BitCopy(data.input, input)

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
	populateRandom(dbs, numSetBits)
	sbs := bitsets.NewSparse(BitFieldSize)
	bitsets.BitCopy(dbs, sbs)
	for i := uint(0); i < BitFieldSize; i++ {
		if d, s := dbs.Test(i), sbs.Test(i); d != s {
			t.Errorf("Bit %d differs: dense=%v, sparse=%v", i, d, s)
		}
	}
}
