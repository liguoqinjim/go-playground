### function as handler
这个直接用function作为handler，有两种方式：

1. 先强制转换到`http.HandlerFunc`，再调用`mux.Handle()`
2. 直接调用`mux.HandleFunc()`