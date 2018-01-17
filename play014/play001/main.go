package main

import "log"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println("a=\t\t\t", a)

	//复制
	b := make([]int, len(a))
	copy(b, a)
	log.Println("copy b=\t\t", b)

	//cut
	b = append(b[:1], b[5:]...)
	log.Println("cut b=\t\t", b)

	//delete
	//b = append(b[:2], b[2+1:]...)//删除index为2的元素
	b = append(b[:len(b)-1], b[len(b):]...) //删除最后一个
	log.Println("delete b=\t", b)

	//delete without preserving order(最后一个赋值到要删除的index位置，然后去掉最后一个元素)
	b[2], b = b[len(b)-1], b[:len(b)-1]
	log.Println("delete2 b=\t", b)

	//pop(队尾弹出)
	x, b := b[len(b)-1], b[:len(b)-1]
	log.Println("pop b=\t\t", b)
	log.Println("pop x=\t\t", x)

	//push
	b = append(b, 10)
	log.Println("push b=\t\t", b)
}
