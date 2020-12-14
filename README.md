# BitSet Benchmark

The benchmark in `bitsets` folder measures the performance of `IsSuperSet` in
[dense][1] and [sparse][2] implementations of BitSet. Run test with `go test ./bitsets/ -run xx -bench . -v`

[1]: https://github.com/willf/bitset
[2]: https://github.com/js-ojus/sparsebitset

Initial results:

```
goos: linux
goarch: amd64
pkg: github.com/misterikkit/bloom-county/bitsets
BenchmarkBitSet_IsSuperSet
    bitsets_test.go:16: Running benchmark with BitFieldSize=1000 and BitFieldDensity=0.1
BenchmarkBitSet_IsSuperSet/sparse
    bitsets_test.go:42: Match ratio is 0/1 hit/miss
    bitsets_test.go:42: Match ratio is 0/100 hit/miss
    bitsets_test.go:42: Match ratio is 0/10000 hit/miss
    bitsets_test.go:42: Match ratio is 5/999995 hit/miss
    bitsets_test.go:42: Match ratio is 5/5336252 hit/miss
BenchmarkBitSet_IsSuperSet/sparse-4              5336257               255 ns/op
BenchmarkBitSet_IsSuperSet/dense
    bitsets_test.go:42: Match ratio is 0/1 hit/miss
    bitsets_test.go:42: Match ratio is 0/100 hit/miss
    bitsets_test.go:42: Match ratio is 0/10000 hit/miss
    bitsets_test.go:42: Match ratio is 0/1000000 hit/miss
    bitsets_test.go:42: Match ratio is 0/41546210 hit/miss
BenchmarkBitSet_IsSuperSet/dense-4              41546210              1290 ns/op
PASS
ok      github.com/misterikkit/bloom-county/bitsets     243.416s
```
