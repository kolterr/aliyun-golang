package aliyun

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	//debian_8_08_64_40G_alibase_20170625.vhd
	a := NewAliyun("", "", "JSON")
	// pips := []string{"123.56.217.122"}
	p := DescribeInstance{Action:"DescribeInstanceStatus",RegionID:"cn-qingdao"}
	res, err := a.DescribeInstances(p)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
