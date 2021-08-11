package main

import (
	"os"
	"os/exec"

	"github.com/goccy/go-graphviz"
)

var outputTypes = map[string]graphviz.Format{
	"SVG":  graphviz.SVG,
	"PNG":  graphviz.PNG,
	"JPG":  graphviz.JPG,
	"XDOT": graphviz.XDOT,
}

var outputTypeList = []string{"PNG", "SVG", "JPG", "XDOT"}

var layouts = map[string]graphviz.Layout{
	"Dot":       graphviz.DOT,
	"FDP":       graphviz.FDP,
	"Neato":     graphviz.NEATO,
	"Circo":     graphviz.CIRCO,
	"Twopi":     graphviz.TWOPI,
	"Osage":     graphviz.OSAGE,
	"Patchwork": graphviz.PATCHWORK,
	"SFDP":      graphviz.SFDP,
}

var layoutList = []string{"Dot", "FDP", "Neato", "Circo", "Twopi", "Osage", "Patchwork", "SFDP"}

func Render(input string, output string, format string, layout string, system bool) error {
	if !system {
		v := graphviz.New()
		g, err := graphviz.ParseFile(input)
		if err != nil {
			return err
		}
		v.SetLayout(layouts[layout])
		return v.RenderFilename(g, outputTypes[format], output)
	}

	cmd := exec.Command(string(layouts[layout]), "-T"+string(outputTypes[format]), input)
	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outFile.Close()

	cmd.Stdout = outFile
	return cmd.Run()
}
