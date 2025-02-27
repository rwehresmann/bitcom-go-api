package ws

import (
	"github.com/rwehresmann/bitcom-go-api/pkg/model/base"
	"github.com/rwehresmann/bitcom-go-api/pkg/model/order"
)

type OrderResponse struct {
	base.WsBaseResponse
	Data []OrderVo `json:"data"`
}

type OrderVo order.OrderActionVo
