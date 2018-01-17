### growing slice
`t := make([]byte, len(s), (cap(s)+1)*2)`，golang内部的对于slice的扩展也是这样操作的。
所以slice的cap都是2的倍数变大的

#### 运行结果
![Imgur](https://i.imgur.com/G8MY3bJ.png)