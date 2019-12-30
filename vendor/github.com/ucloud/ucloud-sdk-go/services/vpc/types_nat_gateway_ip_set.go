// Code is generated by ucloud-model, DO NOT EDIT IT.

package vpc

/*
NatGatewayIPSet - IPSet信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type NatGatewayIPSet struct {

	// 带宽
	Bandwidth int

	// EIP带宽类型
	BandwidthType string

	// 外网IP的 EIPId
	EIPId string

	// 外网IP信息
	IPResInfo []NatGWIPResInfo

	// 权重为100的为出口
	Weight int
}
