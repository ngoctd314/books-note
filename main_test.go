package main

import "testing"

func BenchmarkSum2(b *testing.B) {
	s := []int64{

		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	for i := 0; i < b.N; i++ {
		sum2(s)
	}
}

func BenchmarkSum8(b *testing.B) {
	s := []int64{

		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	for i := 0; i < b.N; i++ {
		sum8(s)
	}
}
