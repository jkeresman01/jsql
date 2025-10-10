package db

import (
	"fmt"

	"github.com/jkeresman01/jsql/internal/db/model"
)

type DatabaseManager struct {
	Databases map[string]*model.Database
	Current   *model.Database
}

func NewManager() *DatabaseManager {
	m := &DatabaseManager{
		Databases: make(map[string]*model.Database),
	}
	m.CreateDatabase("default")
	m.Use("default")
	return m
}

func (m *DatabaseManager) CreateDatabase(name string) {
	if _, exists := m.Databases[name]; exists {
		fmt.Printf("Database '%s' already exists.\n", name)
		return
	}

	db := model.NewDatabase(name)
	m.Databases[name] = db
	fmt.Printf("Database '%s' created.\n", name)
}

func (m *DatabaseManager) DropDatabase(name string) {
	if _, exists := m.Databases[name]; !exists {
		fmt.Printf("Error: database '%s' does not exist.\n", name)
		return
	}

	delete(m.Databases, name)
	fmt.Printf("Database '%s' dropped.\n", name)

	if m.Current != nil && m.Current.Name == name {
		m.Current = nil
	}
}

func (m *DatabaseManager) Use(name string) {
	db, ok := m.Databases[name]
	if !ok {
		fmt.Printf("Error: database '%s' does not exist.\n", name)
		return
	}
	m.Current = db
	fmt.Printf("Connected to database '%s'.\n", name)
}

// CurrentDB returns the active database.
func (m *DatabaseManager) CurrentDB() *model.Database {
	return m.Current
}
