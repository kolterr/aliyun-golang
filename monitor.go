package aliyun

//DescribeInstanceMonitorData ...
func (o *Aliyun) DescribeInstanceMonitorData(param DescribeInstanceMonitorDatas) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//DescribeDiskMonitorData ...
func (o *Aliyun) DescribeDiskMonitorData(param DescribeDiskMonitorDatas) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}
