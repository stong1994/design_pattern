package iterator

import "fmt"

type Student struct{}

type Students struct {
	offset int
	data   []Student
}

func (s *Students) HasNext() bool {
	return s.offset >= len(s.data)
}

func (s *Students) Next() Student {
	if !s.HasNext() {
		panic("index out of range")
	}
	stu := s.data[s.offset]
	s.offset++
	return stu
}

func example(students Students) {
	for students.HasNext() {
		student := students.Next()
		fmt.Println(student)
	}
}

type Dept struct {
	name  string
	level int
	child []*Dept
}

func (d *Dept) Add(dept *Dept) {
	d.child = append(d.child, dept)
}

func (d *Dept) Print() {
	fmt.Printf("name: %s, level: %d", d.name, d.level)
}

func (d *Dept) Iterator(fnc func(dept *Dept)) {
	fnc(d)
	for _, v := range d.child {
		fnc(v)
	}
}

func example2() {
	// 一级部门
	dept1 := &Dept{
		name:  "root",
		level: 1,
		child: nil,
	}
	// 二级部门
	dept2 := &Dept{
		name:  "a",
		level: 2,
		child: nil,
	}
	dept3 := &Dept{
		name:  "b",
		level: 2,
		child: nil,
	}
	// 三级部门
	dept4 := &Dept{
		name:  "c",
		level: 3,
		child: nil,
	}
	dept1.Add(dept2)
	dept1.Add(dept3)
	dept2.Add(dept4)

	dept1.Iterator(func(dept *Dept) {
		dept.Print()
	})
}