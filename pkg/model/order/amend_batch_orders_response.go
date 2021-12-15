package order

import "github.com/rwehresmann/bitcom-go-api/pkg/model/base"

type AmendBatchOrdersResponse struct {
	base.RestBaseResponse
	Data BatchOrderVo `json:"data"`
}
