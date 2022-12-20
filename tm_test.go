package main

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prysmaticlabs/prysm/crypto/bls/blst"
)

// response from tm api
const validators = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "block_height": "10",
        "validators": [
            {
                "address": "1E4144FF0CBDEEFAA23DC6C5FB682B177E423983",
                "pub_key": {
                    "type": "tendermint/PubKeyEd25519",
                    "value": "xY5CnBC0SWCQkOS9wwZQx1aQ2BVzKFgH+FKlq3vMoEQ="
                },
                "voting_power": "100",
                "proposer_priority": "0",
                "relayer_bls_key": "len/DF6Nabeug/HEryieb5eJN0Fd1T6Vu9Jzbv77bfMmsH+6Atdwc6qGlns/PK6u",
                "relayer_address": "zdOTcj8a+B+qPzyHtR2rcrbGgVQ="
            }
        ],
        "count": "1",
        "total": "1"
    }
}
`

func TestDecodeRelayerAddress(t *testing.T) {
	decoded, err := base64.StdEncoding.DecodeString("zdOTcj8a+B+qPzyHtR2rcrbGgVQ=")
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Printf("decoded: %x \n", decoded)
}

func TestDecodeRelayerBlsKey(t *testing.T) {
	decoded, err := base64.StdEncoding.DecodeString("len/DF6Nabeug/HEryieb5eJN0Fd1T6Vu9Jzbv77bfMmsH+6Atdwc6qGlns/PK6u")
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Printf("decoded: %x \n", decoded)
}

func TestGenerateBlsKey(t *testing.T) {
	privKey, _ := blst.RandKey()
	pubKey := privKey.PublicKey().Marshal()
	fmt.Printf("privKey: %x \n", privKey.Marshal())
	fmt.Printf("pubkey: %x \n", pubKey)
}

func TestRecoverBlsKey(t *testing.T) {
	// private key: 38ebe1ea64da024ac6ac7a9b5f484293c6f391b65837c9b5a8b27b47bef96f42
	// public key: 95e9ff0c5e8d69b7ae83f1c4af289e6f978937415dd53e95bbd2736efefb6df326b07fba02d77073aa86967b3f3caeae
	bs := common.Hex2Bytes("38ebe1ea64da024ac6ac7a9b5f484293c6f391b65837c9b5a8b27b47bef96f42")
	secretKey, err := blst.SecretKeyFromBytes(bs)
	if err != nil {
		panic(err)
	}
	pubKey := secretKey.PublicKey()
	fmt.Printf("pubkey: %x \n", pubKey.Marshal())
}
