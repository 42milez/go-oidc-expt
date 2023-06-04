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
		t.Fatalf("cannot create config: %+v", err)
	}

	if got.Port != wantPort {
		t.Errorf("got = %d; want = %d", got.Port, wantPort)
	}

	// --------------------------------------------------

	wantEnv := "dev"

	if got.Env != wantEnv {
		t.Errorf("got = %s; want = %s", got.Env, wantEnv)
	}

	// --------------------------------------------------

	wantErr := env.AggregateError{}

	t.Setenv("DB_PORT", "INVALID_DB_PORT")

	_, err = New()

	if !errors.As(err, &wantErr) {
		t.Errorf("got = %v; want = %v", reflect.TypeOf(err), reflect.TypeOf(wantErr))
	}
}
