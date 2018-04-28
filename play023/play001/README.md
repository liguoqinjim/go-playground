### Hello world

#### 注意点
 - `t.Run`这样可以在一个测试函数里面做subtest
 - t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper. This will help other developers track down problems easier. If you still don't understand, comment it out, make a test fail and observe the test output.

#### 参考资料
 - https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world