package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Printf("Hello SentinelOne! Random UUID: %s\n", id)
}
