package syntax

import (
	"fmt"
	"strings"
)

const (
	TYPE   string = "type"
	QUERY  string = "query"
	MUTATE string = "mutation"
	NONE   string = "NONE"
)

type AST struct {
	Verb string
}

type ParseError struct {
	reason string
}

func (perr *ParseError) Error() string {
	return perr.reason
}

func Parse(data string) (*AST, *ParseError) {
	ast := new(AST)

	if strings.HasPrefix(data, TYPE) {
		fmt.Println("it's a type")
		ast.Verb = TYPE
	} else if strings.HasPrefix(data, QUERY) {
		fmt.Println("it's a query")
		ast.Verb = QUERY
	} else if strings.HasPrefix(data, MUTATE) {
		fmt.Println("it's a mutation")
		ast.Verb = MUTATE
	} else {
		return nil, &ParseError{"Base Command unrecognized. Expected one of [\"type\", \"query\", \"mutation\"]"}
	}

	return ast, nil
}
