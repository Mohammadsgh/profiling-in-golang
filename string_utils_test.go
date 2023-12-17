package profiling_in_golang

import (
	"testing"
)

func BenchmarkCreateStringOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createStringOne(1000)
	}
}

func BenchmarkCreateTestTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createStringTwo(1000)
	}
}
