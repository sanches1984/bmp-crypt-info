package main

import (
	"flag"
	"fmt"
	"github.com/sanches1984/bmp-crypt-info/converter"
	"github.com/sanches1984/bmp-crypt-info/converter/worker"
	"log"
	"os"
)

var (
	level       string
	source      string
	destination string
	phrase      string
	encrypt     bool
	decrypt     bool
)

func init() {
	flag.BoolVar(&encrypt, "encrypt", false, "Encrypt image")
	flag.BoolVar(&decrypt, "decrypt", false, "Decrypt image")
	flag.StringVar(&level, "level", "low", "Crypt level (low, normal, high)")
	flag.StringVar(&source, "src", "", "Source image")
	flag.StringVar(&destination, "dst", "result.bmp", "Destination image")
	flag.StringVar(&phrase, "phrase", "", "Secret phrase")
}

func main() {
	flag.Parse()

	cryptLevel := getCryptLevel()
	if !encrypt && !decrypt || encrypt && decrypt || cryptLevel == 0 {
		fmt.Println("bad command")
		return
	}
	if _, err := os.Stat(source); os.IsNotExist(err) {
		fmt.Println("source file not found")
		return
	}
	if encrypt && phrase == "" {
		fmt.Println("phrase not set")
		return
	}

	if encrypt {
		err := converter.New(cryptLevel).CryptFile(source, destination, "hello world")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Encrypted phrase [%s] to file: %s\n", phrase, destination)
	} else if decrypt {
		secret, err := converter.New(cryptLevel).DecryptFile(source)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Decrypted phrase [%s]\n", secret)
	}
}

func getCryptLevel() worker.CryptLevel {
	switch level {
	case "low":
		return worker.CryptLevelLow
	case "normal":
		return worker.CryptLevelMedium
	case "high":
		return worker.CryptLevelHigh
	}
	return 0
}
