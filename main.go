package main

import (
	"github.com/ipfs/ipfs-cluster/config"
	"os"
)

func main() {
	id, err := config.NewIdentity()
	if err != nil {
		panic(err)
	}
	raw, err := id.ToJSON()
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(raw)
}
