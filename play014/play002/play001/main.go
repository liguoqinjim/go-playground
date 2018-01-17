package main

import "log"

func main() {
	d := []int{10, 20, 30, 40}
	e := d[2:]
	log.Println("d=", d)
	log.Println("e=", e)

	//修改e，对e的修改也会使得d被修改
	e[1] = 100
	log.Println("modify e=", e)
	log.Println("modify d=", d)

	//len和cap
	log.Printf("d.len=%d,cap=%d", len(d), cap(d))
	log.Printf("e.len=%d,cap=%d", len(e), cap(e))

	s := []int{0, 1, 2, 3, 4}
	log.Printf("s.len=%d,cap=%d", len(s), cap(s))
	s = s[2:4]
	log.Printf("s.len=%d,cap=%d", len(s), cap(s))
	s = s[:cap(s)]
	log.Printf("s.len=%d,cap=%d", len(s), cap(s))
}
