package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("It is currently:", time.Now())
	fmt.Println("A number from 1-100", rand.Intn(100))

}

// type go run math-time-count.go for execution
