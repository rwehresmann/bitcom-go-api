package order

import "github.com/rwehresmann/bitcom-go-api/pkg/model/base"

type GetOpenOrdersResponse struct {
	base.RestBaseResponse
	Data []OrderVo `json:"data"`
}
