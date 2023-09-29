package day05_chain_operation

import (
	"database/sql"
	"golangLearningSystem/goproject/go_code/goFramework/ormFramework/day05 chain—operation/dialect"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}
