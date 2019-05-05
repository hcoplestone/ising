package main

import (
	"fmt"
	"math"
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

// Grid direction definitions
const (
	DirectionRight int = 0
	DirectionDown  int = 1
	DirectionLeft  int = 2
	DirectionUp    int = 3
)

// Position is an arbitrary representation for the a grid location that allows us to change the underying grid implementation
// (coordinate origin, data structures, etc) but still have a consistent representation for setting and retrieving values in the grid
type Position struct {
	i int
	j int
}

// NewIsingSystem initialises a new Insing Model simulation
func NewIsingSystem(gridSize int, initialTemperature float64, seed int64, verbose bool) *IsingSystem {
	system := new(IsingSystem)
	system.initialTemperature = initialTemperature

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

	system.Reset()

	return system
}

// Reset restores the system to it's initial state
func (system *IsingSystem) Reset() {
	// Set initial temperature
	system.SetTemperature(system.initialTemperature)

	// Reset all spins to down (-1)
	for i, column := range system.grid {
		for j := range column {
			system.grid[i][j] = -1
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
	// fmt.Printf("beta = %e  M = %e \n", system.beta, system.ComputeMagnetisation())
	fmt.Printf("beta = %f  M = %f  E/J = %f \n", system.beta, system.ComputeMagnetisation(), system.ComputeDimensionlessSystemEnergy())
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

// DeterminePositionOfNeighbouringCell determines the grid coordinate position
// of a neighbouring cell in a given direction, applying periodic boundary conditions.
func (system *IsingSystem) DeterminePositionOfNeighbouringCell(initialPosition *Position, neighbourDirection int) Position {
	neighbourPosition := Position{initialPosition.i, initialPosition.j}

	if neighbourDirection == DirectionRight {
		neighbourPosition.i = (neighbourPosition.i + 1) % system.gridSize
	}

	if neighbourDirection == DirectionLeft {
		neighbourPosition.i = (neighbourPosition.i - 1 + system.gridSize) % system.gridSize
	}

	if neighbourDirection == DirectionUp {
		neighbourPosition.j = (neighbourPosition.j + 1) % system.gridSize
	}

	if neighbourDirection == DirectionDown {
		neighbourPosition.j = (neighbourPosition.j - 1 + system.gridSize) % system.gridSize
	}

	return neighbourPosition
}

// ComputeLocalFieldDividedByTemperature returns the local magnetic field for the spin at the specified position divided by the temperature
// i.e. returns (1/kT) * h_i === (1/T_0*J) * h_i
func (system *IsingSystem) ComputeLocalFieldDividedByTemperature(position *Position) float64 {
	sumOfNearestNeighbourSpins := float64(0.0)
	for direction := 0; direction < 4; direction++ {
		nearestNeighbourPosition := system.DeterminePositionOfNeighbouringCell(position, direction)
		sumOfNearestNeighbourSpins += float64(system.ReadGrid(&nearestNeighbourPosition))
	}
	return sumOfNearestNeighbourSpins * system.beta
}

// ComputeDimensionlessEnergyOfGridLocation computes (E_i)/J For a given grid location, i
// Dimensionless energy = -S_i * [sum over j nearest neighbours](s_j)
func (system *IsingSystem) ComputeDimensionlessEnergyOfGridLocation(position *Position) float64 {
	sumOfNearestNeighbourSpins := 0
	for direction := 0; direction < 4; direction++ {
		nearestNeighbourPosition := system.DeterminePositionOfNeighbouringCell(position, direction)
		sumOfNearestNeighbourSpins += system.ReadGrid(&nearestNeighbourPosition)
	}
	spinAtGridLocation := system.ReadGrid(position)
	return float64(-1 * spinAtGridLocation * sumOfNearestNeighbourSpins)
}

// ComputeDimensionlessSystemEnergy computes the dimensionless energy of the system (E/J)
func (system *IsingSystem) ComputeDimensionlessSystemEnergy() float64 {
	sumOfIndividualGridPointEnergies := 0.0
	for i := 0; i < system.gridSize; i++ {
		for j := 0; j < system.gridSize; j++ {
			sumOfIndividualGridPointEnergies += system.ComputeDimensionlessEnergyOfGridLocation(&Position{i, j})
		}
	}
	systemEnergy := 0.5 * sumOfIndividualGridPointEnergies
	return systemEnergy
}

// ComputeMagnetisation computes the magnetiation (per spin)
func (system *IsingSystem) ComputeMagnetisation() float64 {
	sumOfSpins := 0
	for _, column := range system.grid {
		for _, row := range column {
			sumOfSpins += row
		}
	}
	magnetisation := float64(sumOfSpins) / float64(system.gridSize*system.gridSize)
	return magnetisation
}

// FlipSpin flips the spin at grid position (i, j)
func (system *IsingSystem) FlipSpin(position *Position) {
	// if system.verbose {
	// fmt.Printf("Flipping spin at (%d, %d)", position.i, position.j)
	// }
	system.grid[position.i][position.j] = -1 * system.grid[position.i][position.j]
	// if system.ReadGrid(position) == 1 {
	// system.SetGrid(position, 0)
	// } else {
	// system.SetGrid(position, 1)
	// }
}

// AttemptSpinFlip attempts to flip a spin. Flip is accepted/rejected by Metropolis rule
func (system *IsingSystem) AttemptSpinFlip() {
	i := system.randomGenerator.RandomInt(system.gridSize)
	j := system.randomGenerator.RandomInt(system.gridSize)
	position := Position{i, j}

	hLocal := system.ComputeLocalFieldDividedByTemperature(&position)
	dE := 2.0 * hLocal * float64(system.ReadGrid(&position))

	if dE < 0 {
		system.FlipSpin(&position)
	} else {
		sampleProbability := system.randomGenerator.rand.Float64()
		if sampleProbability < math.Exp(-1*dE) {
			system.FlipSpin(&position)
		}
	}
}

// MCSweep attempts N spin flips, where N is the number of spins in the system
func (system *IsingSystem) MCSweep() {
	for i := 0; i < system.gridSize*system.gridSize; i++ {
		system.AttemptSpinFlip()
	}
}

// Update runs an MC sweep
func (system *IsingSystem) Update() {
	system.MCSweep()
}
