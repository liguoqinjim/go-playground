### 不安全的并发访问
goroutine数量小的时候，比如10程序会有几率出现错误。但是一旦调大goroutine的数量，比如10000，那么基本上都会错的。

出现的错误: `concurrent map read and map write`

#### 参考资料
http://www.jianshu.com/p/10a998089486