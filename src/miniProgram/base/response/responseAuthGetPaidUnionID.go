package response

import (
	"github.com/ArtisanCloud/PowerDouYin/src/kernel/response"
)

type ResponseAuthGetPaidUnionID struct {
	UnionID string `json:"unionid"`

	response.ResponseMiniProgram
}
