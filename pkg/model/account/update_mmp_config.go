package account

import "github.com/rwehresmann/bitcom-go-api/pkg/model/base"

type UpdateMmpResponse struct {
	base.RestBaseResponse
	Data string `json:"data"`
}

type ResetMmpStateResponse struct {
	base.RestBaseResponse
	Data string `json:"data"`
}
