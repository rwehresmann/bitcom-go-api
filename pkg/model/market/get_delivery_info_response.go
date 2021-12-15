package market

import "github.com/rwehresmann/bitcom-go-api/pkg/model/base"

type GetDeliveryInfoResponse struct {
	base.RestBaseResponse
	Data []DeliveryPriceVo `json:"data"`
}

type DeliveryPriceVo struct {
	DeliveryTime int64  `json:"delivery_time"`
	Currency     string `json:"currency"`
	Price        string `json:"price"`
}
