package aliyun

import "testing"

func Test1(t *testing.T) {
	a := NewAliyun("LTAI5CNiUdyqtzK7", "testsecret")
	a.DescribeInstances("ap-southeast-1")
}
