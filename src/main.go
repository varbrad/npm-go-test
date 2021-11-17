package main

import (
	"log"
	"math/rand"
	"time"
)

func GeneratePortNumber() int {
	rand.Seed(time.Now().UnixNano())

	min := 1000
	max := 99999

	port := rand.Intn(max-min+1) + min

	return port
}

func main() {
	log.Println("Random Port:", GeneratePortNumber())
}
