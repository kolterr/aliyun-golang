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
	a := NewAliyun("LTAI5CNiUdyqtzK7", "wP6ooitJXcDm6iYh3IiS2ZFgPZR82U")
	// ids := []string{"i-t4nimkw9we71vezgly9k"}
	// i, _ := json.Marshal(ids)
	// arg := make(map[string]string)
	// arg["InstanceIds"] = string(i)
	c := make(map[string]string)
	c["KeyPairName"] = "CF"
	// c["InternetMaxBandwidthOut"] = "5"
	res, err := a.DeleteInstance("i-2zecno0m7jzzffocosl7")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
