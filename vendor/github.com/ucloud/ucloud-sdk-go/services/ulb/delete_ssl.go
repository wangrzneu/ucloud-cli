//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api ULB DeleteSSL

package ulb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteSSLRequest is request schema for DeleteSSL action
type DeleteSSLRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// SSL证书的ID
	SSLId *string `required:"true"`
}

// DeleteSSLResponse is response schema for DeleteSSL action
type DeleteSSLResponse struct {
	response.CommonBase
}

// NewDeleteSSLRequest will create request of DeleteSSL action.
func (c *ULBClient) NewDeleteSSLRequest() *DeleteSSLRequest {
	req := &DeleteSSLRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteSSL - 删除SSL证书
func (c *ULBClient) DeleteSSL(req *DeleteSSLRequest) (*DeleteSSLResponse, error) {
	var err error
	var res DeleteSSLResponse

	err = c.client.InvokeAction("DeleteSSL", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
