//usr/bin/env go run "$0" "$@"; exit
package main

import (
	"fmt"
	"time"

	"github.com/u6du/ex"
	"golang.org/x/crypto/blake2b"
)

func next(msg []byte) []byte {
	for i := range msg {
		t := msg[i]
		if t != 255 {
			msg[i] = t + 1
			for j := 0; j < i; j++ {
				msg[j] = 0
			}
			return msg
		}
	}
	return make([]byte, len(msg)+1)
}

func Begin0Count(msg []byte) (n uint) {

	for i := range msg {
		if msg[i] == 0 {
			n++
		} else {
			break
		}
	}

	return
}

func Begin0MoreThan(msg []byte, n uint) []byte {
	var salt []byte
	atLest := uint(3)
	begin := uint64(time.Now().UnixNano())
	count := uint(0)

	for {
		salt = next(salt)
		hasher, err := blake2b.New256(msg)

		ex.Panic(err)
		hash := hasher.Sum(salt)

		if Begin0Count(hash) >= atLest {
			count += 1
			cost := (uint64(time.Now().UnixNano()) - begin) / uint64(time.Millisecond) / uint64(count)
			fmt.Printf("%dms/hash %d %d salt %x hash %x\n", cost, count, len(msg), msg, hash)
			return salt
		}
	}
}

func main() {

	Begin0MoreThan([]byte("1gxxxzzcccabcde"), 2)
}
