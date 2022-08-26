package src

import "testing"

// 成功用例
func TestHello(t *testing.T) {
	r := Hello(1)
	if r != "hello" {
		t.Errorf("测试不通过.", r)
	}
}
