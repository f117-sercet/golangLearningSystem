package schema

import (
	"database/sql"
	"golangLearningSystem/goproject/go_code/goFramework/ormFramework/day04-save-query/schema"
	"golangLearningSystem/goproject/go_code/goFramework/ormFramework/day05 chain—operation/clause"
	"golangLearningSystem/goproject/go_code/goFramework/ormFramework/day05 chain—operation/dialect"
	"strings"
)

type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect
	refTable *schema.Schema
	clause   clause.Clause
	sql      strings.Builder
	sqlVars  []interface{}
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

func (s *Session) DB() sql.DB {
	return s.db
}
