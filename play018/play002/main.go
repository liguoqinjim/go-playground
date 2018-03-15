package main

import "log"

type Person struct {
	Name string
}

func main() {
	p := Person{Name: "小明"}
	log.Printf("%p", &p)
	modify(&p)
	log.Println(p.Name)
}

func modify(p *Person) {
	log.Printf("%p", &p)
	p.Name = "小红"
}
