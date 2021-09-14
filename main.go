package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length  = flag.Int("l", 8, "length of the password to generate")
	caps    = flag.Bool("C", true, "wether to use caps in the password or not")
	special = flag.Bool("s", true, "wether to use special characters in the password or not")
	chars   = flag.String("c", "abcdefghijklmnopqrstuvwxyz", "own custom set of characters")
)

func init() {
	flag.Parse()

	if *chars == "abcdefghijklmnopqrstuvwxyz" {
		if *caps {
			*chars = *chars + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		}

		if *special {
			*chars = *chars + "!\"#$%'()*+,-./:;<=>?@[\\]^_`{|}~"
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	charRune := []rune(*chars)
	var pwd string
	for i := 1; i < *length; i++ {
		randomIndex := rand.Intn(len(charRune))
		pick := charRune[randomIndex]
		pwd = pwd + string(pick)
	}
	fmt.Println("Your password is:", pwd)
}
