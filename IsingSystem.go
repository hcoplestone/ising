package main

import (
	"fmt"
)

// IsingSystem represents a closed system for running MC Ising Model simulations
type IsingSystem struct {
	gridSize        int
	verbose         bool
	randomGenerator *RandomGenerator
}

// NewIsingSystem initialises a new Insing Model simulation
func NewIsingSystem(gridSize int, seed int64, verbose bool) *IsingSystem {
	system := new(IsingSystem)

	system.gridSize = gridSize
	system.verbose = verbose

	system.randomGenerator = NewRandomGenerator(seed)

	if system.verbose {
		fmt.Printf("Creating system, grid size %d.\n", gridSize)
	}

	return system
}

// Update runs a MC sweep for every spin in the system
func (system *IsingSystem) Update() {
}
