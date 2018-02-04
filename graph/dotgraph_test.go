package graph

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyGraph(t *testing.T) {
	g := New("my test graph")

	assert.Equal(t, "my test graph", g.name)
	assert.Equal(t, 0, len(g.edges))
}

func TestDotFile(t *testing.T) {
	g := New("test_graph")

	g.AddNode("nodeA")
	g.AddNode("nodeB")
	g.AddNode("nodeC")
	g.AddNode("nodeD")

	g.AddDirectedEdge("nodeA", "nodeB", "")
	g.AddDirectedEdge("nodeC", "nodeB", "fancy")
	g.AddDirectedEdge("nodeA", "nodeC", "")

	g.SetGraphOptions(DotGraphOptions{
		"bgcolor": "#333333",
	})

	assert.Equal(t, 4, len(g.edges))
	assert.Equal(t, 2, len(g.edges[`"nodeA"`]))
	assert.Equal(t, 0, len(g.edges[`"nodeB"`]))
	assert.Equal(t, 1, len(g.edges[`"nodeC"`]))
	assert.Equal(t, 0, len(g.edges[`"nodeD"[attr1="val1 attr2="val2"]`]))

	dotFile := g.String()
	assert.True(t, strings.Contains(dotFile, `digraph test_graph`))
	assert.True(t, strings.Contains(dotFile, `"nodeA"->"nodeB"`))
	assert.True(t, strings.Contains(dotFile, `"nodeA"->"nodeC"`))
	assert.True(t, strings.Contains(dotFile, `"nodeC"->"nodeB"[label="fancy"]`))
	assert.True(t, strings.Contains(dotFile, `"nodeA"`))
	assert.True(t, strings.Contains(dotFile, `"nodeB"`))
	assert.True(t, strings.Contains(dotFile, `"nodeC"`))
	assert.True(t, strings.Contains(dotFile, `"nodeD"`))
}

func TestGetDependencies(t *testing.T) {
	g := New("test_graph")

	g.AddNode("nodeA")
	g.AddNode("nodeB")
	g.AddNode("nodeC")

	g.AddDirectedEdge("nodeA", "nodeB", "")
	g.AddDirectedEdge("nodeB", "nodeC", "fancy")

	assert.Equal(t, []string{`"nodeB"`}, g.GetDependencies("nodeA"))
	assert.Equal(t, []string{}, g.GetDependents("nodeA"))

	assert.Equal(t, []string{`"nodeC"`}, g.GetDependencies("nodeB"))
	assert.Equal(t, []string{`"nodeA"`}, g.GetDependents("nodeB"))

	assert.Equal(t, []string{}, g.GetDependencies("nodeC"))
	assert.Equal(t, []string{`"nodeB"`}, g.GetDependents("nodeC"))
}
