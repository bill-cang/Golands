package main

import "testing"

func BenchmarkFeibonacci1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Feibonacci1(10)
	}
}

func BenchmarkFeibonacci2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Feibonacci2(10)
	}
}


