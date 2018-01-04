package main

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

func main() {
	var t1 = T{"t1"}
	t1.M1()
	t1.M2()

	//错误赋值(因为可以认为t1是没有实现M2的，而是t1的指针实现了M2)
	//var t2 Intf = t1
	//正确赋值(以为t1的指针是实现了M2的，指针实现了的话也可以认为值类型实现了)
	var t2 Intf = &t1
	t2.M1()
	t2.M2()
}
