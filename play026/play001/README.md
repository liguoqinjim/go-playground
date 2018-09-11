# play001

## NOTICE
 - defer后入先出
 - panic 需要等defer 结束后才会向上传递。 出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。

## 运行结果
![Imgur](https://i.imgur.com/ClTAsyE.png)