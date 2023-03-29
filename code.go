package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Read the input file
	filename := "transactions.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the transactions and hash them
	var hashes [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tx := scanner.Text()
		hash, err := hex.DecodeString(tx)
		if err != nil {
			log.Fatal(err)
		}
		hashes = append(hashes, hash)
	}

	// Compute the Merkle root
	root := computeMerkleRoot(hashes)
	fmt.Printf("Merkle root: %x\n", root)
}

// computeMerkleRoot computes the Merkle root from a list of transaction hashes
func computeMerkleRoot(hashes [][]byte) []byte {
	// Handle special cases
	if len(hashes) == 0 {
		return nil
	}
	if len(hashes) == 1 {
		return hashes[0]
	}

	// Compute the parent hashes iteratively
	for len(hashes) > 1 {
		if len(hashes)%2 != 0 {
			hashes = append(hashes, hashes[len(hashes)-1])
		}
		var parentHashes [][]byte
		for i := 0; i < len(hashes); i += 2 {
			concat := append(hashes[i], hashes[i+1]...)
			hash := sha256.Sum256(concat)
			parentHashes = append(parentHashes, hash[:])
		}
		hashes = parentHashes
	}
	return hashes[0]
}
