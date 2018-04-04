package main

func main() {

}

//会内存泄露
func g() []*int {
	s := []*int{new(int), new(int), new(int), new(int)}
	return s[1:3]
}

//解决方法
func h() []*int {
	s := []*int{new(int), new(int), new(int), new(int)}
	s1 := s[1:3]

	//指针重置了
	s[0] = nil
	s[len(s)-1] = nil
	return s1
}
