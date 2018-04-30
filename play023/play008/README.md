###

#### 注意点
 - 我们一开始用time.Sleep的方法来达成暂停一秒的目的，但是这样的话我们每次测试都要等这个几秒钟。
 然而这个几秒钟对于我们的程序实现是没有影响的，time.Sleep总不会错的。所以我们需要mock掉time.Sleep这个功能
 - SpySleeper里面的Calls，是用来记录调用了几次，在测试用例里面可以查看这个方法调用了几次，也就是实际逻辑中，停顿了几秒
 - 我们测试的时候只注意了，sleeper调用了4次，但是没有关注他是什么调用的。
 那么先调用了4次sleeper，然后再打印。测试也是能通过的。
 - 我们又了一个spy，去记录打印和停顿的操作顺序。这样更准确的测试
 - 关于mock的思考可以看参考资料里面的

#### 参考资料
 - https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking#but-isnt-mocking-evil