package aliyun

import (
	"errors"
	"fmt"
)

//AllocatePublicIPAddress ... 需要设置 InternetMaxBandwidthOut ：公网出带宽最大值 [0, 100] 不然报未知错误
//Warning:if you want to deliver ip to instance,you must sure you haven't set InternetMaxBandwidthOut;otherwise you will get a unkownErr
func (o *Aliyun) AllocatePublicIPAddress(id string) (string, error) {
	if id != "" {
		_url := fmt.Sprintf(AllocatePublicIPAddressURL, id)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                                //触发请求
		return res, err
	}
	return "", errors.New("must need instanceID")
}
