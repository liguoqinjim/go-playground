### Defer in a block
要注意block里面的defer不是在block的最后运行的，而是在外层函数的最后运行的

#### 运行结果
##### 错误的情况
![Imgur](https://i.imgur.com/nP8nELg.png)

##### 解决方法
![Imgur](https://i.imgur.com/IjwQEjk.png)