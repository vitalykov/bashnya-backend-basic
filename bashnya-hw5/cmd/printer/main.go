package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"time"

	"github.com/vitalykov/bashnya-backend-basic/bashnya-hw5/printer"
)

func generateValue(rnd *rand.Rand) any {
	switch rnd.Intn(5) {
	case 0:
		return "Hello Bashnya"
	case 1:
		return rand.Intn(10000000)
	case 2:
		return rand.Float64()
	case 3:
		return byte(rand.Intn(128))
	case 4:
		return rune(rand.Intn(65536))
	}
	return 5
}

const timeFormat = "Jan 2, 2006 15:04:05.999"

func main() {
	workersCount := runtime.GOMAXPROCS(-1)
	if len(os.Args) == 2 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("%s: not an number of workers\n", os.Args[1])
			os.Exit(1)
		}
		workersCount = n
	}
	fmt.Println(workersCount)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	in := make(chan any)
	p := printer.NewPrinter(in, os.Stdout, workersCount)
	p.Run()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		select {
		case <-signals:
			tSigint := time.Now()
			p.Shutdown()
			fmt.Printf("SIGINT received at %s.\n", tSigint.Format(timeFormat))
			fmt.Printf("Printer shutdown at %s.\n", time.Now().Format(timeFormat))
			return
		default:
			in <- generateValue(rnd)
		}
	}
}
