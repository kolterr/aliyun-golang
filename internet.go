package aliyun

//AllocatePublicIPAddress ...
//Warning:if you want to deliver ip to instance,you must sure you haven't set InternetMaxBandwidthOut;otherwise you will get a unkownErr
func (o *Aliyun) AllocatePublicIPAddress(param AllocatePublicIPAddresss) (string, error) {
	o.do(param)
	res, err := o.Curl(url, "GET", "")
	return res, err
}
