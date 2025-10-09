package model

// Database represents the entire in-memory database instance.
type Database struct {
	Tables map[string]*Table
}

// NewDatabase initializes a new empty database.
func NewDatabase() *Database {
	return &Database{
		Tables: make(map[string]*Table),
	}
}
