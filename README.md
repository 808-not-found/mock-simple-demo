# 接口测试-续2

再次运行 `go test`

```
exit status 1
FAIL    mock-simple-demo        0.003s
```

我们明明没有更改我们的代码，却出错了，原来是有人不小心修改了 `text.txt` 的内容。

真实情况中，你可能需要连接数据库，但是你不想启动数据库。于是就有了 mock 测试，对于接口的 mock 测试，我们采用 `testify/mock` 并使用 `mockery` 生成代码。

现在我们使用 `mockery --all`，你会发现创建了一个 `mocks` 文件夹，里面的内容我们完全不用修改。

接下来，你需要将 `TestReverseOutput` 这个函数修改为：

```go
import (
	"mock-simple-demo/mocks"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
func TestReverseOutput(t *testing.T) {
	mockFile := mocks.NewFile(t)
	str := "你好\n世界"
	mockFile.On("GetContent").Return(func() string {
		return str
	})
	mockFile.On("GetNthLine", mock.AnythingOfType("int")).Return(func(n int) string {
		lines := strings.Split(str, "\n")
		return lines[n-1]
	})
	assert.Equal(t, "世界\n你好\n", ReverseOutput(mockFile))
}

```

`go run test`:

```
PASS
ok      mock-simple-demo        0.003s
```