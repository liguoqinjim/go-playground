package main

import (
	"log"
	"strconv"
)

type Person struct {
	Name string
	Age  *int
}

func main() {
	i := 19
	p := Person{Name: "小明", Age: &i}
	log.Println("p=", p)
	modify(p)
	log.Println("p=", p)
	modify2(&p)
	log.Println("p=", p)
}

func (p Person) String() string {
	return "姓名为：" + p.Name + ",年龄为：" + strconv.Itoa(*p.Age)
}

func modify(p Person) {
	p.Name = "小红"
	*p.Age = 20
}

func modify2(p *Person) {
	p.Name = "小刚"
	*p.Age = 21
}
