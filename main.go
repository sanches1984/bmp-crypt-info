package main

import (
	"fmt"
	"github.com/sanches1984/bmp-crypt-info/converter"
	"github.com/sanches1984/bmp-crypt-info/converter/worker"
	"log"
)

func main() {
	conv := converter.New(worker.CryptLevelHigh)
	err := conv.CryptFile("test.bmp", "result.bmp", "hello world")
	if err != nil {
		log.Fatal(err)
	}

	secret, err := conv.DecryptFile("result.bmp")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("secret is:", secret)
}
