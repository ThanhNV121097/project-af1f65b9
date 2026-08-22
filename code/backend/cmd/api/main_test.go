package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeRow struct{ err error }

func (r fakeRow) Scan(dest ...any) error {
	if r.err != nil {
		return r.err
	}
	if len(dest) > 0 {
		if s, ok := dest[0].(*string); ok {
			*s = ""
		}
	}
	return nil
}

type fakeDB struct{ row fakeRow }

func (f fakeDB) QueryRowContext(context.Context, string, ...any) rowScanner { return f.row }

func TestGreetingNotFound(t *testing.T) {
	rec := httptest.NewRecorder()
	greeting(fakeDB{row: fakeRow{err: sql.ErrNoRows}}).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/greeting", nil))
	if rec.Code != http.StatusNotFound { t.Fatalf("status = %d", rec.Code) }
	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil { t.Fatal(err) }
	if body["error"].(map[string]any)["code"] != "greeting_not_found" { t.Fatalf("body = %v", body) }
}

func TestGreetingEmpty(t *testing.T) {
	rec := httptest.NewRecorder()
	greeting(fakeDB{row: fakeRow{err: nil}}).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/greeting", nil))
	if rec.Code != http.StatusUnprocessableEntity { t.Fatalf("status = %d", rec.Code) }
}

func TestGreetingInternalError(t *testing.T) {
	rec := httptest.NewRecorder()
	greeting(fakeDB{row: fakeRow{err: errors.New("db down")}}).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/greeting", nil))
	if rec.Code != http.StatusInternalServerError { t.Fatalf("status = %d", rec.Code) }
}
