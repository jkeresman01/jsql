package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jkeresman01/jsql/internal/db"
)

var replCmd = &cobra.Command{
	Use:   "repl",
	Short: "Start the interactive jsql shell",
	Run: func(cmd *cobra.Command, args []string) {
		startREPL()
	},
}

func init() {
	rootCmd.AddCommand(replCmd)
}

func startREPL() {
	fmt.Println("Welcome to jsql v0.1")
	fmt.Println("Type \\help for help, \\exit to quit.\n")

	reader := bufio.NewReader(os.Stdin)
	var statement strings.Builder 

	database := db.NewDatabase()

	for {
		if statement.Len() == 0 {
			fmt.Print("jsql> ")
		} else {
			fmt.Print("....> ")
		}

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		line = strings.TrimSpace(line)

		if statement.Len() == 0 {
			switch line {
			case "\\exit", "\\quit":
				fmt.Println("Goodbye!")
				return
			case "\\help":
				fmt.Println("Available commands:")
				fmt.Println("  SQL-like: INSERT INTO table VALUES (...);")
				fmt.Println("             SELECT * FROM table;")
				fmt.Println("  Meta: \\help, \\exit")
				continue
			case "":
				continue
			}
		}

		statement.WriteString(" ")
		statement.WriteString(line)

		if strings.HasSuffix(line, ";") {
			query := strings.TrimSpace(statement.String())
			statement.Reset()

			executeSQL(database, query)
		}
	}
}

func executeSQL(database *db.Database, query string) {
	parser := db.NewParser(query)
	stmt, err := parser.Parse()
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	switch stmt.Type {
	case "INSERT":
		database.Insert(stmt.Table, stmt.Values)
	case "SELECT":
		database.SelectAll(stmt.Table)
	default:
		fmt.Println("Unknown statement type.")
	}
}
