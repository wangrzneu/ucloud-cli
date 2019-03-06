//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB BackupUDBInstanceSlowLog

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// BackupUDBInstanceSlowLogRequest is request schema for BackupUDBInstanceSlowLog action
type BackupUDBInstanceSlowLogRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// DB实例Id,该值可以通过DescribeUDBInstance获取
	DBId *string `required:"true"`

	// 过滤条件:起始时间(时间戳)
	BeginTime *int `required:"true"`

	// 过滤条件:结束时间(时间戳)
	EndTime *int `required:"true"`

	// 备份文件名称
	BackupName *string `required:"true"`
}

// BackupUDBInstanceSlowLogResponse is response schema for BackupUDBInstanceSlowLog action
type BackupUDBInstanceSlowLogResponse struct {
	response.CommonBase
}

// NewBackupUDBInstanceSlowLogRequest will create request of BackupUDBInstanceSlowLog action.
func (c *UDBClient) NewBackupUDBInstanceSlowLogRequest() *BackupUDBInstanceSlowLogRequest {
	req := &BackupUDBInstanceSlowLogRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// BackupUDBInstanceSlowLog - 备份UDB指定时间段的slowlog分析结果
func (c *UDBClient) BackupUDBInstanceSlowLog(req *BackupUDBInstanceSlowLogRequest) (*BackupUDBInstanceSlowLogResponse, error) {
	var err error
	var res BackupUDBInstanceSlowLogResponse

	err = c.client.InvokeAction("BackupUDBInstanceSlowLog", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
