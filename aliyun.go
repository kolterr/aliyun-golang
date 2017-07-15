package aliyun

import (
	"fmt"
)

//Aliyun ...
type Aliyun struct {
	Account Account
}

//Account ...
type Account struct {
	AccessID     string
	AccessSecret string
}

//NewAliyun ...
func NewAliyun(a, p string) *Aliyun {
	o := &Aliyun{Account: Account{AccessID: a, AccessSecret: p}}
	return o
}

//CreateInstance ...  创建实例
func (o *Aliyun) CreateInstance() {

}

//DescribeInstances ...  查询实例列表
func (o *Aliyun) DescribeInstances(Reg string) {
	_url := "http://ecs.aliyuncs.com/?TimeStamp=2016-02-23T12:46:24Z&Format=XML&AccessKeyId=testid&Action=DescribeRegions&SignatureMethod=HMAC-SHA1&SignatureNonce=3ee8c1b8-83d3-44af-a94f-4e0ad82fd6cf&Version=2014-05-26&SignatureVersion=1.0"
	param := o.makeURLQuery(_url)
	fmt.Println(param)
	fmt.Println(o.makeURLEncode("GET", param))
	s := o.hmacSha1(o.makeURLEncode("GET", param))
	fmt.Println(s)
	// com, randStr, t := o.makeCommon()
	// fmt.Println(com)
	// _url := fmt.Sprintf(InstanceURL, Reg) + com //生成对应的请求连接
	// param := o.makeURLQuery(_url)
	// fmt.Println(param)
	// s := o.hmacSha1(o.makeURLEncode("GET", param))
	// fmt.Println(s)                                                                                   //对query部分签名
	// url := fmt.Sprintf(InstanceURL, Reg) + fmt.Sprintf(CommonURL, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
	// fmt.Println(url)
	// _, err := o.Curl(url, "GET", "") //触发请求
	// return err
}

//StartInstance ... 启动实例
func (o *Aliyun) StartInstance() {

}

//DeleteInstance ... 删除实例
func (o *Aliyun) DeleteInstance() {

}

//ModifyInstanceAutoReleaseTime ...设置实例释放时间
func (o *Aliyun) ModifyInstanceAutoReleaseTime() {

}
