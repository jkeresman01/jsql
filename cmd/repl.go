package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var replCmd = &cobra.Command{
	Use:   "repl",
	Short: "Start the interactive jsql shell",
	Long: `Launches the jsql REPL (Read-Eval-Print Loop)
where you can type SQL-like commands or meta commands (\help, \exit).`,
	Run: func(cmd *cobra.Command, args []string) {
		startREPL()
	},
}

func init() {
	rootCmd.AddCommand(replCmd)
}

func startREPL() {
	fmt.Println("Welcome to jsql")
	fmt.Println("Type \\help for help, \\exit to quit.\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("jsql> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		line := strings.TrimSpace(input)

		switch {
		case line == "\\exit" || line == "\\quit":
			fmt.Println("Goodbye!")
			return

		case line == "\\help":
			fmt.Println("Available commands:")
			fmt.Println("  SQL-like: CREATE, INSERT, SELECT, etc.")
			fmt.Println("  Meta: \\help, \\exit, \\tables, \\schema")
			continue

		case line == "":
			continue

		default:
			// placeholder for SQL command handling
			if strings.HasSuffix(line, ";") {
				fmt.Println("Executing:", line)
			} else {
				fmt.Println("Statements must end with ';'")
			}
		}
	}
}

