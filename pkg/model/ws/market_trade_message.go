package ws

import "github.com/rwehresmann/bitcom-go-api/pkg/model/base"

type MarketTradeResponse struct {
	base.WsBaseResponse
	Data []TradeMessage `json:"data"`
}
