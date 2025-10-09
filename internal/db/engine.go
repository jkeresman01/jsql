package db

import (
	"fmt"
	"strings"
)

type Row []string

type Table struct {
	Name string
	Rows []Row
}

type Database struct {
	Tables map[string]*Table
}

func NewDatabase() *Database {
	return &Database{
		Tables: make(map[string]*Table),
	}
}

func (db *Database) Insert(table string, values []string) {
	t, ok := db.Tables[table]
	if !ok {
		t = &Table{Name: table}
		db.Tables[table] = t
	}
	t.Rows = append(t.Rows, values)
	fmt.Printf("%d row inserted.\n", 1)
}

func (db *Database) SelectAll(table string) {
	t, ok := db.Tables[table]
	if !ok {
		fmt.Printf("Error: table '%s' does not exist.\n", table)
		return
	}

	if len(t.Rows) == 0 {
		fmt.Println("(0 rows)")
		return
	}

	// Print simple table
	fmt.Println("+--------------------------+")
	fmt.Printf("| %s |\n", strings.ToUpper(table))
	fmt.Println("+--------------------------+")
	for _, row := range t.Rows {
		fmt.Println(strings.Join(row, " | "))
	}
	fmt.Printf("(%d rows)\n", len(t.Rows))
}

