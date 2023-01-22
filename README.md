# 函数测试

这次我们甚至没有 `main.go` 里面代码非常简单，请快速浏览。

测试代码在 `bad_test.go` 也十分简单。

`go test -gcflags="all=-l -N" -v ./...` 编译时需要禁用内联和编译优化，否则可能会 mock 失败或者报错

```
=== RUN   TestBad

  TestMockXXX 


0 total assertions

--- PASS: TestBad (0.00s)
PASS
ok      mock-simple-demo/bad    (cached)
?       mock-simple-demo/good   [no test files]
```

不支持真正的偷梁换柱成我们自定义逻辑的函数，但是 mock 一般也就是进行固定返回值的 mock

部分简单逻辑的判断可以参考 `[条件mock](https://github.com/bytedance/mockey/issues/5)`