### Calling recover outside of a deferred func

#### 注意点
 - 出错的时候是recover在panic前调用，且不是defer调用的

#### 运行结果
![Imgur](https://i.imgur.com/qoM1cs0.png)