package order

import "github.com/rwehresmann/bitcom-go-api/pkg/model/base"

type ClosePositionsResponse struct {
	base.RestBaseResponse
	Data OrderActionVo `json:"data"`
}
