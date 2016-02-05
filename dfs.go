// package dfs implements depth first search of graphs.
// See https://www.csd.uoc.gr/~hy583/reviewed_notes/dfs_dags.pdf.
package dfs

const (
	colorWhite = iota
	colorGray
	colorBlack
)

// Direction controls how the
// edges are explored from each vertex.
const (
	DirectionForward = iota
	DirectionBackward
)

// Vertex is a vertex in a graph.
type Vertex struct {
	Value interface{}
	Edges []Vertex
	color int
}

// Visitor is a callback invoked when exploring a graph.
type Visitor func(Vertex)

// Options provides control over the search algorithms.
type Options struct {
	Direction int
}

// Visit visits every node in a graph.
func Visit(root Vertex, visitor Visitor, options Options) {
	root.color = colorGray
	if options.Direction == DirectionForward {
		for _, v := range root.Edges {
			if v.color == colorWhite {
				Visit(v, visitor, options)
			}
		}
	} else if options.Direction == DirectionBackward {
		for i := len(root.Edges) - 1; i >= 0; i-- {
			if root.Edges[i].color == colorWhite {
				Visit(root.Edges[i], visitor, options)
			}
		}
	} else {
	}
	root.color = colorBlack
	visitor(root)
}

// TopologicalSort sorts the vertices of a directed acyclic graph.
// This just does a depth first search and outputs the vertices in
// reverse order.
func TopologicalSort(root Vertex, options Options) []interface{} {
	vals := []interface{}{}

	visitor := func(v Vertex) {
		vals = append(vals, v.Value)
	}

	Visit(root, visitor, options)

	var (
		numVals = len(vals)
		ret     = make([]interface{}, numVals)
	)
	for i, val := range vals {
		ret[numVals-1-i] = val
	}
	return ret
}
