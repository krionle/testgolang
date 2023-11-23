package studentcontral

import "fmt"

type usersys struct {
	num      int
	students []*student
}

// num 学生数量
func newUsersys() *usersys {
	return &usersys{
		num:      0,
		students: make([]*student, 0, 100),
	}
}

// 添加学生
func (sys *usersys) addstudent(stu *student) {
	sys.students = append(sys.students, stu)
	sys.num += 1
	println("=======添加成功=======")

}

// 学生信息更新
func (sys *usersys) updatestudent(stu *student) {
	for i, v := range sys.students {
		if stu.id == v.id {
			sys.students[i] = stu
			fmt.Println("======修改成功======")
			return
		}
	}
}

// 列出所有学生信息
func (sys *usersys) showStudents() {
	for _, v := range sys.students {
		fmt.Println("编号：", v.id, "姓名：")
	}
}
