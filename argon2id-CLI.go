package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/argon2"
	"encoding/hex"
	"os"
	"io/ioutil"
	"path/filepath"
)

func main() {
	// Define command-line flags
	passwordFlag := flag.String("p", "", "Password for Argon2id hashing.")
	saltFlag := flag.String("s", "", "Salt for Argon2id hashing.")
	lengthFlag := flag.Int("l", 32, "Key length for Argon2id hashing.")
	fileFlag := flag.String("w", "", "File to write the output.")

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
	if *fileFlag != "" {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Failed to get current directory: %v\n", err)
			return
		}
		err = ioutil.WriteFile(filepath.Join(dir, *fileFlag), []byte(result), 0600)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println(result)
	}
}

