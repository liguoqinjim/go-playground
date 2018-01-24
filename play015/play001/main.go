package main

import "log"

func main() {
	//遍历的时候增加元素
	a := []int{1, 2, 3, 4, 5}
	for k, v := range a {
		if k == 0 {
			a = append(a, 6)
			log.Println("a.length=", len(a))
		}

		log.Printf("k=%d,v=%d", k, v)
	}

	//遍历的时候减去元素
	for k, v := range a {
		if k == 0 {
			a = append(a[:2], a[2+1:]...)
			log.Println("a.length=", len(a))
		}

		log.Printf("k=%d,v=%d", k, v)
	}
}
