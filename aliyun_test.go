package aliyun

import "testing"
import "fmt"

func Test1(t *testing.T) {
	a := NewAliyun("xxx", "xxx")
	res, err := a.DeleteInstance("i-t4nh4wencasuwj1lkshm")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
