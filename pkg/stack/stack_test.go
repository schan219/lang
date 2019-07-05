package parser_test


import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var TEST_STRS = []string{"1","2","3","4"}
var TEST_NUMS = []int{1,2,3,4,-1}

func TestTop(t *testing.T) {
    assert := assert.New(t)
    
    s := stack.NewStack()

    for str,_ := range TEST_STRS {
        s.Push(str)
        assert.Equal(s.Top(), str)
    }

    for str,_ := range TEST_STRS {
        s.Push(str)
        assert.Equal(s.Top(), str)
    }

}
func TestPush(t *testing.T) {
}

func TestPop(t *testing.T) {

}
