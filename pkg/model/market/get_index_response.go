package market

import "github.com/rwehresmann/bitcom-go-api/pkg/model/base"

type GetIndexResponse struct {
	base.RestBaseResponse
	Data IndexVo `json:"data"`
}

type IndexVo struct {
	Currency   string `json:"name" example:"BTC"`
	IndexPrice string `json:"index_price" example:"8000"`
}
