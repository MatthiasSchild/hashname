package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
)

type hasherConstructor func() hash.Hash

var hasherConstructors = map[string]hasherConstructor{
	"sha1":   sha1.New,
	"sha256": sha256.New,
	"sha512": sha512.New,
	"md5":    md5.New,
}

func hashFile(filename string) (string, error) {
	hasherConst, ok := hasherConstructors[optionMethod]
	if !ok {
		fmt.Println("Unknown hashing method:", optionMethod)
		os.Exit(1)
	}
	hasher := hasherConst()

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Failed to close file", filename, ":", err)
		}
	}()

	_, err = io.Copy(hasher, file)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
