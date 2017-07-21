package aliyun

import (
	"encoding/json"
	"errors"
	"fmt"
)

//CreateKeyPair ...
//Warning: each of aliyun accounts in diffrent region have less than 500 keys
func (o *Aliyun) CreateKeyPair(Reg, name string) (string, error) {
	if name != "" && Reg != "" {
		_url := fmt.Sprintf(CreateKeyPair, Reg, name)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                                //触发请求
		return res, err
	}
	return "", errors.New("must need a KeyPairName")
}

//DescribeKeyPairs ...
func (o *Aliyun) DescribeKeyPairs(Reg string) (string, error) {
	if Reg != "" {
		_url := fmt.Sprintf(DescribeKeyPairs, Reg)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                                //触发请求
		return res, err
	}
	return "", errors.New("must need regID")
}

//DeleteKeyPairs ...
func (o *Aliyun) DeleteKeyPairs(Reg string, names []string) (string, error) {
	if len(names) != 0 && Reg != "" {
		arr := []string{}
		for i := 0; i < len(names); i++ {
			arr = append(arr, names[i])
		}
		json, _ := json.Marshal(arr)
		_url := fmt.Sprintf(DeleteKeyPairs, Reg, o.percentEncode(string(json)))
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                                //触发请求
		return res, err
	}
	return "", errors.New("must need a KeyPairName")
}

//ImportKeyPair ...
func (o *Aliyun) ImportKeyPair(Reg, key, name string) (string, error) {
	if Reg != "" && key != "" && name != "" {
		_url := fmt.Sprintf(ImportKeyPair, Reg, o.percentEncode(key), name)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                                //触发请求
		return res, err
	}
	return "", errors.New("more information needed")
}
