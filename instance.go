package aliyun

import (
	"errors"
	"fmt"
)

//CreateInstance ...  创建实例  创建后为停止状态
func (o *Aliyun) CreateInstance(Reg, ImageID, InstanceType string, args ...map[string]string) (string, error) {
	if Reg != "" && ImageID != "" && InstanceType != "" {
		str := o.makeMapArgs(args)
		_url := fmt.Sprintf(CreateURL, Reg, ImageID, InstanceType) + str
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                      //触发请求
		return res, err
	}
	return "", errors.New("must need instanceID")
}

//DescribeInstances ...  查询实例列表
func (o *Aliyun) DescribeInstances(Reg string, args ...map[string]string) (string, error) {
	str := o.makeMapArgs(args)
	_url := fmt.Sprintf(InstanceURL, Reg) + str
	s, randStr, t := o.makeCommonURL(_url)                                  //对query部分签名
	url := _url + fmt.Sprintf(CommonURL, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
	res, err := o.Curl(url, "GET", "")                                      //触发请求
	return res, err
}

//StartInstance ... 启动实例
func (o *Aliyun) StartInstance(id string) (string, error) {
	if id != "" {
		_url := fmt.Sprintf(StartURL, id)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                      //触发请求
		return res, err
	}
	return "", errors.New("must need instanceID")
}

//DeleteInstance ... 删除实例
//实例状态必须为 Stopped，才可以进行删除操作。删除后，实例的状态为 Deleted，表示资源已释放，删除完成。
func (o *Aliyun) DeleteInstance(id string) (string, error) {
	if id != "" {
		_url := fmt.Sprintf(DeleteURL, id)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                      //触发请求
		return res, err
	}
	return "", errors.New("must need instanceID")
}

//StopInstance ...停止实例
func (o *Aliyun) StopInstance(id string) (string, error) {
	if id != "" {
		_url := fmt.Sprintf(StopInstance, id)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                      //触发请求
		return res, err
	}
	return "", errors.New("must need instanceID")
}

//ModifyInstanceAutoReleaseTime ...设置实例释放时间
//自动释放时间。按照 ISO8601 标准表示，并需要使用 UTC 时间。 格式为：yyyy-MM-ddTHH:mm:ssZ。 如果秒不是 00，则自动取为当前分钟开始时。最少在当前时间之后半小时；最多不能超过当前时间起三年。
func (o *Aliyun) ModifyInstanceAutoReleaseTime(id string, args ...string) (string, error) {
	if id != "" {
		_url := fmt.Sprintf(ModifyInstanceAutoReleaseTime, id)
		if args != nil {
			_url += "&AutoReleaseTime=" + args[0]
		}
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                      //触发请求
		return res, err
	}
	return "", errors.New("must need instanceID")
}

//InstanceURLStatus ...查询实例状态
func (o *Aliyun) InstanceURLStatus(Reg string, args ...map[string]interface{}) (string, error) {
	if Reg != "" {
		_url := fmt.Sprintf(InstanceURLStatus, Reg)
		if len(args) > 0 {
			//TODO

		}
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                      //触发请求
		return res, err
	}
	return "", errors.New("must need RegonID")
}

//RebootInstance ...重启实例 参数依次为 实例ID 、若为 false 则走正常关机流程;若为 true 则强制关机。 如果不指定，则默认值为 false。
func (o *Aliyun) RebootInstance(id string, args ...string) (string, error) {
	if id != "" {
		_url := fmt.Sprintf(RestartInstanceURL, id)
		if args != nil {
			_url += "&ForceStop=" + args[0]
		}
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, s, randStr, o.Account.AccessID, t) //生成对应的请求链接 带 signure
		res, err := o.Curl(url, "GET", "")                                      //触发请求
		return res, err
	}
	return "", errors.New("must need instanceID")
}
