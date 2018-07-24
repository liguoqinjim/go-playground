# channel的使用

|文件夹|简介|说明|
|---|---|---|
|play001|channel有无buffer的区别||
|play002|channel的size|
|play003|range|

## 注意点
### play001
 - channel要是没有receiver的情况下，send数据到channel里面，程序是会报错的。deadlock
 - 最简单的例子，一行代码是`ch <- 1`，它的下面一行就是recever`<-ch`，但是这样还是会报deadlock的。
 - 要是没有数据send了，这个时候还有receiver在监听，那么也会deadlock

