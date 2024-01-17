package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ranon-rat/secretos24/src/router"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println(router.Setup())
}
