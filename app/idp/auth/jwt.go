package auth

import (
	_ "embed"
)

//go:embed cert/secret.pem
var privateKey []byte

//go:embed cert/public.pem
var publicKey []byte
