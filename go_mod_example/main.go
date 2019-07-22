package main

import (
	"fmt"
	"modules/base"

	"golang.org/x/crypto/acme"

	"aaa.com/work"
)

func main() {

	base.Test()
	work.SayHello()
	fmt.Println(acme.LetsEncryptURL)

}
