### slice

#### 注意点
 - 这里用Person这个结构体去模拟slice
 ```
 type slice struct {
     array unsafe.Pointer
     len   int
     cap   int
 }
 ```
 - 可以看到modify里面，Name是无法修改的，Age是可以修改的。和slice是一样的，array是可以修改的，但是len和cap是没法修改的。
 - 如果都想修改，是要传递slice的指针才可以的

#### 运行结果
![Imgur](https://i.imgur.com/6wLEyKX.png)

#### 参考资料
https://mp.weixin.qq.com/s/iKhhrIQ7zEMhJiVuegzPYg