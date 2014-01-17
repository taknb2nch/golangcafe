package main

import (
	"testing"
)

func BenchmarkNoCompression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoCompression()
	}
}

func BenchmarkBestSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BestSpeed()
	}
}

func BenchmarkBestCompression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BestCompression()
	}
}

func BenchmarkDefaultCompression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DefaultCompression()
	}
}
