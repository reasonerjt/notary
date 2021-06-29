package main

import (
	"fmt"
	"os"

	jose "github.com/dvsekhvalnov/jose2go"
)

// Constants
const (
	EncryptionAlg = jose.A256GCM
	KeywrapAlg    = jose.PBES2_HS256_A128KW
)

var goversion string

func main() {
	args:= os.Args
	key := []byte{48,119,2,1,1,4,32,123,130,10,169,120,240,111,225,215,228,25,255,27,111,89,211,45,214,47,194,48,158,208,200,179,199,247,96,175,20,103,87,160,10,6,8,42,134,72,206,61,3,1,7,161,68,3,66,0,4,37,169,250,23,228,226,195,27,210,8,95,61,17,159,92,97,95,65,6,2,183,37,225,109,197,215,112,69,239,110,188,7,7,58,166,216,184,254,116,24,27,211,92,109,21,90,253,205,244,121,45,211,15,220,91,27,102,36,57,213,182,119,87,166}
	passphrase := "tcITiqT6"
	fmt.Println("go version:", goversion)
	if len(args)  == 1 {
		fmt.Println("Need to specify encrypt or decrypt.")
		return
	}
	if args[1] == "encrypt" {
		encrypted, err := jose.Encrypt(string(key), KeywrapAlg, EncryptionAlg, passphrase)
		if err != nil {
			fmt.Println("error in encrypt: ", err)
			return
		}
		fmt.Println("encrypted: ", encrypted)
	} else if args[1] == "decrypt" {
		if len(args) != 3 {
			fmt.Println("Invalid args for decode, len: ", len(args))
			return
		}
		ds, _, err := jose.Decode(args[2], passphrase)
		if err != nil {
			fmt.Println("error in decrypt: ", err)
			return
		}
		decryptedKey := []byte(ds)
		fmt.Printf("Decrypted private key: %v\n", decryptedKey)
		if comp(key, decryptedKey) {
			fmt.Println("SUCCESS: Decrypted key identical to input key")
			return
		}
		fmt.Println("FAIL: Decrypted key NOT identical to input key")

	} else {
		fmt.Println("Invalid subcommand:", args[0])
	}
}

func comp(exp, in []byte) bool {
	if len(exp) != len(in) {
		fmt.Printf("Difference in length, expected: %d, actual: %d\n", len(exp), len(in))
		return false
	}
	for i:=0; i< len(exp); i++ {
		if exp[i] != in[i] {
			fmt.Printf("Index: %d, expect %v, acutal %v\n", i, exp[i], in[i])
			return false
		}
	}
	return true
}

