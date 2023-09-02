package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/42milez/go-oidc-server/pkg/xrandom"
)

var (
	lenFlag = flag.Int("len", 0, "the length of random string")
)

func main() {
	flag.Parse()

	if *lenFlag == 0 {
		log.Fatal("length must be greater than 0")
	}

	ret, err := xrandom.MakeCryptoRandomString(*lenFlag)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(ret)
}
