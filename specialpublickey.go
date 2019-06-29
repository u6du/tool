//usr/bin/env go run "$0" "$@"; exit
package main

import (
	"crypto/rand"
	"io/ioutil"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/u6du/ex"
	"github.com/u6du/go-rfc1924/base85"
	"golang.org/x/crypto/ed25519"
)

func main() {
	count := 0
NEXT:
	for {
		_, private, err := ed25519.GenerateKey(rand.Reader)
		ex.Panic(err)
		count++
		if count%100000 == 0 {
			log.Info().Int("count", count).Msg("")
		}
		public := base85.EncodeToString(private.Public().(ed25519.PublicKey))

		if strings.Index(strings.ToLower(public), "6du") != 0 {
			continue
		}
		println("> ", public)
		for _, c := range "<>&`$%=-|@{}()*#;_!^?~+" {
			if strings.Index(public, string(c)) >= 0 {
				continue NEXT
			}
		}

		println(public)

		filepath := "6du.private"
		ex.Panic(ioutil.WriteFile(filepath, private.Seed(), 0600))
		break
	}
}
