package renderer

import (
	"fmt"
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
}

//Render renders graph content
func (r PNGRenderer) Render(graphContent string) {
	outFile := r.getOutputFile()
	dotFile := outFile + ".dot"

	fmt.Println(graphContent)

	err := ioutil.WriteFile(dotFile, []byte(graphContent), os.ModePerm)
	if err != nil {
		log.Fatal("Error writing dot file.", err.Error())
	}

	_, err = exec.Command("dot", "-Tpng", dotFile, "-o", outFile).Output()
	if err != nil {
		log.Fatal("Error creating png.", err.Error())
	}

	fmt.Println("Written to " + outFile)
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
