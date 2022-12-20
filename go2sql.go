/**
 * Author: Andrew FU
 * File: go2sql.go
 */

package go2sql

import (
	"fmt"
)

type SqlQuery struct {
	queryCmd string `default:""`
}

func NewSqlQuery() *SqlQuery {
	sqlQuery := new(SqlQuery)
	return sqlQuery
}

func (sqlQuery *SqlQuery) Select(querys []string) {
	sqlQuery.queryCmd += "SELECT\n  "
	for index, query := range querys {
		sqlQuery.queryCmd += query
		if index != len(querys)-1 {
			sqlQuery.queryCmd += ",\n  "
		} else {
			sqlQuery.queryCmd += "\n"
		}
	}
}

func (sqlQuery *SqlQuery) SelectDistinct(querys []string) {
	sqlQuery.queryCmd += "SELECT DISTINCT\n  "
	for index, query := range querys {
		sqlQuery.queryCmd += query
		if index != len(querys)-1 {
			sqlQuery.queryCmd += ",\n  "
		} else {
			sqlQuery.queryCmd += "\n"
		}
	}
}

func (sqlQuery *SqlQuery) Where(querys []string) {
	sqlQuery.queryCmd += "WHERE\n  "
	for index, query := range querys {
		sqlQuery.queryCmd += query
		if index != len(querys)-1 {
			sqlQuery.queryCmd += "\n  AND "
		} else {
			sqlQuery.queryCmd += "\n"
		}
	}
}

func (sqlQuery *SqlQuery) GroupBy(querys []string) {
	sqlQuery.queryCmd += "GROUP BY\n  "
	for index, query := range querys {
		sqlQuery.queryCmd += query
		if index != len(querys)-1 {
			sqlQuery.queryCmd += ",\n  "
		} else {
			sqlQuery.queryCmd += "\n"
		}
	}
}

func (sqlQuery *SqlQuery) OrderBy(query string) {
	sqlQuery.queryCmd += "ORDER BY\n  "
	sqlQuery.queryCmd += query
	sqlQuery.queryCmd += "\n"
}

func (sqlQuery *SqlQuery) From(query string) {
	sqlQuery.queryCmd += "FROM\n  "
	sqlQuery.queryCmd += query
	sqlQuery.queryCmd += "\n"
}

func (sqlQuery *SqlQuery) ToSql() string {
	return sqlQuery.queryCmd
}

func RegexpExtract(str string, rule string) string {
	queryCmd := fmt.Sprintf(`REGEXP_EXTRACT(%s, %s)`, str, rule)
	return queryCmd
}

func Sum(str string) string {
	return fmt.Sprintf(`SUM(%s)`, str)
}

func Cast(str string, typeStr string) string {
	return fmt.Sprintf(`CAST(%s AS %s)`, str, typeStr)
}

func As(str string, as string) string {
	return fmt.Sprintf(`%s AS %s`, str, as)
}

func Between(str string, from string, to string) string {
	return fmt.Sprintf(`%s BETWEEN %s AND %s`, str, from, to)
}

func Equal(col string, value string) string {
	return fmt.Sprintf(`%s = "%s"`, col, value)
}
