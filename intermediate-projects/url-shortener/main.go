package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	shortenUrl()
}

func shortenUrl() {
	// Provides more flexibility, allowing to call Write multiple times to build the data incrementally before finalizing with Sum(nil).
	// hash := sha256.New()
	// hash.Write([]byte("test"))
	// shortHash := hash.Sum(nil)

	// It directly returns the SHA256 hash of the input data as a fixed-size array (specifically, [32]byte), without the need to create a hash. Hash object manually.
	shortHash := sha256.Sum256([]byte("http://example.com/some/long/path"))
	fmt.Println(hex.EncodeToString(shortHash[:])[:8])
}
