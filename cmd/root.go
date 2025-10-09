package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jsql",
	Short: "jsql is a lightweight SQL REPL built in Go",
	Long: `jsql is a mini psql-style REPL that lets you run simple SQL-like
commands on an in-memory database. Built with Go and Cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		startREPL()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

