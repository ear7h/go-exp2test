# Go Exp2 Test

Benchmarks for functions which raise powers of 2 with `func(uint64)uint64` signature. Two functions are tested, one with `math.Exp2` and the other with `<<`. It's important to note that the bit shifting implementation is limited to integers and cannot accept a power greater than 63.

Implementations:
```go
func uint64exp2(p uint64) (uint64) {
	if p > 63 {
		panic("power cannot be higher than 63 for integer exp2")
	}

	return uint64(1) << p;
}

func float64exp2(p uint64) (uint64) {

	return uint64(math.Exp2(float64(p)));
}
```

## Side Notes
* The test functions are named `Benchmark{type}{power}`.
* Results from macOS w/ Intel 2.3GHz i5-7360U processor

```
go test -bench=Uint

goos: darwin
goarch: amd64
pkg: github.com/ear7h/go-exp2test
BenchmarkUint0-4    	1000000000	         2.25 ns/op
BenchmarkUint4-4    	1000000000	         2.39 ns/op
BenchmarkUint16-4   	1000000000	         2.27 ns/op
BenchmarkUint32-4   	1000000000	         2.34 ns/op
BenchmarkUint63-4   	1000000000	         2.25 ns/op
PASS
ok  	github.com/ear7h/go-exp2test	12.671s
```

```
go test -bench=Float

goos: darwin
goarch: amd64
pkg: github.com/ear7h/go-exp2test
BenchmarkFloat0-4    	50000000	        28.0 ns/op
BenchmarkFloat4-4    	50000000	        29.5 ns/op
BenchmarkFloat16-4   	50000000	        28.7 ns/op
BenchmarkFloat32-4   	50000000	        29.3 ns/op
BenchmarkFloat63-4   	50000000	        31.4 ns/op
PASS
ok  	github.com/ear7h/go-exp2test	7.499s
```
