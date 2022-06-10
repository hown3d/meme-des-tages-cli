package main

import (
	"github.com/hown3d/meme-des-tages-cli/wasm"
)

func main() {
	c := make(chan struct{})
	wasm.Register(c)
	<-c
}
