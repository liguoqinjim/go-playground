### learn-go-with-tests

|实验|简介|说明|
|---|---|---|
|play001|hello world||
|play002|Integers||
|play003|Iteration|迭代|
|play004|Arrays and slices|
|play005|Structs,methods&interfaces|
|play006|Pointers & errors|
|play007|Dependency Injection|
|play008|Mocking|
|play009|Concurrency|
|play010|Select|

#### tips
 - `go test`测试
 - `go test -v`正确的时候也会输出结果
 - `go test -v -run FuncName`指定测试FuncName方法
 - `t.Run()`方法可以做出分组测试
 - `t.Helper()`可以指定方法为Helper方法，在Helper方法里面出错的话，报错的行数不会是Helper方法的行数
 - 单元测试是Test开头的，
 - 性能测试是Benchmark开头的，注意循环里面使用`for i := 0; i < b.N; i++ {...}`
 - `go test -bench=.`用这个命令的话可以测试benchmark
 - 示例函数是Example开头的，注意最后要加Example
 - `go test -cover`，可以查看现在的测试覆盖率。不要盲目的追求100%的覆盖率，
  只要是按照TDD的开发的话，覆盖率应该都是可以的
 - `reflect.DeepEqual`，这个在使用的时候要注意下，这个是不判断type的，也就是所有的数据类型都可以调用。所以我们调用的时候要注意一下。
 有的时候要比对slice的话，可以用`reflect.DeepEqual`
 - %+v,%#v的区别，两个的format不一样
 - Remember to only do enough to make the tests run. We need to make sure our test fails correctly with a clear error message.
   也就是说，我们一开始只要测试用例可以build过就可以了。不需要去关注这个测试结果正确与否。
 - `https://github.com/kisielk/errcheck`，这个工具可以检查，我们是否有漏掉的对error的检查。
 - mock测试，play008
 - `go test -race`，可以竞争检测
 - http测试的时候可以看下`net/http/httptest`