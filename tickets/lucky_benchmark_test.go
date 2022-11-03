package main

import "testing"

func BenchmarkLuckyN1(b *testing.B) {
	b.ResetTimer()
	lucky := Lucky{}

	for i := 0; i < b.N; i++ {
		lucky.GetLuckyCount(1)
	}
}

func BenchmarkLuckyN2(b *testing.B) {
	b.ResetTimer()
	lucky := Lucky{}

	for i := 0; i < b.N; i++ {
		lucky.GetLuckyCount(2)
	}
}

func BenchmarkLuckyN3(b *testing.B) {
	b.ResetTimer()
	lucky := Lucky{}

	for i := 0; i < b.N; i++ {
		lucky.GetLuckyCount(3)
	}
}

func BenchmarkLuckyN4(b *testing.B) {
	b.ResetTimer()
	lucky := Lucky{}

	for i := 0; i < b.N; i++ {
		lucky.GetLuckyCount(4)
	}
}
