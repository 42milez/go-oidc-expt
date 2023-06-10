package auth

import (
	"bytes"
	"testing"
)

func TestEmbed(t *testing.T) {
	want := []byte("-----BEGIN PUBLIC KEY-----")
	if !bytes.Contains(publicKey, want) {
		t.Errorf("want = %s; got = %s", want, publicKey)
	}

	want = []byte("-----BEGIN PRIVATE KEY-----")
	if !bytes.Contains(privateKey, want) {
		t.Errorf("want = %s; got = %s", want, privateKey)
	}
}
