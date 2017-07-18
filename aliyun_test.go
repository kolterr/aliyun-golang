package aliyun

import "testing"
import "fmt"

func Test1(t *testing.T) {
	//debian_8_08_64_40G_alibase_20170625.vhd
	a := NewAliyun("XXX", "XXX")
	// ids := []string{"i-t4nimkw9we71vezgly9k"}
	// i, _ := json.Marshal(ids)
	// arg := make(map[string]string)
	// arg["InstanceIds"] = string(i)
	res, err := a.DeleteKeyPairs("ap-southeast-1", []string{"CF2"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
