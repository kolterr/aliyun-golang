package aliyun

//CreateInstance ...  which's status is stoped after instance be created,and it didn't hava publicIPaddress
//you can use AllocatePublicIPAddress api to deliver;
func (o *Aliyun) CreateInstance(param CreateInstances) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//DescribeInstances ...  query list of instance
func (o *Aliyun) DescribeInstances(param DescribeInstance) (string, error) {
	o.do(param)
	res, err := o.Curl(o.url, "GET", "")
	return res, err
}

//StartInstance ... start instance
func (o *Aliyun) StartInstance(param StartInstances) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//DeleteInstance ... delete instance
//Warning:the status of instance must be stoped  before delete instance
func (o *Aliyun) DeleteInstance(param DeleteInstances) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//StopInstance ...stop instance
func (o *Aliyun) StopInstance(param StopInstances) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//ModifyInstanceAutoReleaseTime ...
//Set the instance release time;Format  yyyy-MM-ddTHH:mm:ssZ UTC time; At least half an hour after the current time; up to three years beyond the current time
func (o *Aliyun) ModifyInstanceAutoReleaseTime(param ModifyInstanceAutoReleaseTimes) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//InstanceURLStatus ... query the status of instance
func (o *Aliyun) InstanceURLStatus(param DescribeInstanceStatus) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//RebootInstance ... restart instance
//the args have two options,true which means force stop the instance and another is normaly stop
func (o *Aliyun) RebootInstance(param RebootInstances) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}
