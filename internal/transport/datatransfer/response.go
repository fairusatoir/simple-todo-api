package datatransfer

import (
	"encoding/json"
	"net/http"
	"simple-to-do/internal/config"
	"simple-to-do/internal/utils/constants"
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

func Response(c int, d interface{}) *BaseResponse {
	return &BaseResponse{
		Code:   c,
		Status: http.StatusText(c),
		Data:   d,
	}
}

func errProd(c int, err *error) {
	if config.IsAppProd() {
		switch c {
		case http.StatusInternalServerError:
			*err = constants.Err500Prod
		}
	}
}

func ErrorResponse(c int, err error) *BaseResponse {
	errProd(c, &err)
	return &BaseResponse{
		Code:   c,
		Status: http.StatusText(c),
		Error:  &Error{Message: err.Error()},
	}
}

func Write(w http.ResponseWriter, resp *BaseResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)

	return json.NewEncoder(w).Encode(resp)
}
