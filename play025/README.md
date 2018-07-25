# channel的使用

|文件夹|简介|说明|
|---|---|---|
|play001|channel有无buffer的区别||
|play002|非缓存类型的channel的发送和接收||
|play003|从closed的channel里面read|
|play004|channel的size|
|play005|range|
|play006|channel为nil下的一些行为|


## 注意点
### play001
 - channel要是没有receiver的情况下，send数据到channel里面，程序是会报错的。deadlock
 - 最简单的例子，一行代码是`ch <- 1`，它的下面一行就是recever`<-ch`，但是这样还是会报deadlock的。
 - 要是没有数据send了，这个时候还有receiver在监听，那么也会deadlock

### play002

### play003

### play004
 - 非缓存型的channel，len一直是返回0
 - 缓存型的channel，len是返回还有多少未被read的数据
 - channel有没有被close是不会影响len的返回的

### play005
 - range使用在channel的时候，只能得到一个返回值。`for v := range channel`
 - 显式的关闭channel(`close(channel)`)，会导致range的退出