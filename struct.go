package aliyun

//Aliyun url ...
var (
	CommonURL       = "&Format=JSON&Version=2014-05-26&Signature=%s&SignatureMethod=HMAC-SHA1&SignatureNonce=%v&SignatureVersion=1.0&AccessKeyId=%s&Timestamp=%v"
	CommonURLNoSign = "&Format=JSON&Version=2014-05-26&SignatureMethod=HMAC-SHA1&SignatureNonce=%v&SignatureVersion=1.0&AccessKeyId=%s&Timestamp=%v"

	CreateURL                     = "https://ecs.aliyuncs.com/?Action=CreateInstance&RegionId=%s&ImageId=%s&InstanceType=%s"
	StartURL                      = "https://ecs.aliyuncs.com/?Action=StartInstance&InstanceId=%s"
	InstanceURLStatus             = "https://ecs.aliyuncs.com/?Action=DescribeInstanceStatus&RegionId=%s"
	DeleteURL                     = "https://ecs.aliyuncs.com/?Action=DeleteInstance&InstanceId=%s"
	InstanceURL                   = "https://ecs.aliyuncs.com/?Action=DescribeInstances&RegionId=%s"
	RestartInstanceURL            = "https://ecs.aliyuncs.com/?Action=RebootInstance&InstanceId=%s"
	ModifyInstanceAutoReleaseTime = "https://ecs.aliyuncs.com/?Action=ModifyInstanceAutoReleaseTime&InstanceId=%s"
	StopInstance                  = "https://ecs.aliyuncs.com/?Action=StopInstance&InstanceId=%s"
)
