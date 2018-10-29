package renderer

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

//PNGRenderer renders a graph as a png file
type PNGRenderer struct {
	HomeDir    string
	Prefix     string
	OutputFile string
	executor   DotExecutor
}

//DotExecutor creates a dot graph
type DotExecutor interface {
	Render(dotfile, outfile string) error
}

type dot struct{}

func (e *dot) Render(dotfile, outfile string) error {
	_, err := exec.Command("dot", "-Tpng", dotfile, "-o", outfile).Output()
	return err
}

//Render renders graph content
func (r PNGRenderer) Render(graphContent string) {
	outFile := r.getOutputFile()
	dotFile := outFile + ".dot"

	err := ioutil.WriteFile(dotFile, []byte(graphContent), os.ModePerm)
	if err != nil {
		log.Fatal("Error writing dot file.", err.Error())
	}

	if r.executor == nil {
		r.executor = &dot{}
	}

	err = (r.executor).Render(dotFile, outFile)
	if err != nil {
		log.Fatal("Error creating png.", err.Error())
	}
}

func (r PNGRenderer) getOutputFile() string {
	outFile := r.OutputFile
	if outFile == "" {
		prefix := strings.Replace(r.Prefix, "/", "_", -1)
		prefix = strings.Replace(prefix, ".", "_", -1)

		outFile = r.HomeDir + "/" + prefix + "_" + time.Now().Format("20060102150405") + ".png"
	}
	return outFile
}
