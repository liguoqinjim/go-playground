package main

import "log"

func main() {
	//increase the capacity of a slice
	s := []byte{1, 2, 3}
	log.Println("s=", s)
	log.Printf("s.len=%d,cap=%d", len(s), cap(s))
	t := make([]byte, len(s), (cap(s)+1)*2)
	for i := range s {
		t[i] = s[i]
	}
	s = t
	log.Println("s=", s)
	log.Printf("s.len=%d,cap=%d", len(s), cap(s))

	//copy
	s2 := []byte{1, 2, 3}
	log.Println("s2=", s2)
	log.Printf("s2.len=%d,cap=%d", len(s2), cap(s2))
	t2 := make([]byte, len(s), (cap(s2)+1)*2)
	copy(t2, s2)
	s2 = t2
	log.Println("s2=", s2)
	log.Printf("s2.len=%d,cap=%d", len(s2), cap(s2))

	//append
	s3 := make([]int, 1)
	log.Println("s3=", s3)
	log.Printf("s3.len=%d,cap=%d", len(s3), cap(s3))
	s3 = append(s3, 1, 2, 3)
	log.Println("s3=", s3)
	log.Printf("s3.len=%d,cap=%d", len(s3), cap(s3))
	s4 := []int{5, 6, 7}
	s3 = append(s3, s4...)
	log.Println("s3=", s3)
	log.Printf("s3.len=%d,cap=%d", len(s3), cap(s3))
}
