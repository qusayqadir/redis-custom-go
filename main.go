// packages the entire together

package main

import (
	"log"

	"github.com/qusayqadir/redis-custom-go/internal/server"
)

func main() {
	log.Println("redis-custom-go: starting up ")
	if err := server.Run(":6379"); err != nil {
		log.Fatal(err)
	}
}
