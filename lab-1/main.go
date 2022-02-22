package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"sync"
	"time"
)

type Range struct {
	from int64
	to   int64
}

func main() {
	start := time.Now()
	filename, timeout, ranges, err := argumentParser(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Time spent on getting arguments", time.Since(start).Seconds())

	start = time.Now()
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Time spent on creating file", time.Since(start).Seconds())

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*time.Duration(timeout))

	var waiter sync.WaitGroup
	start = time.Now()
	for i := 0; i < len(ranges); i++ {
		waiter.Add(1)
		go looper(ctx, &waiter, ranges[i].from, ranges[i].to, file)
	}
	fmt.Println("Time spent on initializing all goroutines file", time.Since(start).Seconds())

	waiter.Wait()
}

func looper(ctx context.Context, w *sync.WaitGroup, from, to int64, f *os.File) {
	start := time.Now()
	if from >= to {
		return
	}
	var a []int64
	for i := from; i < to; i++ {
		if isPrimeNumber(i) {
			a = append(a, i)
		}
		select {
		case <-ctx.Done():
			log.Println("not enough time")
			fmt.Fprint(f, "prime numbers for range of ", from, ":", to, a, "\n")
			w.Done()
			return
		default:
			continue
		}
	}
	fmt.Fprint(f, "prime numbers for range of ", from, ":", to, a, "\n")
	fmt.Println("Time spent on computing for range of", from, to, "is", time.Since(start).Seconds())
	w.Done()
}

func isPrimeNumber(num int64) bool {
	return big.NewInt(num).ProbablyPrime(0)
}

func parseRange(r string) (from, to int64) {
	_, err := fmt.Sscanf(r, "%d:%d", &from, &to)
	if err != nil {
		log.Fatal("could not parse a range")
	}
	return
}

func argumentParser(args []string) (string, int, []*Range, error) {
	var (
		filename    string
		timeout     int
		ranging     string
		rangeResult []*Range
	)

	for i, v := range args {

		switch v {
		case "--file":
			filename = args[i+1]
			args[i] = ""
			args[i+1] = ""
			continue
		case "--range":
			ranging = args[i+1]
			from, to := parseRange(ranging)
			rangeResult = append(rangeResult, &Range{
				from: from,
				to:   to,
			})
			args[i] = ""
			args[i+1] = ""
			continue
		case "--timeout":
			// dont touch this part
			// it looks studpid but it works
			t, err := strconv.Atoi(args[i+1])
			timeout = t
			if err != nil {
				return "", 0, nil, err
			}
			args[i] = ""
			args[i+1] = ""
			continue
		default:
			continue
		}
	}

	return filename, timeout, rangeResult, nil
}
