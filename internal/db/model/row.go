package model

// Row represents one record in a table.
// All values are stored as strings for now (typed columns come later).
type Row struct {
	Values []string
}
