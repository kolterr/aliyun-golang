package aliyun

//Aliyun url ...
var (
	CommonURL       = "&Format=%s&Version=2014-05-26&Signature=%s&SignatureMethod=HMAC-SHA1&SignatureNonce=%v&SignatureVersion=1.0&AccessKeyId=%s&Timestamp=%v"
	CommonURLNoSign = "&Format=%s&Version=2014-05-26&SignatureMethod=HMAC-SHA1&SignatureNonce=%v&SignatureVersion=1.0&AccessKeyId=%s&Timestamp=%v"

	CreateURL                     = "https://ecs.aliyuncs.com/?Action=CreateInstance&RegionId=%s&ImageId=%s&InstanceType=%s"
	StartURL                      = "https://ecs.aliyuncs.com/?Action=StartInstance&InstanceId=%s"
	InstanceURLStatus             = "https://ecs.aliyuncs.com/?Action=DescribeInstanceStatus&RegionId=%s"
	DeleteURL                     = "https://ecs.aliyuncs.com/?Action=DeleteInstance&InstanceId=%s"
	InstanceURL                   = "https://ecs.aliyuncs.com/?Action=DescribeInstances&RegionId=%s"
	RestartInstanceURL            = "https://ecs.aliyuncs.com/?Action=RebootInstance&InstanceId=%s"
	ModifyInstanceAutoReleaseTime = "https://ecs.aliyuncs.com/?Action=ModifyInstanceAutoReleaseTime&InstanceId=%s"
	StopInstance                  = "https://ecs.aliyuncs.com/?Action=StopInstance&InstanceId=%s"

	//镜像
	DescribeImages = "https://ecs.aliyuncs.com/?Action=DescribeImages&RegionId=%s"

	//密匙对
	DescribeKeyPairs = "https://ecs.aliyuncs.com/?action=DescribeKeyPairs&RegionId=%s"
	CreateKeyPair    = "https://ecs.aliyuncs.com/?action=CreateKeyPair&RegionId=%s&KeyPairName=%s"
	DeleteKeyPairs   = "https://ecs.aliyuncs.com/?action=DeleteKeyPairs&RegionId=%s&KeyPairNames=%s"
	ImportKeyPair    = "https://ecs.aliyuncs.com/?action=ImportKeyPair&RegionId=%s&PublicKeyBody=%s&KeyPairName=%s"

	//网络
	AllocatePublicIPAddressURL = "https://ecs.aliyuncs.com/?Action=AllocatePublicIpAddress&InstanceId=%s"

	//地域
	DescribeRegions = "https://ecs.aliyuncs.com/?Action=DescribeRegions"
	DescribeZones   = "https://ecs.aliyuncs.com/?Action=DescribeZones&RegionId=%s"
	//监控
	DescribeInstanceMonitorData = "https://ecs.aliyuncs.com/?Action=DescribeInstanceMonitorData&InstanceId=%s&StartTime=%s&EndTime=%s"
	DescribeDiskMonitorData     = "https://ecs.aliyuncs.com/?Action=DescribeDiskMonitorData&DiskId=%s&StartTime=%s&EndTime=%s"
	//磁盘
	DescribeDisks = "https://ecs.aliyuncs.com/?action=DescribeDisks&RegionId=%s"
)

const host = "https://ecs.aliyuncs.com/?"

//KeyPairs ...
type KeyPairs struct {
	KeyPairName string `json:"KeyPairName"`
}

//InstanceID ...
type InstanceID struct {
	ID string
}

type checkinstance struct {
	Instances struct {
		Instance []struct {
			InstanceID      string
			PublicIPAddress []string
		}
	}
}

type item struct {
	IPAddress []string
}

type createBack struct {
	RequestID  string `json:"RequestId"`
	InstanceID string `json:"InstanceId"`
}

//CreateInstances ...
type CreateInstances struct {
	Action                      string `json:"Action"`
	RegionID                    string `json:"RegionId"`
	ZoneID                      string `json:"ZoneId"`
	ImageID                     string `json:"ImageId"`
	InstanceType                string `json:"InstanceType"`
	SecurityGroupID             string `json:"SecurityGroupId"`
	InstanceName                string `json:"InstanceName"`
	Description                 string `json:"Description"`
	InternetChargeType          string `json:"InternetChargeType"`
	InternetMaxBandwidthIn      string `json:"InternetMaxBandwidthIn"`
	InternetMaxBandwidthOut     string `json:"InternetMaxBandwidthOut"`
	HostName                    string `json:"HostName"`
	Password                    string `json:"Password"`
	IoOptimized                 string `json:"IoOptimized"`
	SystemDiskCategory          string `json:"SystemDisk.Category"`
	SystemDiskSize              int    `json:"SystemDisk.Size"`
	SystemDiskDiskName          string `json:"SystemDisk.DiskName"`
	SystemDiskDescription       string `json:"SystemDisk.Description"`
	DataDisknSize               int    `json:"DataDisk.n.Size"`
	DataDisknCategory           string `json:"DataDisk.n.Category"`
	DataDisknSnapshotID         string `json:"DataDisk.n.SnapshotId"`
	DataDisknDiskName           string `json:"DataDisk.n.DiskName"`
	DataDisknDescription        string `json:"DataDisk.n.Description"`
	DataDisknDeleteWithInstance string `json:"DataDisk.n.DeleteWithInstance"`
	VSwitchID                   string `json:"VSwitchId"`
	PrivateIPAddress            string `json:"PrivateIpAddress"`
	InstanceChargeType          string `json:"InstanceChargeType"`
	Period                      int    `json:"Period"`
	AutoRenew                   string `json:"AutoRenew"`
	AutoRenewPeriod             int    `json:"AutoRenewPeriod"`
	UserData                    string `json:"UserData"`
	ClientToken                 string `json:"ClientToken"`
	KeyPairName                 string `json:"KeyPairName"`
	DeploymentSetID             string `json:"DeploymentSetId"`
	RAMRoleName                 string `json:"RamRoleName"`
	SecurityEnhancementStrategy string `json:"SecurityEnhancementStrategy"`
	TagnKey                     string `json:"Tag.n.Key"`
	TagnValue                   string `json:"Tag.n.Value"`
}

//DescribeInstance ...
type DescribeInstance struct {
	Action              string   `json:"Action"`
	RegionID            string   `json:"RegionId"`
	VpcID               string   `json:"VpcId"`
	VSwitchID           string   `json:"VSwitchId"`
	ZoneID              string   `json:"ZoneId"`
	InstanceIds         []string `json:"InstanceIds"`
	InstanceType        string   `json:"InstanceType"`
	InstanceTypeFamily  string   `json:"InstanceTypeFamily"`
	InstanceNetworkType string   `json:"InstanceNetworkType"`
	PrivateIPAddresses  []string `json:"PrivateIpAddresses"`
	InnerIPAddresses    []string `json:"InnerIpAddresses"`
	PublicIPAddresses   []string `json:"PublicIpAddresses"`
	SecurityGroupID     string   `json:"SecurityGroupId"`
	InstanceChargeType  string   `json:"InstanceChargeType"`
	InternetChargeType  string   `json:"InternetChargeType"`
	InstanceName        string   `json:"InstanceName"`
	ImageID             string   `json:"ImageId"`
	Status              string   `json:"Status"`
	IoOptimized         string   `json:"IoOptimized"`
	TagnKey             string   `json:"Tag.n.Key"`
	TagnValue           string   `json:"Tag.n.Value"`
	PageNumber          int      `json:"PageNumber"`
	PageSize            int      `json:"PageSize"`
}

//StopInstances ...
type StopInstances struct {
	Action      string `json:"Action"`
	InstanceID  string `json:"InstanceId"`
	ForceStop   bool   `json:"ForceStop"`
	ConfirmStop bool   `json:"ConfirmStop"`
}

//StartInstances ...
type StartInstances struct {
	Action     string `json:"Action"`
	InstanceID string `json:"InstanceId"`
}

//RebootInstances ...
type RebootInstances struct {
	Action     string `json:"Action"`
	InstanceID string `json:"InstanceId"`
	ForceStop  bool   `json:"ForceStop"`
}

//DescribeInstanceStatus ...
type DescribeInstanceStatus struct {
	Action     string `json:"Action"`
	RegionID   string `json:"RegionId"`
	ZoneID     string `json:"ZoneId"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
}

//ModifyInstanceAutoReleaseTimes ...
type ModifyInstanceAutoReleaseTimes struct {
	Action          string `json:"Action"`
	InstanceID      string `json:"InstanceId"`
	AutoReleaseTime string `json:"AutoReleaseTime"`
}

//DeleteInstances ...
type DeleteInstances struct {
	Action     string `json:"Action"`
	InstanceID string `json:"InstanceId"`
}

//AllocatePublicIPAddresss ...
type AllocatePublicIPAddresss struct {
	Action     string `json:"Action"`
	InstanceID string `json:"InstanceId"`
}

//DescribeRegionss ...
type DescribeRegionss struct {
	Action string `json:"Action"`
}

//DescribeZoness ...
type DescribeZoness struct {
	Action   string `json:"Action"`
	RegionID string `json:"RegionId"`
}

//DescribeInstanceMonitorDatas ...
type DescribeInstanceMonitorDatas struct {
	Action     string `json:"Action"`
	InstanceID string `json:"InstanceId"`
	StartTime  string `json:"StartTime"`
	EndTime    string `json:"EndTime"`
	Period     int    `json:"Period"`
}

//DescribeDiskMonitorDatas ...
type DescribeDiskMonitorDatas struct {
	Action    string `json:"Action"`
	DiskID    string `json:"DiskId"`
	StartTime string `json:"StartTime"`
	EndTime   string `json:"EndTime"`
	Period    int    `json:"Period"`
}

//CreateKeyPairs ...
type CreateKeyPairs struct {
	Action      string `json:"Action"`
	RegionID    string `json:"RegionId"`
	KeyPairName string `json:"KeyPairName"`
}

//ImportKeyPairs ...
type ImportKeyPairs struct {
	Action        string `json:"Action"`
	RegionID      string `json:"RegionId"`
	KeyPairName   string `json:"KeyPairName"`
	PublicKeyBody string `json:"PublicKeyBody"`
}

//DescribeKeyPairss ...
type DescribeKeyPairss struct {
	Action             string `json:"Action"`
	RegionID           string `json:"RegionId"`
	KeyPairFingerPrint string `json:"KeyPairFingerPrint"`
	KeyPairName        string `json:"KeyPairName"`
	PageNumber         int    `json:"PageNumber"`
	PageSize           int    `json:"PageSize"`
}

//AttachKeyPairs ...
type AttachKeyPairs struct {
	Action      string   `json:"Action"`
	RegionID    string   `json:"RegionId"`
	KeyPairName string   `json:"KeyPairName"`
	InstanceIds []string `json:"InstanceIds"`
}

//DetachKeyPair ...
type DetachKeyPairs struct {
	Action      string   `json:"Action"`
	RegionID    string   `json:"RegionId"`
	KeyPairName string   `json:"KeyPairName"`
	InstanceIds []string `json:"InstanceIds"`
}

//DeleteKeyPairss ...
type DeleteKeyPairss struct {
	Action      string   `json:"Action"`
	RegionID    string   `json:"RegionId"`
	KeyPairName []string `json:"KeyPairName"`
}

//DescribeImagess ...
type DescribeImagess struct {
	Action          string `json:"Action"`
	RegionID        string `json:"RegionId"`
	ImageID         string `json:"ImageId"`
	Status          string `json:"Status"`
	SnapshotID      string `json:"SnapshotId"`
	ImageName       string `json:"ImageName"`
	ImageOwnerAlias string `json:"ImageOwnerAlias"`
	Usage           string `json:"Usage"`
	TagnKey         string `json:"Tag.n.Key"`
	TagnValue       string `json:"Tag.n.Value"`
	PageNumber      int    `json:"PageNumber"`
	PageSize        int    `json:"PageSize"`
}
