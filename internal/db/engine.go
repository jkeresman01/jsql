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

	fmt.Println("+--------------------------+")
	fmt.Printf("| %s |\n", strings.ToUpper(table))
	fmt.Println("+--------------------------+")
	for _, row := range t.Rows {
		fmt.Println(strings.Join(row.Values, " | "))
	}
	fmt.Printf("(%d rows)\n", len(t.Rows))
}
