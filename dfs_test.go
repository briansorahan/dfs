package dfs

import (
	"reflect"
	"testing"
)

var graph1 = Vertex{
	Value: 1,
	Edges: []Vertex{
		{Value: 2},
		{
			Value: 3,
			Edges: []Vertex{
				{
					Value: 4,
					Edges: []Vertex{
						{Value: 5},
						{Value: 6},
					},
				},
			},
		},
	},
}

func TestVisit(t *testing.T) {
	vals := []int{}

	visitor := func(v Vertex) {
		val, ok := v.Value.(int)
		if !ok {
			t.Fatalf("Expected int, got %T", v.Value)
		}
		vals = append(vals, val)
	}

	Visit(graph1, visitor, Options{})

	expected, got := []int{2, 5, 6, 4, 3, 1}, vals
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func TestTopo(t *testing.T) {
	var (
		topo     = TopologicalSort(graph1, Options{})
		expected = make([]interface{}, len(topo))
		got      = topo
	)
	expected[0] = 1
	expected[1] = 3
	expected[2] = 4
	expected[3] = 6
	expected[4] = 5
	expected[5] = 2
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}
