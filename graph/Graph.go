package graph

import "github.com/ssedrick/gql/nodes"

type Graph struct {
	nodes map[string]nodes.Node
	edges map[string][]Edge
}

type Edge struct {
	to   string
	from string
}
