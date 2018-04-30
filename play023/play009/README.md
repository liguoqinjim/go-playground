###

#### 注意点
 - go test -race。可以竞争检测
 - 注意一下reuslt这个struct，没有给定fieldName，就是一个string和一个bool。
 但是这里也只能有一个string或者一个bool，之后可以直接用string或者bool来获取字段的值