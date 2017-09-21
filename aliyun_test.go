package aliyun

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	type allocateIP struct {
		RequestID string
		IPAddress string
	}
	//debian_8_08_64_40G_alibase_20170625.vhd
	a := NewAliyun("LTAI5CNiUdyqtzK7", "wP6ooitJXcDm6iYh3IiS2ZFgPZR82U", "JSON")
	// pips := []string{"123.56.217.122"}
	res, err := a.StopInstance("i-2zec4vask3fyg1mhjx2a")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
