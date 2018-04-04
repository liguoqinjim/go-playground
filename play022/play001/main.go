package main

import "strings"

var s0 string

func main() {

	s1 := `cdsacnjdsanjk
cdascdsa
cdsacdsacdsa
cdsacdsacdsacdsavfdnjkc
sacdsanjckdsanjkcdlsanjkl
cdsacdsa`

	f(s1)
}

//会泄露内存的方法
func f(s1 string) {
	s0 = s1[:50]
}

//修改方法1
func f1(s1 string) {
	s0 = string([]byte(s1[:50]))
}

//修改方法2
func f2(s1 string) {
	s0 = (" " + s1[:50])[1:]
}

//修改方法3
func f3(s1 string) {
	var b strings.Builder
	b.Grow(50)
	b.WriteString(s1[:50])
	s0 = b.String()
	//b.Reset()
}
