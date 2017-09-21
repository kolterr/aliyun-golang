package aliyun

//AllocatePublicIPAddress ...
//Warning:if you want to deliver ip to instance,you must sure you haven't set InternetMaxBandwidthOut;otherwise you will get a unkownErr
func (o *Aliyun) AllocatePublicIPAddress(param AllocatePublicIPAddresss) (string, error) {
	regs, _ := o.Struct2Map(param)
	params := o.validator(regs)
	url := o.makeURLInternal(params)
	res, err := o.Curl(url, "GET", "")
	return res, err
}
