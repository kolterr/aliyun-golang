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
	a := NewAliyun("xxx", "xxx", "JSON")
	// ids := []string{"i-t4nimkw9we71vezgly9k"}
	// i, _ := json.Marshal(ids)
	// arg := make(map[string]string)
	// arg["InstanceIds"] = string(i)
	c := make(map[string]string)
	c["KeyPairName"] = "CF"
	// c["InternetMaxBandwidthOut"] = "5"
	res, err := a.DescribeDiskMonitorData("d-2zea456im5lkgeph7kj9", "2017-07-21T12:07:00Z", "2017-07-21T20:07:00Z")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
