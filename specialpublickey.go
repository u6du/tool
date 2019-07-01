//usr/bin/env go run "$0" "$@"; exit
package main

import (
	"crypto/rand"
	"io/ioutil"
	"runtime"
	"strings"

	"github.com/u6du/ex"
	"github.com/u6du/go-rfc1924/base85"
	"golang.org/x/crypto/ed25519"
)

func run(ch chan string) {
	count := 0
NEXT:
	for {
		if count%10000 == 0 {
			println("> ", count)
		}
		count++

		_, private, err := ed25519.GenerateKey(rand.Reader)
		ex.Panic(err)
		public := base85.EncodeToString(private.Public().(ed25519.PublicKey))

		if strings.Index(strings.ToLower(public), "6du") < 0 {
			continue
		}

		for _, c := range "<>&`$%=-|@{}()*#;_!^?~+" {
			if strings.Index(public, string(c)) >= 0 {
				continue NEXT
			}
		}

		filepath := "6du.private"
		ex.Panic(ioutil.WriteFile(filepath, private.Seed(), 0600))

		ch <- public
		break
	}
}

func main() {
	ch := make(chan string)
	num := runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < num; i++ {
		go run(ch)
	}
	public := <-ch
	println(public)
}
