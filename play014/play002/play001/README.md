### slice的赋值，slice的len和cap

#### 注意点
 - slice不会复制值，所以都子slice的修改会修改到slice本身
 - len和cap，`[]`符号，最大可以到cap，而不是最大到len，但是超过cap会报错

#### 运行结果
![Imgur](https://i.imgur.com/DkY8DI7.png)