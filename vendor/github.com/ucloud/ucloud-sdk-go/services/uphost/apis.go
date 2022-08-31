// Code is generated by ucloud-model, DO NOT EDIT IT.

package uphost

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UPHost API Schema

/*
CreatePHostParamDisks is request schema for complex param
*/
type CreatePHostParamDisks struct {

	// 裸金属机型参数->云盘代金券id。不适用于系统盘。请通过DescribeCoupon接口查询，或登录用户中心查看
	CouponId *string `required:"false"`

	// 裸金属机型参数->是否是系统盘。枚举值： True，是系统盘。 False，是数据盘（默认）。Disks数组中有且只能有一块盘是系统盘。
	IsBoot *string `required:"false"`

	// 裸金属机型参数->磁盘大小，单位GB，必须是10GB的整数倍。系统盘20-500GB，数据盘单块盘20-32000GB。
	Size *int `required:"false"`

	// 裸金属机型参数->磁盘类型：枚举值：CLOUD_RSSD
	Type *string `required:"false"`
}

// CreatePHostRequest is request schema for CreatePHost action
type CreatePHostRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"true"`

	// 计费模式，枚举值为：year, 按年付费； month,按月付费；默认为按月付费
	ChargeType *string `required:"false"`

	// 网络环境，可选千兆：1G ，万兆：10G， 默认1G。智能网卡可以选择25G。
	Cluster *string `required:"false"`

	// 【该字段已废弃，请谨慎使用】
	Count *int `required:"false" deprecated:"true"`

	// 代金券
	CouponId *string `required:"false"`

	//
	Disks []CreatePHostParamDisks `required:"false"`

	// ImageId，可以通过接口 [DescribePHostImage](api/uphost-api/describe_phost_image.html)获取
	ImageId *string `required:"true"`

	// 物理机名称，默认为phost
	Name *string `required:"false"`

	// 密码（密码需使用base64进行编码）
	Password *string `required:"true"`

	// 购买时长，1-10个月或1-10年；默认值为1。月付时，此参数传0，代表购买至月末，1代表整月。
	Quantity *string `required:"false"`

	// Raid配置，默认Raid10  支持:Raid0、Raid1、Raid5、Raid10，NoRaid
	Raid *string `required:"false"`

	// 物理机备注，默认为空
	Remark *string `required:"false"`

	// 防火墙ID，默认：Web推荐防火墙。如何查询SecurityGroupId请参见 [DescribeFirewall](api/unet-api/describe_firewall.html)。
	SecurityGroupId *string `required:"false"`

	// 子网ID，不填为默认，VPC2.0下需要填写此字段。
	SubnetId *string `required:"false"`

	// 业务组，默认为default
	Tag *string `required:"false"`

	// 物理机类型，默认为：db-2(基础型-SAS-V3)
	Type *string `required:"false"`

	// VPC ID，不填为默认，VPC2.0下需要填写此字段。
	VPCId *string `required:"false"`

	// 指定内网ip创建
	VpcIp *string `required:"false"`
}

// CreatePHostResponse is response schema for CreatePHost action
type CreatePHostResponse struct {
	response.CommonBase

	// PHost的资源ID数组
	PHostId []string
}

// NewCreatePHostRequest will create request of CreatePHost action.
func (c *UPHostClient) NewCreatePHostRequest() *CreatePHostRequest {
	req := &CreatePHostRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreatePHost

指定数据中心，根据资源使用量创建指定数量的UPHost物理云主机实例。
*/
func (c *UPHostClient) CreatePHost(req *CreatePHostRequest) (*CreatePHostResponse, error) {
	var err error
	var res CreatePHostResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreatePHost", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CreatePHostImageRequest is request schema for CreatePHostImage action
type CreatePHostImageRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"true"`

	// 镜像描述
	ImageDescription *string `required:"false"`

	// 镜像名称
	ImageName *string `required:"true"`

	// UPHost实例ID
	PHostId *string `required:"true"`
}

// CreatePHostImageResponse is response schema for CreatePHostImage action
type CreatePHostImageResponse struct {
	response.CommonBase

	// 镜像ID
	ImageId string
}

// NewCreatePHostImageRequest will create request of CreatePHostImage action.
func (c *UPHostClient) NewCreatePHostImageRequest() *CreatePHostImageRequest {
	req := &CreatePHostImageRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreatePHostImage

创建裸金属2.0用户自定义镜像
*/
func (c *UPHostClient) CreatePHostImage(req *CreatePHostImageRequest) (*CreatePHostImageResponse, error) {
	var err error
	var res CreatePHostImageResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreatePHostImage", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeBaremetalMachineTypeRequest is request schema for DescribeBaremetalMachineType action
type DescribeBaremetalMachineTypeRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"true"`

	// 具体机型。若不填写，则返回全部机型
	Type *string `required:"false"`
}

// DescribeBaremetalMachineTypeResponse is response schema for DescribeBaremetalMachineType action
type DescribeBaremetalMachineTypeResponse struct {
	response.CommonBase

	// 机型列表，模型：PHostCloudMachineTypeSet
	MachineTypes []PHostCloudMachineTypeSet
}

// NewDescribeBaremetalMachineTypeRequest will create request of DescribeBaremetalMachineType action.
func (c *UPHostClient) NewDescribeBaremetalMachineTypeRequest() *DescribeBaremetalMachineTypeRequest {
	req := &DescribeBaremetalMachineTypeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeBaremetalMachineType

获取裸金属机型的详细描述信息
*/
func (c *UPHostClient) DescribeBaremetalMachineType(req *DescribeBaremetalMachineTypeRequest) (*DescribeBaremetalMachineTypeResponse, error) {
	var err error
	var res DescribeBaremetalMachineTypeResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeBaremetalMachineType", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribePHostRequest is request schema for DescribePHost action
type DescribePHostRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"false"`

	// 返回数据长度，默认为20
	Limit *int `required:"false"`

	// 数据偏移量，默认为0
	Offset *int `required:"false"`

	// PHost资源ID，若为空，则返回当前Region所有PHost。
	PHostId []string `required:"false"`

	// 要挂载的云盘id，过滤返回能被UDiskId挂载的云主机。目前主要针对rssd云盘使用
	UDiskIdForAttachment *string `required:"false"`

	// ULB使用参数，获取同VPC下机器信息。
	VPCId *string `required:"false"`
}

// DescribePHostResponse is response schema for DescribePHost action
type DescribePHostResponse struct {
	response.CommonBase

	// PHost资源列表，参见 PHostSet
	PHostSet []PHostSet

	// 满足条件的PHost总数
	TotalCount int
}

// NewDescribePHostRequest will create request of DescribePHost action.
func (c *UPHostClient) NewDescribePHostRequest() *DescribePHostRequest {
	req := &DescribePHostRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribePHost

获取物理机详细信息
*/
func (c *UPHostClient) DescribePHost(req *DescribePHostRequest) (*DescribePHostResponse, error) {
	var err error
	var res DescribePHostResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribePHost", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribePHostImageRequest is request schema for DescribePHostImage action
type DescribePHostImageRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"true"`

	// 镜像ID
	ImageId []string `required:"false"`

	// 镜像类别，枚举值，Base是基础镜像；Custom是自制镜像。
	ImageType *string `required:"false"`

	// 返回数据长度，默认为20
	Limit *int `required:"false"`

	// 机器型号，只支持当前zone的展示机型
	MachineType *string `required:"false"`

	// 数据偏移量，默认为0
	Offset *int `required:"false"`
}

// DescribePHostImageResponse is response schema for DescribePHostImage action
type DescribePHostImageResponse struct {
	response.CommonBase

	// 镜像列表 PHostImageSet
	ImageSet []PHostImageSet

	// 满足条件的镜像总数
	TotalCount int
}

// NewDescribePHostImageRequest will create request of DescribePHostImage action.
func (c *UPHostClient) NewDescribePHostImageRequest() *DescribePHostImageRequest {
	req := &DescribePHostImageRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribePHostImage

获取物理云主机镜像列表
*/
func (c *UPHostClient) DescribePHostImage(req *DescribePHostImageRequest) (*DescribePHostImageResponse, error) {
	var err error
	var res DescribePHostImageResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribePHostImage", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribePHostMachineTypeRequest is request schema for DescribePHostMachineType action
type DescribePHostMachineTypeRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"true"`

	// 具体机型。若不填写，则返回全部机型
	Type *string `required:"false"`
}

// DescribePHostMachineTypeResponse is response schema for DescribePHostMachineType action
type DescribePHostMachineTypeResponse struct {
	response.CommonBase

	// 机型列表，模型：PHostMachineTypeSet
	MachineTypes []PHostMachineTypeSet
}

// NewDescribePHostMachineTypeRequest will create request of DescribePHostMachineType action.
func (c *UPHostClient) NewDescribePHostMachineTypeRequest() *DescribePHostMachineTypeRequest {
	req := &DescribePHostMachineTypeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribePHostMachineType

获取物理云机型的详细描述信息
*/
func (c *UPHostClient) DescribePHostMachineType(req *DescribePHostMachineTypeRequest) (*DescribePHostMachineTypeResponse, error) {
	var err error
	var res DescribePHostMachineTypeResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribePHostMachineType", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribePHostTagsRequest is request schema for DescribePHostTags action
type DescribePHostTagsRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

}

// DescribePHostTagsResponse is response schema for DescribePHostTags action
type DescribePHostTagsResponse struct {
	response.CommonBase

	// 具体参见 PHostTagSet
	TagSet []PHostTagSet

	// Tag的个数
	TotalCount int
}

// NewDescribePHostTagsRequest will create request of DescribePHostTags action.
func (c *UPHostClient) NewDescribePHostTagsRequest() *DescribePHostTagsRequest {
	req := &DescribePHostTagsRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribePHostTags

获取物理机tag列表（业务组）
*/
func (c *UPHostClient) DescribePHostTags(req *DescribePHostTagsRequest) (*DescribePHostTagsResponse, error) {
	var err error
	var res DescribePHostTagsResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribePHostTags", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetPHostDiskUpgradePriceRequest is request schema for GetPHostDiskUpgradePrice action
type GetPHostDiskUpgradePriceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 磁盘大小，单位GB，必须是10GB的整数倍。系统盘20-500GB，数据盘单块盘20-32000GB。
	DiskSpace *int `required:"true"`

	// UPHost实例ID。
	PHostId *string `required:"true"`

	// 是否重装价格获取。复用此接口。扩容只能增加云盘大小。重装不限制。枚举值：true/false
	ReinstallTag *bool `required:"false"`

	// 磁盘 ID。获取扩容价格必填（只能扩不能减）；重装时候不需要填（根据所选盘大小决定）
	UDiskId *string `required:"false"`
}

// GetPHostDiskUpgradePriceResponse is response schema for GetPHostDiskUpgradePrice action
type GetPHostDiskUpgradePriceResponse struct {
	response.CommonBase

	// 升价差价原价。精度为小数点后2位。
	OriginalPrice float64

	// 升级差价。精度为小数点后2位。
	Price float64
}

// NewGetPHostDiskUpgradePriceRequest will create request of GetPHostDiskUpgradePrice action.
func (c *UPHostClient) NewGetPHostDiskUpgradePriceRequest() *GetPHostDiskUpgradePriceRequest {
	req := &GetPHostDiskUpgradePriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetPHostDiskUpgradePrice

获取物理云裸金属挂载磁盘的升级价格
*/
func (c *UPHostClient) GetPHostDiskUpgradePrice(req *GetPHostDiskUpgradePriceRequest) (*GetPHostDiskUpgradePriceResponse, error) {
	var err error
	var res GetPHostDiskUpgradePriceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetPHostDiskUpgradePrice", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

/*
GetPHostPriceParamDisks is request schema for complex param
*/
type GetPHostPriceParamDisks struct {

	// 裸金属机型参数->枚举值：\\ > True，是系统盘 \\ > False，是数据盘（默认）。Disks数组中有且只能有一块盘是系统盘。
	IsBoot *string `required:"false"`

	// 裸金属机型参数->磁盘大小，单位GB，必须是10GB的整数倍。系统盘20-500GB。数据盘是20-32000G。
	Size *string `required:"false"`

	// 裸金属机型参数->磁盘类型：枚举值：CLOUD_RSSD
	Type *string `required:"false"`
}

// GetPHostPriceRequest is request schema for GetPHostPrice action
type GetPHostPriceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"false"`

	// 计费模式，枚举值为： Year/Month
	ChargeType *string `required:"true"`

	// 网络环境，可选千兆：1G ；万兆：10G；25G网络：25G。
	Cluster *string `required:"false"`

	// 购买数量，范围[1-5]
	Count *int `required:"true"`

	//
	Disks []GetPHostPriceParamDisks `required:"false"`

	// 购买时长，1-10个月或1-10年；默认值为1。月付时，此参数传0，代表购买至月末，1代表整月。
	Quantity *int `required:"true"`

	// 默认为：DB(数据库型)，可以通过接口 [DescribePHostMachineType](api/uphost-api/describe_phost_machine_type.html)获取
	Type *string `required:"false"`
}

// GetPHostPriceResponse is response schema for GetPHostPrice action
type GetPHostPriceResponse struct {
	response.CommonBase

	// 价格列表 见 PHostPriceSet
	PriceSet []PHostPriceSet
}

// NewGetPHostPriceRequest will create request of GetPHostPrice action.
func (c *UPHostClient) NewGetPHostPriceRequest() *GetPHostPriceRequest {
	req := &GetPHostPriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetPHostPrice

获取物理机价格列表
*/
func (c *UPHostClient) GetPHostPrice(req *GetPHostPriceRequest) (*GetPHostPriceResponse, error) {
	var err error
	var res GetPHostPriceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetPHostPrice", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetPhostDiskUpgradePriceRequest is request schema for GetPhostDiskUpgradePrice action
type GetPhostDiskUpgradePriceRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 磁盘ID。
	DiskId *string `required:"true"`

	// 裸金属机型参数->磁盘大小，单位GB，必须是10GB的整数倍。系统盘20-500GB，数据盘单块盘20-32000GB。
	DiskSpace *int `required:"true"`

	// UPHost实例ID。
	PHostId *string `required:"true"`
}

// GetPhostDiskUpgradePriceResponse is response schema for GetPhostDiskUpgradePrice action
type GetPhostDiskUpgradePriceResponse struct {
	response.CommonBase

	// 升级差价。精度为小数点后2位。
	Price float64
}

// NewGetPhostDiskUpgradePriceRequest will create request of GetPhostDiskUpgradePrice action.
func (c *UPHostClient) NewGetPhostDiskUpgradePriceRequest() *GetPhostDiskUpgradePriceRequest {
	req := &GetPhostDiskUpgradePriceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetPhostDiskUpgradePrice

获取物理云裸金属挂载磁盘的升级价格
*/
func (c *UPHostClient) GetPhostDiskUpgradePrice(req *GetPhostDiskUpgradePriceRequest) (*GetPhostDiskUpgradePriceResponse, error) {
	var err error
	var res GetPhostDiskUpgradePriceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetPhostDiskUpgradePrice", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ModifyPHostImageInfoRequest is request schema for ModifyPHostImageInfo action
type ModifyPHostImageInfoRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"true"`

	// 镜像ID
	ImageId *string `required:"true"`

	// 镜像名称
	Name *string `required:"false"`

	// 备注
	Remark *string `required:"false"`
}

// ModifyPHostImageInfoResponse is response schema for ModifyPHostImageInfo action
type ModifyPHostImageInfoResponse struct {
	response.CommonBase

	// 镜像ID
	ImageId string
}

// NewModifyPHostImageInfoRequest will create request of ModifyPHostImageInfo action.
func (c *UPHostClient) NewModifyPHostImageInfoRequest() *ModifyPHostImageInfoRequest {
	req := &ModifyPHostImageInfoRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ModifyPHostImageInfo

修改自定义镜像名称和备注
*/
func (c *UPHostClient) ModifyPHostImageInfo(req *ModifyPHostImageInfoRequest) (*ModifyPHostImageInfoResponse, error) {
	var err error
	var res ModifyPHostImageInfoResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ModifyPHostImageInfo", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ModifyPHostInfoRequest is request schema for ModifyPHostInfo action
type ModifyPHostInfoRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// 物理机名称，默认不更改
	Name *string `required:"false"`

	// 物理机资源ID
	PHostId *string `required:"true"`

	// 物理机备注，默认不更改
	Remark *string `required:"false"`

	// 业务组，默认不更改
	Tag *string `required:"false"`
}

// ModifyPHostInfoResponse is response schema for ModifyPHostInfo action
type ModifyPHostInfoResponse struct {
	response.CommonBase

	// PHost 的资源ID
	PHostId string
}

// NewModifyPHostInfoRequest will create request of ModifyPHostInfo action.
func (c *UPHostClient) NewModifyPHostInfoRequest() *ModifyPHostInfoRequest {
	req := &ModifyPHostInfoRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ModifyPHostInfo

更改物理机信息
*/
func (c *UPHostClient) ModifyPHostInfo(req *ModifyPHostInfoRequest) (*ModifyPHostInfoResponse, error) {
	var err error
	var res ModifyPHostInfoResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ModifyPHostInfo", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// PoweroffPHostRequest is request schema for PoweroffPHost action
type PoweroffPHostRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// PHost资源ID
	PHostId *string `required:"true"`
}

// PoweroffPHostResponse is response schema for PoweroffPHost action
type PoweroffPHostResponse struct {
	response.CommonBase

	// PHost 的资源ID
	PHostId string
}

// NewPoweroffPHostRequest will create request of PoweroffPHost action.
func (c *UPHostClient) NewPoweroffPHostRequest() *PoweroffPHostRequest {
	req := &PoweroffPHostRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: PoweroffPHost

断电物理云主机
*/
func (c *UPHostClient) PoweroffPHost(req *PoweroffPHostRequest) (*PoweroffPHostResponse, error) {
	var err error
	var res PoweroffPHostResponse

	reqCopier := *req

	err = c.Client.InvokeAction("PoweroffPHost", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// RebootPHostRequest is request schema for RebootPHost action
type RebootPHostRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// PHost资源ID
	PHostId *string `required:"true"`
}

// RebootPHostResponse is response schema for RebootPHost action
type RebootPHostResponse struct {
	response.CommonBase

	// PHost 的资源ID
	PHostId string
}

// NewRebootPHostRequest will create request of RebootPHost action.
func (c *UPHostClient) NewRebootPHostRequest() *RebootPHostRequest {
	req := &RebootPHostRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: RebootPHost

重启物理机
*/
func (c *UPHostClient) RebootPHost(req *RebootPHostRequest) (*RebootPHostResponse, error) {
	var err error
	var res RebootPHostResponse

	reqCopier := *req

	err = c.Client.InvokeAction("RebootPHost", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ReinstallPHostRequest is request schema for ReinstallPHost action
type ReinstallPHostRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 裸金属机型参数->系统盘大小。 单位：GB， 范围[20,500]， 步长：10
	BootDiskSpace *int `required:"false"`

	// 镜像Id，参考镜像列表，默认使用原镜像
	ImageId *string `required:"false"`

	// 物理机名称，默认不更改
	Name *string `required:"false"`

	// PHost资源ID
	PHostId *string `required:"true"`

	// 密码
	Password *string `required:"true"`

	// 不保留数据盘重装，可选Raid
	Raid *string `required:"false"`

	// 物理机备注，默认为不更改。
	Remark *string `required:"false"`

	// 是否保留数据盘，保留：Yes，不报留：No， 默认：Yes
	ReserveDisk *string `required:"false"`

	// 业务组，默认不更改。
	Tag *string `required:"false"`
}

// ReinstallPHostResponse is response schema for ReinstallPHost action
type ReinstallPHostResponse struct {
	response.CommonBase

	// PHost 的资源ID
	PHostId string
}

// NewReinstallPHostRequest will create request of ReinstallPHost action.
func (c *UPHostClient) NewReinstallPHostRequest() *ReinstallPHostRequest {
	req := &ReinstallPHostRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ReinstallPHost

重装物理机操作系统
*/
func (c *UPHostClient) ReinstallPHost(req *ReinstallPHostRequest) (*ReinstallPHostResponse, error) {
	var err error
	var res ReinstallPHostResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ReinstallPHost", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ResetPHostPasswordRequest is request schema for ResetPHostPassword action
type ResetPHostPasswordRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 裸金属实例ID
	PHostId *string `required:"true"`

	// PHost新密码（密码格式使用BASE64编码）
	Password *string `required:"true"`
}

// ResetPHostPasswordResponse is response schema for ResetPHostPassword action
type ResetPHostPasswordResponse struct {
	response.CommonBase

	// 裸金属实例ID
	PHostId string
}

// NewResetPHostPasswordRequest will create request of ResetPHostPassword action.
func (c *UPHostClient) NewResetPHostPasswordRequest() *ResetPHostPasswordRequest {
	req := &ResetPHostPasswordRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ResetPHostPassword

重置裸金属实例的管理员密码
*/
func (c *UPHostClient) ResetPHostPassword(req *ResetPHostPasswordRequest) (*ResetPHostPasswordResponse, error) {
	var err error
	var res ResetPHostPasswordResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ResetPHostPassword", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ResizePHostAttachedDiskRequest is request schema for ResizePHostAttachedDisk action
type ResizePHostAttachedDiskRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"true"`

	// 裸金属机型参数->磁盘大小，单位GB，必须是10GB的整数倍。系统盘20-500GB，数据盘单块盘20-32000GB。
	DiskSpace *int `required:"false"`

	// UPHost实例ID。
	PHostId *string `required:"false"`

	// 磁盘ID。
	UDiskId *string `required:"false"`
}

// ResizePHostAttachedDiskResponse is response schema for ResizePHostAttachedDisk action
type ResizePHostAttachedDiskResponse struct {
	response.CommonBase

	// 改配成功的磁盘id
	UDiskId string
}

// NewResizePHostAttachedDiskRequest will create request of ResizePHostAttachedDisk action.
func (c *UPHostClient) NewResizePHostAttachedDiskRequest() *ResizePHostAttachedDiskRequest {
	req := &ResizePHostAttachedDiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ResizePHostAttachedDisk

修改裸金属物理云已经挂载的云盘容量大小
*/
func (c *UPHostClient) ResizePHostAttachedDisk(req *ResizePHostAttachedDiskRequest) (*ResizePHostAttachedDiskResponse, error) {
	var err error
	var res ResizePHostAttachedDiskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ResizePHostAttachedDisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// StartPHostRequest is request schema for StartPHost action
type StartPHostRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// PHost资源ID
	PHostId *string `required:"true"`
}

// StartPHostResponse is response schema for StartPHost action
type StartPHostResponse struct {
	response.CommonBase

	// PHost 的资源ID
	PHostId string
}

// NewStartPHostRequest will create request of StartPHost action.
func (c *UPHostClient) NewStartPHostRequest() *StartPHostRequest {
	req := &StartPHostRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: StartPHost

启动物理机
*/
func (c *UPHostClient) StartPHost(req *StartPHostRequest) (*StartPHostResponse, error) {
	var err error
	var res StartPHostResponse

	reqCopier := *req

	err = c.Client.InvokeAction("StartPHost", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// StopPHostRequest is request schema for StopPHost action
type StopPHostRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"false"`

	// PHost资源ID
	PHostId *string `required:"true"`
}

// StopPHostResponse is response schema for StopPHost action
type StopPHostResponse struct {
	response.CommonBase

	// PHost 的资源ID
	PHostId string
}

// NewStopPHostRequest will create request of StopPHost action.
func (c *UPHostClient) NewStopPHostRequest() *StopPHostRequest {
	req := &StopPHostRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: StopPHost

关闭物理机
*/
func (c *UPHostClient) StopPHost(req *StopPHostRequest) (*StopPHostResponse, error) {
	var err error
	var res StopPHostResponse

	reqCopier := *req

	err = c.Client.InvokeAction("StopPHost", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// TerminatePHostRequest is request schema for TerminatePHost action
type TerminatePHostRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// PHost资源ID
	PHostId *string `required:"true"`

	// 是否释放绑定的EIP。true: 解绑EIP后，并释放；其他值或不填：解绑EIP。
	ReleaseEIP *bool `required:"false"`

	// 裸金属机型参数->删除主机时是否同时删除挂载的数据盘。默认为false。
	ReleaseUDisk *bool `required:"false"`
}

// TerminatePHostResponse is response schema for TerminatePHost action
type TerminatePHostResponse struct {
	response.CommonBase

	// PHost 的资源ID
	PHostId string
}

// NewTerminatePHostRequest will create request of TerminatePHost action.
func (c *UPHostClient) NewTerminatePHostRequest() *TerminatePHostRequest {
	req := &TerminatePHostRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: TerminatePHost

删除物理云主机
*/
func (c *UPHostClient) TerminatePHost(req *TerminatePHostRequest) (*TerminatePHostResponse, error) {
	var err error
	var res TerminatePHostResponse

	reqCopier := *req

	err = c.Client.InvokeAction("TerminatePHost", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// TerminatePHostImageRequest is request schema for TerminatePHostImage action
type TerminatePHostImageRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"true"`

	// 自制镜像ID
	ImageId *string `required:"true"`
}

// TerminatePHostImageResponse is response schema for TerminatePHostImage action
type TerminatePHostImageResponse struct {
	response.CommonBase

	// 自制镜像ID
	ImageId string
}

// NewTerminatePHostImageRequest will create request of TerminatePHostImage action.
func (c *UPHostClient) NewTerminatePHostImageRequest() *TerminatePHostImageRequest {
	req := &TerminatePHostImageRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: TerminatePHostImage

删除裸金属2.0用户自定义镜像
*/
func (c *UPHostClient) TerminatePHostImage(req *TerminatePHostImageRequest) (*TerminatePHostImageResponse, error) {
	var err error
	var res TerminatePHostImageResponse

	reqCopier := *req

	err = c.Client.InvokeAction("TerminatePHostImage", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
