package session

import (
	"database/sql"
	"golangLearningSystem/goproject/go_code/goFramework/ormFramework/02Day-databaseSql/log"
	"strings"
)

// Session 数据库连接相关操作
type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
}

// New 创建一个Session实例
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

// Clear 初始化一个连接
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

// DB 返回一个DB
func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// 多参数查询
func (s *Session) Raw(sql string, values ...interface{}) *Session {

	s.sql.WriteString(sql)
	s.sql.WriteString("")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}
