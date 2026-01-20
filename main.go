package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	samplemodule "github.com/pratheesh-s1/snyk-test/sample-module"
)

func main() {
	id := uuid.New()
	fmt.Printf("Hello SentinelOne! Random UUID: %s\n", id)
	fmt.Println("Time : ", time.Now())

	samplemodule.Sample()
}
