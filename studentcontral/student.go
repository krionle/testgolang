package studentcontral

type student struct {
	id   int    //编号
	name string //姓名
	age  int    //年龄
}

func newStudent(id int, name string, age int) *student {
	return &student{
		id:   id,
		name: name,
		age:  age,
	}
}
