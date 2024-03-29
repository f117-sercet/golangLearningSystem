package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type mySql struct{}

var _Dialect = (*mySql)(nil)

func init() {
	RegisterDialect("mySql", &mySql{})
}

func (m mySql) DataTypeOf(typ reflect.Value) string {

	switch typ.Kind() {

	case reflect.Bool:
		return "bool"

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"

	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"

	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}

	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

func (m mySql) TableExistSQL(tableName string) (string, []interface{}) {

	args := []interface{}{tableName}
	return "SELECT name FROM tb_master WHERE type='table' and name = ?", args
}
