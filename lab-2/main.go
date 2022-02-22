package main

import (
	"errors"
	"log"
	"math/rand"
	"os"
	"time"
)

var logger log.Logger

func init() {
	f, err := os.Create("logging.txt")
	if err != nil {
		panic(err)
	}
	logger.SetOutput(f)
}

func main() {
	// all the main functionality in tests
}

func randomSlice(length int) ([]int, error) {
	if length <= 0 {
		return nil, errors.New("length can't be less or equal to 0")
	}
	rand.Seed(time.Now().Unix())
	var a []int = make([]int, length)

	for i := 0; i < length; i++ {
		a[i] = rand.Intn(length)
	}
	return a, nil
}
