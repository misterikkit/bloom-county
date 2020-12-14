# BitSet Benchmark

The benchmark in `bitsets` folder measures the performance of `IsSuperSet` in
[dense][1] and [sparse][2] implementations of BitSet. Run test with `go test ./bitsets/ -run xx -bench . -benchmem`

[1]: https://github.com/willf/bitset
[2]: https://github.com/js-ojus/sparsebitset

Initial results:

```
Running benchmark with BitFieldSize=1000 and BitFieldDensity=0.1
goos: linux
goarch: amd64
pkg: github.com/misterikkit/bloom-county/bitsets
BenchmarkBitSet_IsSuperSet/dense/positive-4         	 1271924	       929 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitSet_IsSuperSet/dense/negative-4         	96602800	        12.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitSet_IsSuperSet/sparse/positive-4        	11112118	       105 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitSet_IsSuperSet/sparse/negative-4        	11898327	        98.2 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/misterikkit/bloom-county/bitsets	5.916s
```

```
Running benchmark with BitFieldSize=10000 and BitFieldDensity=0.1
goos: linux
goarch: amd64
pkg: github.com/misterikkit/bloom-county/bitsets
BenchmarkBitSet_IsSuperSet/dense/positive-4         	  120638	      9841 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitSet_IsSuperSet/dense/negative-4         	83552077	        13.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitSet_IsSuperSet/sparse/positive-4        	 1237958	       933 ns/op	       0 B/op	       0 allocs/op
BenchmarkBitSet_IsSuperSet/sparse/negative-4        	 1274845	       945 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/misterikkit/bloom-county/bitsets	6.723s
```
