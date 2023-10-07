package session

import (
	"golangLearningSystem/goproject/go_code/goFramework/ormFramework/day04-save-query/log"
	"reflect"
)

// 钩子变量
const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

func (s *Session) CallMethod(method string, value interface{}) {

	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method)

	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	}

	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {

		if v := fm.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {

				log.Error(err)
			}
		}
	}
	return

}

// 之后调用即可
