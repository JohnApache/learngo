package main

import (
	fib1 "mypkg/fib1"
	fib2 "mypkg/fib2"
	"testing"
)

func BenchmarkTestA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib1.Run()
	}
}
func BenchmarkTestB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib2.Run()
	}
}
