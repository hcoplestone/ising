package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func runIsingSystem(seed int64, wg *sync.WaitGroup, systemID int) {
	defer wg.Done()

	isingSystem := NewIsingSystem(40, seed, true)

	for i := 0; i < 10; i++ {
		isingSystem.Update()
	}

	fmt.Println("System " + strconv.Itoa(systemID) + " finished!")
}

func main() {
	numberOfCores := runtime.NumCPU()
	fmt.Printf("Using " + strconv.Itoa(numberOfCores) + " cores...\n")
	runtime.GOMAXPROCS(numberOfCores)

	var wg sync.WaitGroup

	i := 0
	for i < 1 {
		wg.Add(1)
		fmt.Printf("Starting system %d\n", i)
		go runIsingSystem(int64(i), &wg, i)
		i++
	}
	fmt.Println("")

	wg.Wait()
	fmt.Println("\nAll systems complete!!!")
}
