package aliyun

import (
	"errors"
	"fmt"
)

//DescribeRegions ...
func (o *Aliyun) DescribeRegions() (string, error) {
	s, randStr, t := o.makeCommonURL(DescribeRegions)
	url := DescribeRegions + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//DescribeZones ...
func (o *Aliyun) DescribeZones(Reg string) (string, error) {
	if Reg != "" {
		_url := fmt.Sprintf(DescribeZones, Reg)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t)
		res, err := o.Curl(url, "GET", "")
		return res, err
	}
	return "", errors.New("must need regionID")
}
