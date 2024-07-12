package xrandom

import (
	"testing"
)

func BenchmarkMakeCryptoRandomString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := GenerateCryptoRandomString(20); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMakeCryptoRandomStringNoCache(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := GenerateCryptoRandomStringNoCache(20); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMakeMathRandomString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GenerateMathRandomString(20)
	}
}
