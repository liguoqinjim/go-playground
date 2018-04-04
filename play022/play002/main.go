package main

var s0 []int

func main() {

}

//会内存泄漏
func g(s1 []int) {
	s0 = s1[len(s1)-30:]
}

func g1(s1 []int) {
	s0 = append([]int{nil}, s1[len(s1)-30:]...)
}
