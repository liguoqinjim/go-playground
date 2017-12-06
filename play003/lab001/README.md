### 示例
go_version_test.go是测试文件

#### 运行代码
`mockgen.exe -destination spider/mock_spider.go -package spider play003/lab001/spider Spider`

命令里面的destination表示生成的文件名，package是生成文件的包名。`play003/lab001/spider Spider`表示这个文件的里面的这个接口