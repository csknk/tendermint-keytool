package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/tendermint/crypto/bcrypt"
	"github.com/tendermint/tendermint/crypto/ed25519"
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

type PrivKey = ed25519.PrivKey
type MyPrivKey ed25519.PrivKey

func NewPrivKey() (*PrivKey, error) {
	k := ed25519.GenPrivKey()
	return &k, nil
}

// NewPrivKeyFromSeed outputs a private key from an input seed.
// FIXME add a proper salt.
func NewPrivKeyFromSeed(seed string) (*PrivKey, error) {
	salt := []byte("1234567890abcdef")
	secret, err := bcrypt.GenerateFromPassword(salt, []byte(seed), 0)
	if err != nil {
		log.Fatal(err)
	}
	k := ed25519.GenPrivKeyFromSecret(secret)
	return &k, nil
}

// OutputKeyPair outputs keypair to stdout in JSON format for Tendermint
func OutputKeyPair(key *PrivKey, outfile io.Writer) {
	pubKey := key.PubKey()
	outputPrivKey := KeyJSON{
		Type:  "tendermint/PubKeyEd25519",
		Value: base64.StdEncoding.EncodeToString(*key),
	}
	address := GetAddress(key)
	kp := KeyPair{
		Address: address,
		PrivKey: outputPrivKey,
		PubKey: KeyJSON{
			Type:  "tendermint/PubKeyEd25519",
			Value: base64.StdEncoding.EncodeToString(pubKey.Bytes()),
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
func GetAddress(k *PrivKey) string {
	add := k.PubKey().Address()
	return strings.ToUpper(hex.EncodeToString(add))
}

// KeyPairFromPrivKey outputs a keypair computed from the input private key
func KeyPairFromPrivKey(b64KeyString string, outfile io.Writer) {
	hexBytes, _ := base64.StdEncoding.DecodeString(b64KeyString)
	var privKey ed25519.PrivKey = hexBytes
	OutputKeyPair(&privKey, outfile)
}
