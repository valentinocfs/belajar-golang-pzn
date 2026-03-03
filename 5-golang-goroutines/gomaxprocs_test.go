package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Threads", totalThreads)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines", totalGoroutines)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	// Mengubah jumlah thread yang digunakan oleh Go
	runtime.GOMAXPROCS(20)
	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Threads", totalThreads)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines", totalGoroutines)

	group.Wait()
}
