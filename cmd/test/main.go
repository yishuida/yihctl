package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	token := "ESyF6HdLzZmT_6xZTLmQ"
	dbKeyBase := "bf2e47b68d6cafaef1d767e628b619365becf27571e10f196f98dc85e7771042b9203199d39aff91fcb6837c8ed83f2a912b278da50999bb11a2fbc0fba52964"
	str := fmt.Sprintf("%s%s", token, dbKeyBase[:32])

	hash := sha256.New()
	hash.Write([]byte(str))

	// to lowercase hexits
	hex.EncodeToString(hash.Sum(nil))

	// to base64
	fmt.Println(base64.StdEncoding.EncodeToString(hash.Sum(nil)))
}
