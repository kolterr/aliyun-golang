package aliyun

import "testing"
import "fmt"

func Test1(t *testing.T) {
	//debian_8_08_64_40G_alibase_20170625.vhd
	a := NewAliyun("xxx", "xxx")
	arr := []string{"CF2"}
	res, err := a.DeleteKeyPairs("ap-southeast-1", arr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
