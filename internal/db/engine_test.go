package db

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/jkeresman01/jsql/internal/db/model"
)

func captureOutput(f func()) string {
	// Temporarily redirect stdout
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	done := make(chan bool)
	go func() {
		_, _ = buf.ReadFrom(r)
		done <- true
	}()

	f()
	_ = w.Close()
	os.Stdout = old
	<-done
	return buf.String()
}

func TestInsertCreatesTableAndAddsRow(t *testing.T) {
	dbInstance := model.NewDatabase()

	out := captureOutput(func() {
		Insert(dbInstance, "users", []string{"1", "Alice"})
	})

	if len(dbInstance.Tables) != 1 {
		t.Fatalf("expected 1 table, got %d", len(dbInstance.Tables))
	}

	users, ok := dbInstance.Tables["users"]
	if !ok {
		t.Fatalf("expected table 'users' to exist")
	}

	if len(users.Rows) != 1 {
		t.Fatalf("expected 1 row, got %d", len(users.Rows))
	}

	if got := users.Rows[0].Values; len(got) != 2 || got[0] != "1" || got[1] != "Alice" {
		t.Fatalf("unexpected row values: %v", got)
	}

	if !strings.Contains(out, "1 row inserted.") {
		t.Fatalf("expected output to contain '1 row inserted.', got %q", out)
	}
}

func TestSelectAllPrintsRows(t *testing.T) {
	dbInstance := model.NewDatabase()
	Insert(dbInstance, "users", []string{"1", "Alice"})
	Insert(dbInstance, "users", []string{"2", "Bob"})

	out := captureOutput(func() {
		SelectAll(dbInstance, "users")
	})

	if !strings.Contains(out, "USERS") {
		t.Errorf("expected table header 'USERS' in output, got:\n%s", out)
	}

	if !strings.Contains(out, "1 | Alice") || !strings.Contains(out, "2 | Bob") {
		t.Errorf("expected both rows in output, got:\n%s", out)
	}

	if !strings.Contains(out, "(2 rows)") {
		t.Errorf("expected '(2 rows)' footer, got:\n%s", out)
	}
}

func TestSelectAllEmptyTable(t *testing.T) {
	dbInstance := model.NewDatabase()
	dbInstance.Tables["users"] = &model.Table{Name: "users"}

	out := captureOutput(func() {
		SelectAll(dbInstance, "users")
	})

	if !strings.Contains(out, "(0 rows)") {
		t.Errorf("expected '(0 rows)' for empty table, got:\n%s", out)
	}
}

func TestSelectAllMissingTable(t *testing.T) {
	dbInstance := model.NewDatabase()

	out := captureOutput(func() {
		SelectAll(dbInstance, "nonexistent")
	})

	expected := fmt.Sprintf("Error: table '%s' does not exist.", "nonexistent")
	if !strings.Contains(out, expected) {
		t.Errorf("expected error message %q, got:\n%s", expected, out)
	}
}
