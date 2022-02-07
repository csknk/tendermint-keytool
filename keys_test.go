package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

const keyPairFromPrivKeyCorrect = `{
  "address": "876DA0A61304879B2CDC506C877142F2BBE3E8C5",
  "pub_key": {
    "type": "tendermint/PubKeyEd25519",
    "value": "1Ymtk9rXFRZScNjTdXDXQ8Ef3p1zLO7YaywKtrOHTeM="
  },
  "priv_key": {
    "type": "tendermint/PubKeyEd25519",
    "value": "ZXEwxDK0xbzboZ0Rv/twTM9tanCbqW7Gug8QQ8w8hdPVia2T2tcVFlJw2NN1cNdDwR/enXMs7thrLAq2s4dN4w=="
  }
}
`

func TestKeyPairFromPrivKey(t *testing.T) {
	input := "ZXEwxDK0xbzboZ0Rv/twTM9tanCbqW7Gug8QQ8w8hdPVia2T2tcVFlJw2NN1cNdDwR/enXMs7thrLAq2s4dN4w=="
	var b bytes.Buffer
	KeyPairFromPrivKey(input, &b)
	assert.Equal(t, b.String(), keyPairFromPrivKeyCorrect)
}

func TestGoTest(t *testing.T) {
	fmt.Println("hello")

}
