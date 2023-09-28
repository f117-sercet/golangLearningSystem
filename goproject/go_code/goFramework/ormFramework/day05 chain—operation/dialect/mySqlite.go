package dialect

import (
	"reflect"
)

type mySqlite struct{}

var _ Dialect = (*mySqlite)(nil)

func init() {
	RegisterDialect("mySqlite", &mySqlite{})
}

func (m mySqlite) DataTypeOf(typ reflect.Value) string {
	//TODO implement me

	switch typ.Kind() {

	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"

	}
	panic("implement me")
}

func (m mySqlite) TableExistSQL(tableName string) (string, []interface{}) {
	//TODO implement me
	panic("implement me")
}
