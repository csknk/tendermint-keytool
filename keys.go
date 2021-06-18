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

	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/libs/bytes"
)

const (
	// AddressSize is the size of a pubkey address.
	AddressSize = tmhash.TruncatedSize
)

// An address is a []byte, but hex-encoded even in JSON.
// []byte leaves us the option to change the address length.
// Use an alias so Unmarshal methods (with ptr receivers) are available too.
type Address = bytes.HexBytes

func AddressHash(bz []byte) Address {
	return Address(tmhash.SumTruncated(bz))
}

type PubKey interface {
	Address() Address
	Bytes() []byte
	VerifySignature(msg []byte, sig []byte) bool
	Equals(PubKey) bool
	Type() string
}

//var _ PubKey := (*PubKey)(nil)

type PrivKey interface {
	Bytes() []byte
	Sign(msg []byte) ([]byte, error)
	PubKey() PubKey
	Equals(PrivKey) bool
	Type() string
}

type Symmetric interface {
	Keygen() []byte
	Encrypt(plaintext []byte, secret []byte) (ciphertext []byte)
	Decrypt(ciphertext []byte, secret []byte) (plaintext []byte, err error)
}

type KeyPair struct {
	Address string  `json:"address"`
	PubKey  KeyJSON `json:"pub_key"`
	PrivKey KeyJSON `json:"priv_key"`
}

type KeyJSON struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewPrivKey() (*ed25519.PrivKey, error) {
	k := ed25519.GenPrivKey()
	return &k, nil
}
func OutputKeyPair(privKey *ed25519.PrivKey, outfile io.Writer) {
	pubKey := privKey.PubKey()
	outputPrivKey := KeyJSON{
		Type:  "tendermint/PubKeyEd25519",
		Value: base64.StdEncoding.EncodeToString(*privKey),
	}
	address := GetAddress(privKey)
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
	fmt.Fprintf(outfile, "%s", out)
}

// See if there is a concrete implementation of PubKey()
func GetAddress(k *ed25519.PrivKey) string {
	add := k.PubKey().Address()
	return strings.ToUpper(hex.EncodeToString(add))
}
