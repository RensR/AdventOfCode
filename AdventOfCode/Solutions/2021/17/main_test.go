package main

import (
	"testing"
)

func BenchmarkMain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(puzzle)
	}
}

func TestMainM(t *testing.T) {
	run(puzzle)
}
