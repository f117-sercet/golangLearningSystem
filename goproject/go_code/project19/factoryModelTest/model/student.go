package model

type Student struct {
	Name string
	Age  int64
}
type student struct {
	Name string
	Age  int64
}

type Student1 struct {
	Name string
	age  int64
}

// 由于是私有。因此只能在model使用
// 通过工厂模式来解决

func NewStudent(n string, s int64) *student {

	return &student{

		Name: n,
		Age:  s,
	}

}

// 若有字段小写，用方法调用
func getAge(s *Student1) int64 {

	return s.age
}
