package database

import (
	"github.com/ssedrick/gql/graph"
	"github.com/ssedrick/gql/syntax"
)

type Database struct {
	tables map[string]graph.Graph
}

func (db *Database) Load() {
}

func (db *Database) AddTable(data *syntax.AST) {

}
