package market

import "github.com/rwehresmann/bitcom-go-api/pkg/model/base"

type GetCurrenciesResponse struct {
	base.RestBaseResponse
	Data CurrenciesVo `json:"data"`
}

type CurrenciesVo struct {
	Currencies []string `json:"currencies"`
}
