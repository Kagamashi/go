package main

import (
	"testing"
)

/* BENCHMARKING (performance)
Go supports benchmarking tests to measure the performance of functions using testing package.
It helps identify performance bottlenecks and optimize critical code paths.
- go test -bench to run all benchmarks in a package
- Benchmark functions must start with Benchmark and accept a *testing.B parameter

- Common functions:
b.ResetTimer() - resets the timer so that only the code after this point is benchmarked
b.StopTimer() and b.StartTimer() - pause and resume timing during the benchmark
b.N - the number of iterations the benchmark will run, used to control the loop
*/

func IntMinBen(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BenchmarkIntMin(b *testing.B) { // Benchmark tests typically in _test.go files and are named beginning with Benchmark keyword.

	for i := 0; i < b.N; i++ { // Testng runner executes each benchmark function several times, increasing b.N on each run until it collects a precise measurement.
		IntMinBen(1, 2)
	}

	/*
	> go test -bench=.

	goos: darwin
	goarch: arm64
	pkg: examples/testing
	BenchmarkIntMin-8 1000000000 0.3136 ns/op
	PASS
	ok      examples/testing-and-benchmarking    0.351s
	*/

}
