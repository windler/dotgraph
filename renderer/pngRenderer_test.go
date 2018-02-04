package renderer

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/windler/dotgraph/graph"
)

func TestPngRenderer(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "godepgtest")

	graph := graph.New("testgraph")
	renderer := &PNGRenderer{
		HomeDir: tmpDir,
		Prefix:  "my.pkg",
	}

	renderer.Render(graph.String())

	files, _ := ioutil.ReadDir(tmpDir)

	assert.True(t, len(files) == 2)
	assert.True(t, strings.HasPrefix(files[0].Name(), "my_pkg"))
	assert.True(t, strings.HasSuffix(files[0].Name(), "png"))
}
