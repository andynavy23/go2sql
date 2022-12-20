/**
 * Author: Andrew FU
 * File: go2sql.go
 */

package go2sql

import (
	"testing"
)

func TestSelect(t *testing.T) {
	sqlQuery := NewSqlQuery()
	sqlQuery.Select([]string{
		`col_a`,
		`col_b`,
	})
	expected := `SELECT
  col_a,
  col_b
`

	if got := sqlQuery.ToSql(); got != expected {
		t.Errorf("Got: %q", got)
		t.Errorf("Expected: %q", expected)
	}
}

func TestSelectDistinct(t *testing.T) {
	sqlQuery := NewSqlQuery()
	sqlQuery.SelectDistinct([]string{
		`col_a`,
		`col_b`,
	})
	expected := `SELECT DISTINCT
  col_a,
  col_b
`

	if got := sqlQuery.ToSql(); got != expected {
		t.Errorf("Got: %q", got)
		t.Errorf("Expected: %q", expected)
	}
}
