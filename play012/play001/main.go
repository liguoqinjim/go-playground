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

func main() {
	t1 := T{"t1"}

	//调用M1
	log.Println("M1调用前:", t1.Name)
	t1.M1()
	log.Println("M1调用后:", t1.Name)
	//调用M2
	log.Println("M2调用前:", t1.Name)
	t1.M2()
	log.Println("M2调用后:", t1.Name)

	t2 := &T{Name: "t2"}
	//调用M1
	log.Println("M1调用前:", t2.Name)
	t2.M1()
	log.Println("M1调用后:", t2.Name)
	//调用M2
	log.Println("M2调用前:", t2.Name)
	t2.M2()
	log.Println("M2调用后:", t2.Name)
}
