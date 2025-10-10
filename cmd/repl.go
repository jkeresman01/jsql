package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jkeresman01/jsql/internal/db"
	"github.com/jkeresman01/jsql/internal/parser"
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

	manager := db.NewManager()

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
			switch {
			case line == "\\exit" || line == "\\quit":
				fmt.Println("Goodbye!")
				return

			case line == "\\help":
				fmt.Println("Available commands:")
				fmt.Println("  SQL-like: CREATE DATABASE name;")
				fmt.Println("             DROP DATABASE name;")
				fmt.Println("             INSERT INTO table VALUES (...);")
				fmt.Println("             SELECT * FROM table;")
				fmt.Println("  Meta: \\help, \\exit, \\connect <db>")
				continue

			case strings.HasPrefix(line, "\\connect "):
				parts := strings.SplitN(line, " ", 2)
				if len(parts) == 2 {
					manager.Use(parts[1])
				} else {
					fmt.Println("Usage: \\connect <database>")
				}
				continue

			case line == "":
				continue
			}
		}

		// accumulate statement lines until ';'
		statement.WriteString(" ")
		statement.WriteString(line)

		if strings.HasSuffix(line, ";") {
			query := strings.TrimSpace(statement.String())
			statement.Reset()

			executeSQL(manager, query)
		}
	}
}

func executeSQL(manager *db.DatabaseManager, query string) {
	p := parser.NewParser(query)
	stmt, err := p.Parse()
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	switch stmt.Type {
	case "CREATE_DATABASE":
		manager.CreateDatabase(stmt.Name)

	case "DROP_DATABASE":
		manager.DropDatabase(stmt.Name)

	case "INSERT":
		current := manager.CurrentDB()
		if current == nil {
			fmt.Println("Error: no database selected.")
			return
		}
		db.Insert(current, stmt.Table, stmt.Values)

	case "SELECT":
		current := manager.CurrentDB()
		if current == nil {
			fmt.Println("Error: no database selected.")
			return
		}
		db.SelectAll(current, stmt.Table)

	default:
		fmt.Println("Unknown statement type.")
	}
}
