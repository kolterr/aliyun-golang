package aliyun

//DescribeInstanceMonitorData ...
func (o *Aliyun) DescribeInstanceMonitorData(param DescribeInstanceMonitorDatas) (string, error) {
	regs, _ := o.Struct2Map(param)
	params := o.validator(regs)
	url := o.makeURLInternal(params)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//DescribeDiskMonitorData ...
func (o *Aliyun) DescribeDiskMonitorData(param DescribeDiskMonitorDatas) (string, error) {
	regs, _ := o.Struct2Map(param)
	params := o.validator(regs)
	url := o.makeURLInternal(params)
	res, err := o.Curl(url, "GET", "")
	return res, err
}
