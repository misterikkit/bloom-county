package bitsets_test

import (
	"fmt"
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
	// fmt.Printf instead of b.Logf because we can get nice output without the -v flag
	fmt.Printf("Running benchmark with BitFieldSize=%v and BitFieldDensity=%v\n", BitFieldSize, BitFieldDensity)
	// 1 and 2 are random. 3 is a superset of 2.
	dense1 := bitsets.NewDense(BitFieldSize)
	populateRandom(dense1, numSetBits)
	dense2 := bitsets.NewDense(BitFieldSize)
	populateRandom(dense2, numSetBits)
	dense3 := bitsets.NewDense(BitFieldSize)
	bitsets.BitCopy(dense2, dense3)
	populateRandom(dense3, numSetBits)

	sparse1 := bitsets.NewSparse(BitFieldSize)
	bitsets.BitCopy(dense1, sparse1)
	sparse2 := bitsets.NewSparse(BitFieldSize)
	bitsets.BitCopy(dense2, sparse2)
	sparse3 := bitsets.NewSparse(BitFieldSize)
	bitsets.BitCopy(dense3, sparse3)

	b.Run("dense/positive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dense3.IsSuperSet(dense2)
		}
	})

	b.Run("dense/negative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dense3.IsSuperSet(dense1)
		}
	})

	b.Run("sparse/positive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sparse3.IsSuperSet(sparse2)
		}
	})

	b.Run("sparse/negative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sparse3.IsSuperSet(sparse1)
		}
	})
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
