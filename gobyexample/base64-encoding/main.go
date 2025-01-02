package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	data := "abc123!?$*&()'-=@~"

	sEnv := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnv)

	sDec, _ := b64.StdEncoding.DecodeString(sEnv)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
