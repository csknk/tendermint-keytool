package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	b64PrivKeyString string = ""
	verbose                 = false
)

type Key struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func init() {
	flag.StringVar(&b64PrivKeyString, "secret", "", "Private Key, Base 64 string")
	flag.BoolVar(&verbose, "v", false, "Verbose")
}

func main() {
	flag.Parse()
	if b64PrivKeyString != "" {
		if verbose {
			fmt.Println("Generating a public key from the supplied base64 private key...")
		}
		KeyPairFromPrivKey(b64PrivKeyString, os.Stdout)
		return
	}

	// var privKey *PrivKey
	// var err error
	// if len(os.Args) == 2 {
	// 	if verbose {
	// 		fmt.Println("Generating key pair from the supplied seed...")
	// 	}
	// 	privKey, err = NewPrivKeyFromSeed(os.Args[1])
	// } else {
	// 	if verbose {
	// 		fmt.Println("Generating pseuso-random private key and associated public key...")
	// 	}
	// 	privKey, err = NewPrivKey()
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// OutputKeyPair(privKey, os.Stdout)
}
