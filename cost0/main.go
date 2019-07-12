//usr/bin/env go run "$0" "$@"; exit
package main

import (
	"fmt"

	"golang.org/x/crypto/blake2b"
)

func main() {
	var msg []byte

	msg = append(msg, byte(0))
	fmt.Println("vim-go", blake2b.Sum256(msg))
}
