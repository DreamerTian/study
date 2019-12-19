package main

import (
	"crypto/sha256"
	"fmt"
)

func main()  {
	hasher := sha256.New()

	hasher.Write([]byte("http://baidu.com"))

	b := hasher.Sum(nil)

	fmt.Println(b)
}
