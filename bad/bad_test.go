package bad_test

import (
	"mock-simple-demo/bad"
	"mock-simple-demo/good"
	"testing"

	. "github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

func TestBad(t *testing.T) {
	PatchConvey("TestMockXXX", t, func() {
		Mock(good.Good).Return("yes").Build() // mock函数
		assert.Equal(t, bad.Bad(), "yes")
	})
}
