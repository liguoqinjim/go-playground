package main

import "log"

type T struct {
	Name string
}

func (t T) M1() {
	t.Name = "name1"
}

func (t *T) M2() {
	t.Name = "name2"
}

type Intf interface {
	M1()
	M2()
}

//嵌入一个值类型
type S1 struct{ T }

//嵌入一个引用类型
type S2 struct{ *T }

func main() {
	//S1
	t1 := T{"t1"}
	s1 := S1{t1}
	log.Println("M1调用前:", s1.Name)
	s1.M1()
	log.Println("M1调用后:", s1.Name)

	log.Println("M2调用前:", s1.Name)
	s1.M2()
	log.Println("M2调用后:", s1.Name)
	//这里t1的值没有变是因为这时候，s里面的t是一个副本，和t1是没有关系的
	log.Println("M2调用后:t1.Name=", t1.Name)

	//S2
	t2 := &T{"t2"}
	s2 := S2{t2}
	log.Println("M1调用前:", s2.Name)
	s2.M1()
	log.Println("M1调用后:", s2.Name)

	log.Println("M2调用前:", s2.Name)
	s2.M2()
	log.Println("M2调用后:", s2.Name)
	//这里的值是会变的因为传递给s2的是一个指针
	log.Println("M2调用后:t2.Name=", t2.Name)

}
