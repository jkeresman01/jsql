package parser

import (
	"reflect"
	"testing"
)

func TestParseInsert(t *testing.T) {
	query := "INSERT INTO users VALUES (1, 'Alice');"

	p := NewParser(query)
	stmt, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if stmt.Type != "INSERT" {
		t.Fatalf("expected Type=INSERT, got %s", stmt.Type)
	}
	if stmt.Table != "users" {
		t.Fatalf("expected Table=users, got %s", stmt.Table)
	}

	expectedVals := []string{"1", "Alice"}
	if !reflect.DeepEqual(stmt.Values, expectedVals) {
		t.Fatalf("expected Values=%v, got %v", expectedVals, stmt.Values)
	}
}

func TestParseSelect(t *testing.T) {
	query := "SELECT * FROM users;"

	p := NewParser(query)
	stmt, err := p.Parse()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if stmt.Type != "SELECT" {
		t.Fatalf("expected Type=SELECT, got %s", stmt.Type)
	}
	if stmt.Table != "users" {
		t.Fatalf("expected Table=users, got %s", stmt.Table)
	}
}
