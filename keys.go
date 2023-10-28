package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/tendermint/tendermint/crypto/tmhash"
)

const (
	// AddressSize is the size of a pubkey address.
	AddressSize = tmhash.TruncatedSize
)

type KeyPair struct {
	Address string  `json:"address"`
	PubKey  KeyJSON `json:"pub_key"`
	PrivKey KeyJSON `json:"priv_key"`
}

type KeyJSON struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type (
	PrivKey   = ed25519.PrivateKey
	MyPrivKey ed25519.PrivateKey
)

//	func NewPrivKey() (*PrivKey, error) {
//		k := ed25519.GenerateKey()
//		return &k, nil
//	}
//
// NewPrivKeyFromSeed outputs a private key from an input seed.
// FIXME add a proper salt.
//
//	func NewPrivKeyFromSeed(seed string) (*PrivKey, error) {
//		salt := []byte("1234567890abcdef")
//		secret, err := bcrypt.GenerateFromPassword(salt, []byte(seed), 0)
//		if err != nil {
//			log.Fatal(err)
//		}
//		k := ed25519.GenPrivKeyFromSecret(secret)
//		return &k, nil
//	}
//
// OutputKeyPair outputs keypair to stdout in JSON format for Tendermint
func OutputKeyPair(key PrivKey, outfile io.Writer) {
	publicKey := key.Public()
	outputPrivKey := KeyJSON{
		Type:  "tendermint/PubKeyEd25519",
		Value: base64.StdEncoding.EncodeToString(key),
	}
	fmt.Printf("%#v\n", publicKey)
	// address := GetAddress(key)
	// Perform a type assertion to get the Ed25519 public key
	ed25519PublicKey, ok := publicKey.(ed25519.PublicKey)
	if !ok {
		fmt.Println("Error: publicKey is not an Ed25519 public key")
		return
	}
	// Convert the Ed25519 public key to a byte slice
	publicKeyBytes := []byte(ed25519PublicKey)

	kp := KeyPair{
		// Address: address,
		PrivKey: outputPrivKey,
		PubKey: KeyJSON{
			Type:  "tendermint/PubKeyEd25519",
			Value: base64.StdEncoding.EncodeToString(publicKeyBytes),
		},
	}
	out, err := json.MarshalIndent(&kp, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Fprintf(outfile, "%s\n", out)
}

// GetAddress returns the node id in the standard Tendermint format
// func GetAddress(k *PrivKey) string {
// 	add := k.PubKey().Address()
// 	return strings.ToUpper(hex.EncodeToString(add))
// }

// KeyPairFromPrivKey outputs a keypair computed from the input private key
func KeyPairFromPrivKey(b64KeyString string, outfile io.Writer) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(b64KeyString)
	if err != nil {
		log.Fatal(err)
	}
	privateKey := ed25519.NewKeyFromSeed(privateKeyBytes)
	publicKey := privateKey.Public().(ed25519.PublicKey)

	fmt.Printf("Full (64 byte, as per rfc 8032):   %s\n", base64.StdEncoding.EncodeToString(privateKey))
	fmt.Printf("Partial private key (start point): %s\n", b64KeyString)
	fmt.Printf("Public Key:                        %s\n", base64.StdEncoding.EncodeToString(publicKey))
}
