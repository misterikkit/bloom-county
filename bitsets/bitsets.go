package bitsets

import densebs "github.com/willf/bitset"

type BitSet interface {
	IsSuperSet(other BitSet) bool
	Set(n uint) BitSet
}

func NewDense(length uint) BitSet {
	return dense{densebs.New(length)}
}

type dense struct {
	*densebs.BitSet
}

func (d dense) IsSuperSet(other BitSet) bool {
	return d.BitSet.IsSuperSet(other.(dense).BitSet)
}

func (d dense) Set(n uint) BitSet {
	d.BitSet.Set(n) // ignore return val
	return d
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
//  ReadFrom(r io.Reader) (int64, error)
Set(n uint64) BitSet
SetTo(n uint64, val bool) BitSet
//  SymmetricDifference(c BitSet) BitSet
//  SymmetricDifferenceCardinality(c BitSet) (uint64, error)
Test(n uint64) bool
Union(c BitSet) BitSet
UnionCardinality(c BitSet) (uint64, error)
//  WriteTo(w io.Writer) (int64, error)
*/
