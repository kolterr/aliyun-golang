package aliyun

//DescribeRegions ...
func (o *Aliyun) DescribeRegions(param DescribeRegionss) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//DescribeZones ...
func (o *Aliyun) DescribeZones(param DescribeZoness) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}
