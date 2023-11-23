package studentcontral

import (
	"fmt"
	"os"
)

func main() {
	showMenu()
	println("请输入：")
	var input int
	fmt.Scanln("%d", &input)
	switch input {
	case 1: //添加学生
		fmt.Println("请输入学员学号名字，班级")
		s1, err := newStuByUser()
		if err == nil {
			usersys.addstudent(s1)
		}
	case 2: //修改学生信息
		fmt.Println("请输入该学生的学号，及其要修改的名称和班级：（用空格隔开）")
		s1, err := newStuByUser()
		if err == nil {
			usersys.updatestudent(s1)
		}
	case 3: //列出所有学生信息
		fmt.Println("以下是所有学生信息")
		usersys.showStudents()
	case 4:
		os.Exit(0)
	default:
		println("输入不正确")
	}
}
func showMenu() {
	info := `
	==========================
			学生管理系统
	==========================
	1.添加学生信息
	2.编辑学生信息
	3.展示学生信息
	4.exit
`
	fmt.Print(info)
}
func newStuByUser() (*student, error) {
	var (
		id   int
		name string
		age  int
	)
	_, err := fmt.Scanln(&id, &name, &age)
	if err != nil {
		print(err)
	}
	return newStudent(id, name, age), err
}
