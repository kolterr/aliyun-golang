package aliyun

//CreateKeyPair ...
//Warning: each of aliyun accounts in diffrent region have less than 500 keys
func (o *Aliyun) CreateKeyPair(param CreateKeyPairs) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//DescribeKeyPairs ...
func (o *Aliyun) DescribeKeyPairs(param DescribeKeyPairss) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//DeleteKeyPairs ...
func (o *Aliyun) DeleteKeyPairs(param DeleteKeyPairss) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//ImportKeyPair ...
func (o *Aliyun) ImportKeyPair(param ImportKeyPairs) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//AttachKeyPair ...
func (o *Aliyun) AttachKeyPair(param AttachKeyPairs) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}

//DetachKeyPair ...
func (o *Aliyun) DetachKeyPair(param DetachKeyPairs) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}
