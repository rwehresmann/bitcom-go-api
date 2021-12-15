package ws

import "github.com/rwehresmann/bitcom-go-api/pkg/model/base"

type GetWsAuthTokenResponse struct {
	base.RestBaseResponse
	Data WsToken `json:"data"`
}

type WsToken struct {
	Token string `json:"token"`
}
