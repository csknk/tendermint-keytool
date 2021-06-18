package main

import (
	"log"
	"os"
)

type Key struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func main() {
	//	add := pubKey.Address()
	//	fmt.Printf("Private key: %v\n", hex.EncodeToString(privKey))
	//	fmt.Printf("Private key: %v\n", base64.StdEncoding.EncodeToString(privKey))
	//	fmt.Printf("Public key: %v\n", hex.EncodeToString(pubKey.Bytes()))
	//	fmt.Printf("Public key: %v\n", base64.StdEncoding.EncodeToString(pubKey.Bytes()))
	//	fmt.Printf("Address: %v\n", add)
	//	kp := KeyPair{
	//		Address: address,
	//		PrivKey: privKey,
	//		PubKey: Key{
	//			Type:  "tendermint/PubKeyEd25519",
	//			Value: base64.StdEncoding.EncodeToString(pubKey.Bytes()),
	//		},
	//	}
	//}
	//	out, err := json.MarshalIndent(&kp, "", "  ")
	//	if err != nil {
	//		log.Println(err)
	//		os.Exit(1)
	//	}
	//
	//	fmt.Println(string(out))

	privKey, err := NewPrivKey()
	if err != nil {
		log.Fatal(err)
	}

	//	fmt.Println(privKey.Bytes())
	OutputKeyPair(privKey, os.Stdout)

}
