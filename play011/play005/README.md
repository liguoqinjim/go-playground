### defer中调用变量
注意两个使用的方法。原理的话是因为defer是保存一份当时的代码。要是值类型的变量，就保存了值。
要是指针类型的变量，会保存一个新的指针，但是指向的地址不会变。

#### 运行结果
##### 调用非指针变量
![Imgur](https://i.imgur.com/p3IlYYb.png)

##### 调用指针变量
![Imgur](https://i.imgur.com/gG5yWfZ.png)