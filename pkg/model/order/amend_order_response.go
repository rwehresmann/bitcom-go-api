package order

import "github.com/rwehresmann/bitcom-go-api/pkg/model/base"

type AmendOrderResponse struct {
	base.RestBaseResponse
	Data OrderActionVo `json:"data"`
}
