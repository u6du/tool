//usr/bin/env go run "$0" "$@"; exit
package main

import (
	"fmt"
	"math/bits"
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

func Begin0Count(msg []byte) (n int) {

	for i := range msg {
		t := bits.OnesCount8(uint8(0) ^ msg[i])
		n += t
		if t != 8 {
			break
		}
	}

	return
}

func Begin0MoreThan(msg []byte, atLest int) []byte {
	var salt []byte
	begin := uint64(time.Now().UnixNano())
	count := uint(0)

	for {
		salt = next(salt)
		hasher, err := blake2b.New256(nil)
		hasher.Write(msg)
		hasher.Write(salt)
		ex.Panic(err)
		hash := hasher.Sum(nil)

		if Begin0Count(hash) >= atLest {
			count += 1
			cost := (uint64(time.Now().UnixNano()) - begin) / uint64(time.Millisecond) / uint64(count)
			fmt.Printf("%dms/hash %d %d salt %x hash %x\n", cost, count, len(msg), msg, hash)
			return salt
		}
	}
}

func main() {
	var msg []byte
	for {
		msg = next(msg)
		salt := Begin0MoreThan([]byte(msg), 22)
		fmt.Printf("%x\n\n", blake2b.Sum256(append(msg, salt...)))
	}
}
