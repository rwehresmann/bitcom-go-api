package ws

import (
	"github.com/rwehresmann/bitcom-go-api/pkg/model/account"
	"github.com/rwehresmann/bitcom-go-api/pkg/model/base"
)

type PositionResponse struct {
	base.WsBaseResponse
	Data []PositionVo `json:"data"`
}

type PositionVo account.PositionVo
