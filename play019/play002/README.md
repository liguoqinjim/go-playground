### Calling defer in the wrong order

#### 注意点
 - 第一部分的代码会出错是因为没有判断之前的运行是否会出错，不确定res是否会为空，为空的时候就会出错
 - 解决方法就像第二部分代码里面一样，先判断一下是否出错。没错的话再使用defer

#### 运行结果
##### 出错的结果
![Imgur](https://i.imgur.com/q3918fI.png)

##### 正确结果
![Imgur](https://i.imgur.com/Glf67ZP.png)