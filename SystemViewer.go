package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

// SystemViewer is an interface for visualising the Ising system
type SystemViewer struct {
	system     *IsingSystem
	gridSize   int
	spinSize   int
	fakeOrigin int
}

// NewSystemViewer creates a viewing interface for a given ising model system
func NewSystemViewer(isingSystem *IsingSystem, gridSize int, spinSize int) *SystemViewer {
	systemViewer := new(SystemViewer)
	systemViewer.system = isingSystem
	systemViewer.gridSize = gridSize
	systemViewer.spinSize = spinSize
	systemViewer.fakeOrigin = 10
	return systemViewer
}

// Run is the main window loop for our system viewer
func (viewer *SystemViewer) Run() {
	dimension := float64(viewer.gridSize*viewer.spinSize + viewer.fakeOrigin*2)

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, dimension, dimension),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		// win.Clear(colornames.Black)

		imd := imdraw.New(nil)

		for colIndex, column := range viewer.system.grid {
			for rowIndex, row := range column {
				bottomLeftX := float64(rowIndex*viewer.spinSize + viewer.fakeOrigin)
				topRightX := float64((rowIndex+1)*viewer.spinSize + viewer.fakeOrigin)

				bottomLeftY := float64(colIndex*viewer.spinSize + viewer.fakeOrigin)
				topRightY := float64((colIndex+1)*viewer.spinSize + viewer.fakeOrigin)

				if row == 1 {
					imd.Color = pixel.RGB(0, 0, 1)
				} else {
					imd.Color = pixel.RGB(0, 1, 0)
				}

				imd.Push(pixel.V(bottomLeftX, bottomLeftY))
				imd.Push(pixel.V(topRightX, topRightY))
				imd.Rectangle(0)
			}
		}

		imd.Draw(win)
		win.Update()
	}
}
