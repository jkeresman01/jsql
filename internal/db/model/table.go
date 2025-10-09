package model

// Table represents a single in-memory table.
// Columns are inferred for now, but later can include schema info.
type Table struct {
	Name    string
	Columns []string
	Rows    []*Row
}
