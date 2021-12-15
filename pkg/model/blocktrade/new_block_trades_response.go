package blocktrade

import "github.com/rwehresmann/bitcom-go-api/pkg/model/base"

type NewBlockTradeResponse struct {
	base.RestBaseResponse
	Data ParadigmBlockTradeVO `json:"data"`
}

type ParadigmBlockTradeVO struct {
	Label string `json:"label"`
	// BlockOrderStatus
	Status string `json:"status"`
}
