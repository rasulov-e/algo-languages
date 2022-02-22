package main

// import (
// 	"fmt"
// 	"log"
// 	"math/big"
// 	_ "net/http/pprof"
// 	"os"
// 	"time"
// )

// type Range struct {
// 	from int64
// 	to   int64
// }

// func main() {
// 	start := time.Now()
// 	filename, ranges, err := argumentParser(os.Args)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Time spent on getting arguments", time.Since(start))

// 	start = time.Now()
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Time spent on creating file", time.Since(start))

// 	start = time.Now()
// 	looper(ranges.from, ranges.to, file)
// 	fmt.Println("Time spent on computing", time.Since(start))
// }

// func looper(from, to int64, f *os.File) {
// 	if from >= to {
// 		return
// 	}
// 	var a []int64
// 	for i := from; i < to; i++ {
// 		if isPrimeNumber(i) {
// 			a = append(a, i)
// 		}

// 	}
// 	fmt.Fprint(f, "prime numbers for range of ", from, ":", to, a, "\n")
// }

// func isPrimeNumber(num int64) bool {
// 	return big.NewInt(num).ProbablyPrime(0)
// }

// func parseRange(r string) (from, to int64) {
// 	_, err := fmt.Sscanf(r, "%d:%d", &from, &to)
// 	if err != nil {
// 		log.Fatal("could not parse a range")
// 	}
// 	return
// }

// func argumentParser(args []string) (string, *Range, error) {
// 	var (
// 		filename    string
// 		ranging     string
// 		rangeResult *Range
// 	)

// 	for i, v := range args {

// 		switch v {
// 		case "--file":
// 			filename = args[i+1]
// 			args[i] = ""
// 			args[i+1] = ""
// 			continue
// 		case "--range":
// 			ranging = args[i+1]
// 			from, to := parseRange(ranging)
// 			rangeResult = &Range{
// 				from: from,
// 				to:   to,
// 			}
// 			args[i] = ""
// 			args[i+1] = ""
// 			continue
// 		default:
// 			continue
// 		}
// 	}

// 	return filename, rangeResult, nil
// }
