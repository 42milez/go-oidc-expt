package xutil

import (
	"testing"
)

func BenchmarkMakeCryptoRandomString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := MakeCryptoRandomString(20); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMakeCryptoRandomStringNoCache(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := MakeCryptoRandomStringNoCache(20); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMakeMathRandomString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MakeMathRandomString(20)
	}
}
