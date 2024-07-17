package api

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/42milez/go-oidc-expt/cmd/config"
)

func Test_SecretKey(t *testing.T) {
	t.Parallel()
	email := fmt.Sprintf("test@%s", config.Issuer)
	want := regexp.MustCompile(fmt.Sprintf("otpauth://totp/%s:%s?secret=[A-Z2-7]+&issuer=%s", config.Issuer, email, config.Issuer))
	got, err := SecretKey(email)
	if err != nil {
		t.Fatal(err)
	}
	if want.MatchString(got) {
		t.Errorf("SecretKey() = %v; want = %v", got, want)
	}
}

func Test_Verify_RFC6238_TestVector1(t *testing.T) {
	t.Parallel()
	verifier := &TOTP{
		Key: "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ",
		now: func() int64 { return 59 },
	}
	got, _ := verifier.Verify("287082")
	if !got {
		t.Errorf("Verify() = %v; want = %v", got, true)
	}
}

func Test_Verify_RFC6238_TestVector4(t *testing.T) {
	t.Parallel()
	verifier := &TOTP{
		Key: "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ",
		now: func() int64 { return 1111111109 },
	}
	got, _ := verifier.Verify("081804")
	if !got {
		t.Errorf("Verify() = %v; want = %v", got, true)
	}
}

func Test_Verify_RFC6238_TestVector7(t *testing.T) {
	t.Parallel()
	verifier := &TOTP{
		Key: "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ",
		now: func() int64 { return 1111111111 },
	}
	got, _ := verifier.Verify("050471")
	if !got {
		t.Errorf("Verify() = %v; want = %v", got, true)
	}
}

func Test_Verify_RFC6238_TestVector10(t *testing.T) {
	t.Parallel()
	verifier := &TOTP{
		Key: "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ",
		now: func() int64 { return 1234567890 },
	}
	got, _ := verifier.Verify("005924")
	if !got {
		t.Errorf("Verify() = %v; want = %v", got, true)
	}
}

func Test_Verify_RFC6238_TestVector13(t *testing.T) {
	t.Parallel()
	verifier := &TOTP{
		Key: "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ",
		now: func() int64 { return 2000000000 },
	}
	got, _ := verifier.Verify("279037")
	if !got {
		t.Errorf("Verify() = %v; want = %v", got, true)
	}
}

func Test_Verify_RFC6238_TestVector16(t *testing.T) {
	t.Parallel()
	verifier := &TOTP{
		Key: "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ",
		now: func() int64 { return 20000000000 },
	}
	got, _ := verifier.Verify("353130")
	if !got {
		t.Errorf("Verify() = %v; want = %v", got, true)
	}
}
