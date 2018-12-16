package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	// "github.com/ssedrick/gql/database"
	"github.com/ssedrick/gql/syntax"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	fmt.Printf("Welcome to GQL, database based on the GraphQL specification\n")
	// db := new(database.Database)

	for {
		fmt.Printf(">")
		input, err := stdin.ReadString(';')
		fmt.Printf("input: %s\n", input)
		if err == io.EOF {
			fmt.Println("Goodbye")
			return
		} else if err != nil {
			panic(fmt.Errorf("Unexpected error occured: %s\n", err))
		}
		command, err := syntax.Parse(input)
		if err == nil {
			fmt.Printf("Command is of type %s\n", command.Verb)
		} else {
			fmt.Printf("error: %s\n", err)
		}
		// ast := new(syntax.AST)
		// db.AddTable(ast)
	}
}
