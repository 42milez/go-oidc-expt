package config

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/caarlos0/env/v8"
)

func TestNew(t *testing.T) {
	wantPort := 1234
	t.Setenv("PORT", fmt.Sprint(wantPort))
	got, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if got.Port != wantPort {
		t.Errorf("got = %d; want = %d", got.Port, wantPort)
	}

	wantErr := env.AggregateError{}
	t.Setenv("DB1_PORT", "INVALID_DB1_PORT")
	_, err = New()
	if !errors.As(err, &wantErr) {
		t.Errorf("got = %+v; want = %+v", reflect.TypeOf(err), reflect.TypeOf(wantErr))
	}
}
