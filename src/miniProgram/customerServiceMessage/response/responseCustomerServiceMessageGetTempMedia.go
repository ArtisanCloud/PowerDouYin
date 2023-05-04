package response

import (
	"github.com/ArtisanCloud/PowerDouYin/src/kernel/response"
)

type ResponseCustomerServiceMessageGetTempMedia struct {
	response.ResponseMiniProgram
	ContentType string `json:"contentType"`
	Buffer      []byte `json:"buffer"`
}
