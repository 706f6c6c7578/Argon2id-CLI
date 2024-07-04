package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/argon2"
	"encoding/hex"
	"os"
)

func main() {
	// Define command-line flags
	passwordFlag := flag.String("p", "", "Password for Argon2id hashing.")
	saltFlag := flag.String("s", "", "Salt for Argon2id hashing.")
	lengthFlag := flag.Int("l", 32, "Key length for Argon2id hashing.")

	// Set up a usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Parse the command-line flags
	flag.Parse()

	// Check for required flags
	if *passwordFlag == "" || *saltFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Generate the key with Argon2id
	key := argon2.IDKey([]byte(*passwordFlag), []byte(*saltFlag), 1, 64*1024, 4, uint32(*lengthFlag))
	result := hex.EncodeToString(key)

	// Output the resulting key
	fmt.Println(result)
}

