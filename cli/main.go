package main

import (
	"flag"
	"fmt"
	"log"

	helpers "github.com/polisko/encdec"
)

func main() {
	cmd := flag.String("c", "e", "e for encryption, d for decryption")
	pass := flag.String("p", "", "Passphrase")

	flag.Parse()

	switch *cmd {
	case "e", "enc":
		fmt.Println("Encryption")
		cipherText := helpers.Encrypt([]byte(flag.Arg(0)), *pass)
		fmt.Printf("%x\n", cipherText)
	case "d", "dec":
		fmt.Println("Decryption")
		plain, err := helpers.DecryptHexString(flag.Arg(0), *pass)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%s\n", plain)
	}
}
