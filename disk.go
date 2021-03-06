package aliyun

import (
	"errors"
	"fmt"
)

//DescribeDisks ...
func (o *Aliyun) DescribeDisks(Reg string) (string, error) {
	if Reg != "" {
		_url := fmt.Sprintf(DescribeDisks, Reg)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                                //触发请求
		return res, err
	}
	return "", errors.New("Incomplete parameter")
}
