package main

import (
	"fmt"
	"study/publicChain/part45/BLC"
)

func main()  {
	bytes := []byte("http://liyuechun.org")

	b58 := BLC.Base58Encode(bytes)

	fmt.Printf("%x\n",b58)
	fmt.Printf("%s\n",b58)

	byteStr := BLC.Base58Decode(b58)

	fmt.Printf("%s\n",byteStr)
}
