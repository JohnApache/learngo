package goc

import (
	"testing"
)

func testA() {
	grc := Goc(10, 10)
	for i := 0; i < 100; i++ {
		grc.Push(TaskA)
	}
	grc.Join()
	//	fmt.Println(2)
}

func testB() {
	grc := Goc(5, 5)
	for i := 0; i < 100; i++ {
		grc.Push(TaskA)
	}
	grc.Join()
	//	fmt.Println(2)
}

func BenchmarkTestA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testA()
	}
}

func BenchmarkTestB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testB()
	}
}
