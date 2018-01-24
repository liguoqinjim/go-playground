package main

import "log"

func main() {
	//删除元素
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	del := false
	for k, v := range m {
		if !del {
			delete(m, 2)
			del = true
			log.Println("m.length=", len(m))
		}
		log.Printf("k=%d,v=%d", k, v)
	}

	//添加元素
	add := false
	for k, v := range m {
		if !add {
			m[6] = 6
			m[7] = 7
			add = true
			log.Println("m.length=", len(m))
		}
		log.Printf("k=%d,v=%d", k, v)
	}
}
