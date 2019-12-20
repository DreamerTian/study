package main

import (
	"crypto/sha256"
	"fmt"
	"study/publicChain/part48/BLC"
)

func main()  {
	bytes := []byte("http://liyuechun.org")

	hasher := sha256.New()
	hasher.Write(bytes)
	hash := hasher.Sum(nil)

	b58 := BLC.Base58Encode(hash)

	fmt.Printf("%x\n",b58)
	fmt.Printf("%s\n",b58)

	byteStr := BLC.Base58Decode(b58)

	fmt.Printf("%s\n",byteStr)
}
