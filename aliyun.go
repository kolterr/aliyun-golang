package aliyun

//Aliyun ...
type Aliyun struct {
	Account Account
	Format  string //the format of response
}

//Account ...
type Account struct {
	AccessID     string
	AccessSecret string
}

//NewAliyun ...
func NewAliyun(a, p, f string) *Aliyun {
	o := &Aliyun{Account: Account{AccessID: a, AccessSecret: p}, Format: f}
	return o
}
