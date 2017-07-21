package aliyun

import (
	"errors"
	"fmt"
)

//CreateInstance ...  which's status is stoped after instance be created,and it didn't hava publicIPaddress
//you can use AllocatePublicIPAddress api to deliver;
func (o *Aliyun) CreateInstance(Reg, ImageID, InstanceType string, args ...map[string]string) (string, error) {
	if Reg != "" && ImageID != "" && InstanceType != "" {
		str := o.makeMapArgs(args)
		_url := fmt.Sprintf(CreateURL, Reg, ImageID, InstanceType) + str
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t)
		res, err := o.Curl(url, "GET", "")
		return res, err
	}
	return "", errors.New("must need instanceID")
}

//DescribeInstances ...  query list of instance
func (o *Aliyun) DescribeInstances(Reg string, args ...map[string]string) (string, error) {
	str := ""
	if len(args) > 0 {
		str = o.makeMapArgs(args)
	}
	_url := fmt.Sprintf(InstanceURL, Reg) + str
	s, randStr, t := o.makeCommonURL(_url)
	url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//StartInstance ... start instance
func (o *Aliyun) StartInstance(id string) (string, error) {
	if id != "" {
		_url := fmt.Sprintf(StartURL, id)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t)
		res, err := o.Curl(url, "GET", "")
		return res, err
	}
	return "", errors.New("must need instanceID")
}

//DeleteInstance ... delete instance
//Warning:the status of instance must be stoped  before delete instance
func (o *Aliyun) DeleteInstance(id string) (string, error) {
	if id != "" {
		_url := fmt.Sprintf(DeleteURL, id)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t)
		res, err := o.Curl(url, "GET", "")
		return res, err
	}
	return "", errors.New("must need instanceID")
}

//StopInstance ...stop instance
func (o *Aliyun) StopInstance(id string) (string, error) {
	if id != "" {
		_url := fmt.Sprintf(StopInstance, id)
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t)
		res, err := o.Curl(url, "GET", "")
		return res, err
	}
	return "", errors.New("must need instanceID")
}

//ModifyInstanceAutoReleaseTime ...
//Set the instance release time;Format  yyyy-MM-ddTHH:mm:ssZ UTC time; At least half an hour after the current time; up to three years beyond the current time
func (o *Aliyun) ModifyInstanceAutoReleaseTime(id string, args ...string) (string, error) {
	if id != "" {
		_url := fmt.Sprintf(ModifyInstanceAutoReleaseTime, id)
		if args != nil {
			_url += "&AutoReleaseTime=" + args[0]
		}
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t)
		res, err := o.Curl(url, "GET", "")
		return res, err
	}
	return "", errors.New("must need instanceID")
}

//InstanceURLStatus ... query the status of instance
func (o *Aliyun) InstanceURLStatus(Reg string, args ...map[string]interface{}) (string, error) {
	if Reg != "" {
		_url := fmt.Sprintf(InstanceURLStatus, Reg)
		if len(args) > 0 {
			//TODO

		}
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t)
		res, err := o.Curl(url, "GET", "")
		return res, err
	}
	return "", errors.New("must need RegonID")
}

//RebootInstance ... restart instance
//the args have two options,true which means force stop the instance and another is normaly stop
func (o *Aliyun) RebootInstance(id string, args ...string) (string, error) {
	if id != "" {
		_url := fmt.Sprintf(RestartInstanceURL, id)
		if args != nil {
			_url += "&ForceStop=" + args[0]
		}
		s, randStr, t := o.makeCommonURL(_url)
		url := _url + fmt.Sprintf(CommonURL, o.Format, s, randStr, o.Account.AccessID, t)
		res, err := o.Curl(url, "GET", "")
		return res, err
	}
	return "", errors.New("must need instanceID")
}
