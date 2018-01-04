### 查看调用方法的时候，接收者是值类型或者引用类型是否有影响

#### 注意点
 - 不管方法是怎么声明的，值类型或者引用类型(指针类型)都是可以用声明的方法的
 - 接收者是值类型的话，传入方法的只是这个变量的副本，修改副本的值对变量本身是没用的。
 但是接收者是指针类型的话，是可以修改变量的值的

#### 参考资料
http://www.cnblogs.com/zlingh/p/5701785.html?hmsr=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com