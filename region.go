package aliyun

//DescribeRegions ...
func (o *Aliyun) DescribeRegions(param DescribeRegionss) (string, error) {
	regs, _ := o.Struct2Map(param)
	params := o.validator(regs)
	url := o.makeURLInternal(params)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//DescribeZones ...
func (o *Aliyun) DescribeZones(param DescribeZoness) (string, error) {
	regs, _ := o.Struct2Map(param)
	params := o.validator(regs)
	url := o.makeURLInternal(params)
	res, err := o.Curl(url, "GET", "")
	return res, err
}
