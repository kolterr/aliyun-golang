package aliyun

import (
	"encoding/json"
	"errors"
	"fmt"
)

//CreateKeyPair ...
//每个用户账号每个 Region 不能超过 500 个
func (o *Aliyun) CreateKeyPair(Reg, name string) (string, error) {
	if name != "" && Reg != "" {
		_url := fmt.Sprintf(CreateKeyPair, Reg, name)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                      //触发请求
		return res, err
	}
	return "", errors.New("must need a KeyPairName")
}

//DescribeKeyPairs ...
func (o *Aliyun) DescribeKeyPairs(Reg string) (string, error) {
	if Reg != "" {
		_url := fmt.Sprintf(DescribeKeyPairs, Reg)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                      //触发请求
		return res, err
	}
	return "", errors.New("must need regID")
}

//DeleteKeyPairs ...删除密匙对
func (o *Aliyun) DeleteKeyPairs(Reg string, names []string) (string, error) {
	if len(names) != 0 && Reg != "" {
		arr := []KeyPairs{}
		for i := 0; i < len(names); i++ {
			arr = append(arr, KeyPairs{KeyPairName: names[i]})
		}
		json, _ := json.Marshal(arr)
		fmt.Println(string(json))
		_url := fmt.Sprintf(DeleteKeyPairs, Reg, o.percentEncode(string(json)))
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		fmt.Println(url)
		res, err := o.Curl(url, "GET", "") //触发请求
		return res, err
	}
	return "", errors.New("must need a KeyPairName")
}

//ImportKeyPair ... 导入密匙对
func (o *Aliyun) ImportKeyPair(Reg, key, name string) (string, error) {
	if Reg != "" && key != "" && name != "" {
		_url := fmt.Sprintf(ImportKeyPair, Reg, o.percentEncode(key), name)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                      //触发请求
		return res, err
	}
	return "", errors.New("more information needed")
}
