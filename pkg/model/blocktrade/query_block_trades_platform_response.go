package blocktrade

import "github.com/rwehresmann/bitcom-go-api/pkg/model/base"

type QueryBlockTradeByPlatformResponse struct {
	base.RestBaseResponse
	Data []ParadigmBlockTradeQueryVo `json:"data"`
}
