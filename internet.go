package aliyun

import (
	"errors"
	"fmt"
)

//AllocatePublicIPAddress ...
//Warning:if you want to deliver ip to instance,you must sure you haven't set InternetMaxBandwidthOut;otherwise you will get a unkownErr
func (o *Aliyun) AllocatePublicIPAddress(id string) (string, error) {
	if id != "" {
		_url := fmt.Sprintf(AllocatePublicIPAddressURL, id)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t)
		res, err := o.Curl(url, "GET", "")
		return res, err
	}
	return "", errors.New("must need instanceID")
}
