package bitsets

import (
	sparsebs "github.com/js-ojus/sparsebitset"
	densebs "github.com/willf/bitset"
)

type BitSet interface {
	Set(n uint) BitSet
	Test(n uint) bool
	IsSuperSet(other BitSet) bool
	NextSet(n uint) (uint, bool)
}

func NewDense(length uint) BitSet {
	return dense{densebs.New(length)}
}

type dense struct {
	*densebs.BitSet
}

func (d dense) Set(n uint) BitSet {
	d.BitSet.Set(n) // ignore return val
	return d
}
func (d dense) Test(n uint) bool             { return d.BitSet.Test(n) }
func (d dense) IsSuperSet(other BitSet) bool { return d.BitSet.IsSuperSet(other.(dense).BitSet) }
func (d dense) NextSet(n uint) (uint, bool)  { return d.BitSet.NextSet(n) }

func NewSparse(length uint) BitSet {
	return sparse{sparsebs.New(uint64(length))}
}

type sparse struct {
	*sparsebs.BitSet
}

func (s sparse) Set(n uint) BitSet {
	s.BitSet.Set(uint64(n)) // ignore return val
	return s
}
func (s sparse) Test(n uint) bool             { return s.BitSet.Test(uint64(n)) }
func (s sparse) IsSuperSet(other BitSet) bool { return s.BitSet.IsSuperSet(other.(sparse).BitSet) }

func (s sparse) NextSet(n uint) (uint, bool) {
	bit, ok := s.BitSet.NextSet(uint64(n))
	return uint(bit), ok
}

// BitCopy copies the set bits from one BitSet to another. It is done bit-by-bit
// so that sparse and dense sets can be converted.
func BitCopy(from, to BitSet) {
	// TODO: call to.BitSet.Clear
	bit, ok := from.NextSet(0)
	for ok {
		to.Set(bit)
		bit, ok = from.NextSet(bit + 1)
	}
}

/*
All() bool
Any() bool
BinaryStorageSize() int
Cardinality() uint64
Clear(n uint64) BitSet
ClearAll() BitSet
Clone() BitSet
Complement() BitSet
Copy(c BitSet) int
Count() uint64
Difference(c BitSet) BitSet
DifferenceCardinality(c BitSet) (uint64, error)
Equal(c BitSet) bool
Flip(n uint64) BitSet
InPlaceDifference(c BitSet) BitSet
InPlaceIntersection(c BitSet) BitSet
InPlaceSymmetricDifference(c BitSet) BitSet
InPlaceUnion(c BitSet) BitSet
Intersection(c BitSet) BitSet
IntersectionCardinality(c BitSet) (uint64, error)
IsEmpty() bool
IsStrictSuperSet(c BitSet) bool
IsSuperSet(c BitSet) bool
Len() int
NextSet(n uint64) (uint64, bool)
None() bool
Set(n uint64) BitSet
SetTo(n uint64, val bool) BitSet
//  SymmetricDifference(c BitSet) BitSet
//  SymmetricDifferenceCardinality(c BitSet) (uint64, error)
Test(n uint64) bool
Union(c BitSet) BitSet
UnionCardinality(c BitSet) (uint64, error)
*/
