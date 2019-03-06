//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api VPC DescribeVPCIntercom

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeVPCIntercomRequest is request schema for DescribeVPCIntercom action
type DescribeVPCIntercomRequest struct {
	request.CommonBase

	// [公共参数] 源VPC所在地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 源VPC所在项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// VPC短ID
	VPCId *string `required:"true"`

	// 目的VPC所在地域，默认为全部地域
	DstRegion *string `required:"false"`

	// 目的项目ID，默认为全部项目
	DstProjectId *string `required:"false"`
}

// DescribeVPCIntercomResponse is response schema for DescribeVPCIntercom action
type DescribeVPCIntercomResponse struct {
	response.CommonBase

	// 联通VPC信息数组
	DataSet []VPCIntercomInfo
}

// NewDescribeVPCIntercomRequest will create request of DescribeVPCIntercom action.
func (c *VPCClient) NewDescribeVPCIntercomRequest() *DescribeVPCIntercomRequest {
	req := &DescribeVPCIntercomRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeVPCIntercom - 获取VPC互通信息
func (c *VPCClient) DescribeVPCIntercom(req *DescribeVPCIntercomRequest) (*DescribeVPCIntercomResponse, error) {
	var err error
	var res DescribeVPCIntercomResponse

	err = c.client.InvokeAction("DescribeVPCIntercom", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
