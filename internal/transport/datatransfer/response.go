package datatransfer

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  *Error      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type Error struct {
	Message string `json:"message"`
}

func Response(hs int, d interface{}) *BaseResponse {
	return &BaseResponse{
		Code:   hs,
		Status: http.StatusText(hs),
		Data:   d,
	}
}

func ErrorResponse(hs int, err error) *BaseResponse {
	return &BaseResponse{
		Code:   hs,
		Status: http.StatusText(hs),
		Error:  &Error{Message: err.Error()},
	}
}

func Write(w http.ResponseWriter, resp *BaseResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)

	return json.NewEncoder(w).Encode(resp)
}
