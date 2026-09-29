package main

import (
	"image/color"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func plotResults(
	workerCounts []int,
	cpuTimes []float64,
	ioTimes []float64,
) error {

	p := plot.New()

	p.Title.Text = "CPU vs I/O Worker Pool Performance"
	p.X.Label.Text = "Number of Workers"
	p.Y.Label.Text = "Execution Time (seconds)"

	cpuPoints := make(plotter.XYs, len(workerCounts))
	ioPoints := make(plotter.XYs, len(workerCounts))

	for i := range workerCounts {
		cpuPoints[i].X = float64(workerCounts[i])
		cpuPoints[i].Y = cpuTimes[i]

		ioPoints[i].X = float64(workerCounts[i])
		ioPoints[i].Y = ioTimes[i]
	}

	cpuLine, err := plotter.NewLine(cpuPoints)
	if err != nil {
		return err
	}

	ioLine, err := plotter.NewLine(ioPoints)
	if err != nil {
		return err
	}

	cpuLine.LineStyle.Width = vg.Points(2)
	ioLine.LineStyle.Width = vg.Points(2)
	cpuLine.LineStyle.Color = color.RGBA{R: 220, G: 50, B: 47, A: 255}
	ioLine.LineStyle.Color = color.RGBA{R: 38, G: 139, B: 210, A: 255}

	p.Add(cpuLine, ioLine)

	p.Legend.Add("CPU", cpuLine)
	p.Legend.Add("I/O", ioLine)

	return p.Save(
		10*vg.Inch,
		6*vg.Inch,
		"worker_performance.png",
	)
}
