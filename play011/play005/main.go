package main

import "log"

type Car struct {
	model string
}

func (c Car) PrintModel() {
	log.Println("c.model=", c.model)
}

func (c *Car) PrintModel2() {
	log.Println("*c.model=", c.model)
}

func main() {
	//调用非指针
	//c := Car{model: "DeLorean DMC-12"}
	//defer c.PrintModel()
	//c.model = "Chevrolet Impala"

	//调用指针
	c := &Car{model: "Delorean DMC-12"}
	defer c.PrintModel()
	defer c.PrintModel2()
	c.model = "Chevrolet Impala"
}
