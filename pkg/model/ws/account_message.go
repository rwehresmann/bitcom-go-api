package ws

import (
	"github.com/rwehresmann/bitcom-go-api/pkg/model/account"
	"github.com/rwehresmann/bitcom-go-api/pkg/model/base"
)

type AccountResponse struct {
	base.WsBaseResponse
	Data AccountVo `json:"data"`
}

type AccountVo account.AccountVo
