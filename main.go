package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// func runIsingSystem(initialBeta float64, numberOfSweeps int, seed int64, wg *sync.WaitGroup, subSystemID int) {
// 	defer wg.Done()

// 	initialTemperature := 1 / initialBeta
// 	system := NewIsingSystem(40, initialTemperature, seed, false)

// 	filenameComponents := []string{"results/section1final/beta-", strconv.Itoa(int(initialBeta * 100)), "-system", strconv.Itoa(subSystemID), ".csv"}
// 	csvFilename := strings.Join(filenameComponents, "")

// 	f, err := os.OpenFile(csvFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if _, err := f.Write([]byte(fmt.Sprintf("sweep, beta, subsystem_id, M, E/J\n"))); err != nil {
// 		log.Fatal(err)
// 	}

// 	for i := 0; i <= numberOfSweeps; i++ {
// 		if _, err := f.Write([]byte(fmt.Sprintf("%d, %f, %d, %f, %f\n", i, initialBeta, subSystemID, system.ComputeMagnetisation(), system.ComputeDimensionlessSystemEnergy()))); err != nil {
// 			log.Fatal(err)
// 		}
// 		system.Update()
// 	}

// 	if err := f.Close(); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("System finished - sub system %d - beta = %f\n", subSystemID, initialBeta)
// }

func runIsingSystem(initialBeta float64, numberOfSweeps int, seed int64, wg *sync.WaitGroup, subSystemID int) {
	defer wg.Done()

	initialTemperature := 1 / initialBeta
	system := NewIsingSystem(40, initialTemperature, seed, false)

	filenameComponents := []string{"results/section2final3/beta-", strconv.Itoa(int(initialBeta * 1000)), "-system", strconv.Itoa(subSystemID), ".csv"}
	csvFilename := strings.Join(filenameComponents, "")

	f, err := os.OpenFile(csvFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write([]byte(fmt.Sprintf("sweep, beta, subsystem_id, M, E/J\n"))); err != nil {
		log.Fatal(err)
	}

	for i := 0; i <= numberOfSweeps; i++ {
		if _, err := f.Write([]byte(fmt.Sprintf("%d, %f, %d, %f, %f\n", i, initialBeta, subSystemID, system.ComputeMagnetisation(), system.ComputeDimensionlessSystemEnergy()))); err != nil {
			log.Fatal(err)
		}
		system.Update()
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("System finished - sub system %d - beta = %f\n", subSystemID, initialBeta)
}

func main() {
	ensembleCount := 1

	// numberOfSweeps := 1000000
	numberOfSweeps := 200000
	betaStep := 0.001
	betaLowerLimit := 0.25
	betaUpperLimit := 1.0 + betaStep

	numberOfCores := runtime.NumCPU()
	fmt.Printf("Using " + strconv.Itoa(numberOfCores) + " cores...\n")
	runtime.GOMAXPROCS(numberOfCores)

	var wg sync.WaitGroup

	for ensembleSubSystem := 0; ensembleSubSystem < ensembleCount; ensembleSubSystem++ {
		for beta := betaLowerLimit; beta <= betaUpperLimit; beta = beta + betaStep {
			wg.Add(1)
			fmt.Printf("Starting system - sub system %d - beta = %f\n", ensembleSubSystem, beta)
			go runIsingSystem(beta, numberOfSweeps, int64(ensembleSubSystem), &wg, ensembleSubSystem)
		}
	}

	wg.Wait()
	fmt.Println("\nAll systems complete!!!")
}

// func main() {
// 	ensembleCount := 10000

// 	numberOfSweeps := 50
// 	betaStep := 0.1
// 	betaLowerLimit := 0.2
// 	betaUpperLimit := 0.7 + betaStep

// 	numberOfCores := runtime.NumCPU()
// 	fmt.Printf("Using " + strconv.Itoa(numberOfCores) + " cores...\n")
// 	runtime.GOMAXPROCS(numberOfCores)

// 	var wg sync.WaitGroup

// 	for ensembleSubSystem := 0; ensembleSubSystem < ensembleCount; ensembleSubSystem++ {
// 		for beta := betaLowerLimit; beta <= betaUpperLimit; beta = beta + betaStep {
// 			wg.Add(1)
// 			fmt.Printf("Starting system - sub system %d - beta = %f\n", ensembleSubSystem, beta)
// 			go runIsingSystem(beta, numberOfSweeps, int64(ensembleSubSystem), &wg, ensembleSubSystem)
// 		}
// 	}

// 	wg.Wait()
// 	fmt.Println("\nAll systems complete!!!")
// }
