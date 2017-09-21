package aliyun

//DescribeImages ... query images list
func (o *Aliyun) DescribeImages(param DescribeImagess) (string, error) {
	regs, _ := o.Struct2Map(param)
	params := o.validator(regs)
	url := o.makeURLInternal(params)
	res, err := o.Curl(url, "GET", "")
	return res, err
}
