package ws

import (
	"github.com/rwehresmann/bitcom-go-api/pkg/model/base"
	"github.com/rwehresmann/bitcom-go-api/pkg/model/order"
)

type UserTradeResponse struct {
	base.WsBaseResponse
	Data []TradeVo `json:"data"`
}

type TradeVo order.TradeVo
