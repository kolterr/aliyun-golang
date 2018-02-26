package aliyun

//DescribeImages ... query images list
func (o *Aliyun) DescribeImages(param DescribeImagess) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}
