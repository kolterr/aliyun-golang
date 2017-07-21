package aliyun

import (
	"errors"
	"fmt"
)

//DescribeInstanceMonitorData ...
func (o *Aliyun) DescribeInstanceMonitorData(id, startTime, endTime string) (string, error) {
	if startTime != "" && endTime != "" && id != "" {
		_url := fmt.Sprintf(DescribeInstanceMonitorData, id, startTime, endTime)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t)
		res, err := o.Curl(url, "GET", "")
		return res, err
	}
	return "", errors.New("must need regionID")
}

//DescribeDiskMonitorData ...
func (o *Aliyun) DescribeDiskMonitorData(disk, startTime, endTime string) (string, error) {
	if startTime != "" && endTime != "" && disk != "" {
		_url := fmt.Sprintf(DescribeDiskMonitorData, disk, startTime, endTime)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t)
		res, err := o.Curl(url, "GET", "")
		return res, err
	}
	return "", errors.New("must need regionID")
}
