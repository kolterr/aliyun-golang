package aliyun

//CreateKeyPair ...
//Warning: each of aliyun accounts in diffrent region have less than 500 keys
func (o *Aliyun) CreateKeyPair(param CreateKeyPairs) (string, error) {
	regs, _ := o.Struct2Map(param)
	params := o.validator(regs)
	url := o.makeURLInternal(params)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//DescribeKeyPairs ...
func (o *Aliyun) DescribeKeyPairs(param DescribeKeyPairss) (string, error) {
	regs, _ := o.Struct2Map(param)
	params := o.validator(regs)
	url := o.makeURLInternal(params)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//DeleteKeyPairs ...
func (o *Aliyun) DeleteKeyPairs(param DeleteKeyPairss) (string, error) {
	regs, _ := o.Struct2Map(param)
	params := o.validator(regs)
	url := o.makeURLInternal(params)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//ImportKeyPair ...
func (o *Aliyun) ImportKeyPair(param ImportKeyPairs) (string, error) {
	regs, _ := o.Struct2Map(param)
	params := o.validator(regs)
	url := o.makeURLInternal(params)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//AttachKeyPair ...
func (o *Aliyun) AttachKeyPair(param AttachKeyPairs) (string, error) {
	regs, _ := o.Struct2Map(param)
	params := o.validator(regs)
	url := o.makeURLInternal(params)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//DetachKeyPair ...
func (o *Aliyun) DetachKeyPair(param DetachKeyPairs) (string, error) {
	regs, _ := o.Struct2Map(param)
	params := o.validator(regs)
	url := o.makeURLInternal(params)
	res, err := o.Curl(url, "GET", "")
	return res, err
}
