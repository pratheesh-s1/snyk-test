package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Printf("Hello SentinelOne! Random UUID: %s\n", id)
	fmt.Println("Time : ", time.Now())
}
