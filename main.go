package main

import (
	"fmt"
	"os"
)

func main() {
	bytecount := 32
	entropysrc := "/dev/urandom"

	buf := make([]byte, bytecount)
	entropyfile, err := os.Open(entropysrc)
	if err != nil {
		panic(err)
	}
	defer entropyfile.Close()

	count, err := entropyfile.Read(buf)

	if err != nil {
		panic(err)
	}

	if count != bytecount {
		panic("unable to read sufficient entropy")
	}

	for i := 0; i < len(buf); i++ {
		fmt.Printf("%02x", buf[i])
	}
}
