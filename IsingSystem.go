package main

import (
	"fmt"
	"os"
	"os/exec"
)

// IsingSystem represents a closed system for running MC Ising Model simulations
// gridSize: represents the size of the system
// verbose: flag toggling verbose output
// randomGenerator: random number generator for stochastic behaviour
// beta: the inverse temperature parameter
type IsingSystem struct {
	gridSize           int
	grid               [][]int
	verbose            bool
	randomGenerator    *RandomGenerator
	initialTemperature float64
	beta               float64
}

type Position struct {
	i int
	j int
}

// NewIsingSystem initialises a new Insing Model simulation
func NewIsingSystem(gridSize int, seed int64, verbose bool) *IsingSystem {
	system := new(IsingSystem)
	system.initialTemperature = 4.0

	system.gridSize = gridSize
	system.verbose = verbose
	system.randomGenerator = NewRandomGenerator(seed)

	system.grid = make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		system.grid[i] = make([]int, gridSize)
	}

	if system.verbose {
		fmt.Printf("Creating system, grid size %d.\n", gridSize)
	}

	return system
}

// Reset restores the system to it's initial state
func (system *IsingSystem) Reset() {
	// Set initial temperature
	system.SetTemperature(system.initialTemperature)

	// Reset grid
	for i, column := range system.grid {
		for j := range column {
			system.grid[i][j] = 0
		}
	}
}

// SetTemperature changes the temperature of the Ising system
func (system *IsingSystem) SetTemperature(temperature float64) {
	system.beta = 1 / temperature
}

// SetGrid assigns the value of grid[i][j] to value
func (system *IsingSystem) SetGrid(position *Position, value int) {
	system.grid[position.i][position.j] = value
}

// ReadGrid returns the value of the grid position (i ,j)
func (system *IsingSystem) ReadGrid(position *Position) int {
	return system.grid[position.i][position.j]
}

// DisplayGrid prints grid to stdout
func (system *IsingSystem) DisplayGrid() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	// fmt.Printf("\n")
	for _, column := range system.grid {
		for _, row := range column {
			if row == 1 {
				fmt.Printf("● ")
			} else {
				fmt.Printf("◌ ")
			}
		}
		fmt.Printf("\n")
	}
}

// Update runs a MC sweep for every spin in the system
func (system *IsingSystem) Update() {
	system.grid[10][10] = 0
	system.grid[10][15] = 1
}
