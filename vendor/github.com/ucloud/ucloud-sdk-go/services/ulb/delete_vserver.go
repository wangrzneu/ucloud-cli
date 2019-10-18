//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api ULB DeleteVServer

package ulb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteVServerRequest is request schema for DeleteVServer action
type DeleteVServerRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// 负载均衡实例的ID
	ULBId *string `required:"true"`

	// VServer实例的ID
	VServerId *string `required:"true"`
}

// DeleteVServerResponse is response schema for DeleteVServer action
type DeleteVServerResponse struct {
	response.CommonBase
}

// NewDeleteVServerRequest will create request of DeleteVServer action.
func (c *ULBClient) NewDeleteVServerRequest() *DeleteVServerRequest {
	req := &DeleteVServerRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteVServer - 删除VServer实例
func (c *ULBClient) DeleteVServer(req *DeleteVServerRequest) (*DeleteVServerResponse, error) {
	var err error
	var res DeleteVServerResponse

	err = c.Client.InvokeAction("DeleteVServer", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}