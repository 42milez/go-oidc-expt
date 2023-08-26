package xstring

import (
	"bytes"
	"fmt"
	"testing"
)

var data = []byte("Measure performance in converting to string.")

func BenchmarkNewBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bytes.NewBuffer(data).String()
	}
}

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", data)
	}
}

func BenchmarkString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string(data)
	}
}

func BenchmarkByteToString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ByteToString(data)
	}
}
