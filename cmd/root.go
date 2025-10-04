/*

Copyright © 2025 Josip Keresman

*/
package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var version = "v0.1.0"

var rootCmd = &cobra.Command{
	Use:   "jsql",
	Short: "jsql is a lightweight SQL REPL",
	Long: `jsql is a mini psql-style REPL that lets you run simple SQL-like
commands on an in-memory database.`, 
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to jsql! Use --help to see available commands.")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jsql.yaml)")
	rootCmd.Version = version;

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


