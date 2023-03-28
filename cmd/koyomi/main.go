package main

import (
	"context"
	"log"
	"os"

	"github.com/nukokusa/koyomi"
)

var Version string

func init() {
	koyomi.Version = Version
}

func main() {
	ctx := context.Background()
	if err := koyomi.Run(ctx, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
