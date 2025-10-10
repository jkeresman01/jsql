package db

import (
	"fmt"
	"strings"

	"github.com/jkeresman01/jsql/internal/db/model"
)

func Insert(db *model.Database, table string, values []string) {
	t, ok := db.Tables[table]
	if !ok {
		t = &model.Table{Name: table}
		db.Tables[table] = t
	}
	t.Rows = append(t.Rows, &model.Row{Values: values})
	fmt.Printf("%d row inserted.\n", 1)
}

func SelectAll(db *model.Database, table string) {
	t, ok := db.Tables[table]
	if !ok {
		fmt.Printf("Error: table '%s' does not exist.\n", table)
		return
	}

	if len(t.Rows) == 0 {
		fmt.Println("(0 rows)")
		return
	}

	numCols := len(t.Rows[0].Values)

	// Compute column widths
	colWidths := make([]int, numCols)
	for _, row := range t.Rows {
		for i, val := range row.Values {
			if len(val) > colWidths[i] {
				colWidths[i] = len(val)
			}
		}
	}

	// Build a horizontal line
	line := "+"
	for _, w := range colWidths {
		line += strings.Repeat("-", w+2) + "+"
	}

	// Print header
	fmt.Println(line)
	fmt.Printf("| %s |\n", strings.ToUpper(table))
	fmt.Println(line)

	// Print rows
	for _, row := range t.Rows {
		fmt.Print("|")
		for i, val := range row.Values {
			fmt.Printf(" %-*s |", colWidths[i], val)
		}
		fmt.Println()
	}
	fmt.Println(line)

	fmt.Printf("(%d rows)\n", len(t.Rows))
}
