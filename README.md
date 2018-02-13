[![Go Report Card](https://goreportcard.com/badge/github.com/windler/dotgraph)](https://goreportcard.com/report/github.com/windler/dotgraph) [![CircleCI](https://circleci.com/gh/windler/dotgraph.svg?style=svg)](https://circleci.com/gh/windler/dotgraph)
# dotgraph
`dotgraph` is a package that lets you create and render [graphviz dot graphs](https://www.graphviz.org/).

## Installation
```bash
go get github.com/windler/dotgraph/graph
go get github.com/windler/dotgraph/renderer
```

If you want to render graphs make sure you have [graphviz](https://www.graphviz.org/) installed.

## Usage
### Creating a new graph
```go
graph := graph.New("my_graph")
```

### Adding nodes
```go
graph.AddNode("first node")
graph.AddNode("second node")
graph.AddNode("third node")
```
### Adding directed edges
```go
graph.AddDirectedEdge("first node", "second node", "edge label")
```

### Assiging dot graph attributes
You can apply any [dot attributes](https://graphviz.gitlab.io/_pages/doc/info/attrs.html).

#### Graph attributes
```go
graph.SetGraphOptions(DotGraphOptions{
    "bgcolor": "#333333",
})
```

#### Global node attributes
```go
graph.SetGraphOptions(dotgraph.DotGraphOptions{
    "fillcolor": "#336699",
    "style":     "filled",
    "fontcolor": "white",
    "fontname":  "Courier",
    "shape":     "rectangle",
})
```

#### Node attributes for nodes
If you want to assign attributes for nodes matching a pattern you can do the following:

```go
graph.AddNodeGraphPatternOptions("first", dotgraph.DotGraphOptions{
    "shape": "oval",
})
```

#### Global edge attributes
```go
graph.SetEdgeGraphOptions(dotgraph.DotGraphOptions{
    "arrowhead": "open",
    "color":     "white",
    "fontcolor": "white",
    "splines":   "curved",
})
```

#### Edge attributes 
If you want to assign attributes for an edge that references a certain node you can do the following:

```go
graph.AddEdgeGraphPatternOptions("first", dotgraph.DotGraphOptions{
    "color": "black",
})
```

### Render a graph
```go
r := &renderer.PNGRenderer{
   OutputFile: "/tmp/my_graph.png",
}

r.Render(graph.String())
```
