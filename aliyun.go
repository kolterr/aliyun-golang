package aliyun

//Aliyun ...
type Aliyun struct {
	Account Account
}

//Account ...
type Account struct {
	AccessID     string
	AccessSecret string
}

//NewAliyun ...
func NewAliyun(a, p string) *Aliyun {
	o := &Aliyun{Account: Account{AccessID: a, AccessSecret: p}}
	return o
}
