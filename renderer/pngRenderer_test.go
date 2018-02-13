package renderer

import (
	"testing"

	"github.com/windler/dotgraph/graph"
	"github.com/windler/dotgraph/renderer/mocks"
)

func TestPngRenderer(t *testing.T) {
	execMock := &mocks.DotExecutor{}

	graph := graph.New("testgraph")
	renderer := &PNGRenderer{
		OutputFile: "/tmp/out.png",
		executor:   execMock,
	}

	execMock.On("Render", "/tmp/out.png.dot", "/tmp/out.png").Return(nil)
	renderer.Render(graph.String())

	execMock.AssertExpectations(t)
}
