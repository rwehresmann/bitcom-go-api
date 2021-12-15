package ws

import (
	"github.com/rwehresmann/bitcom-go-api/pkg/model/base"
	"github.com/rwehresmann/bitcom-go-api/pkg/model/market"
)

type TickerResponse struct {
	base.WsBaseResponse
	Data TickerMessage `json:"data"`
}

type TickerMessage market.TickerVo
