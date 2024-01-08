package datatransfer

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  error       `json:"error"`
	Data   interface{} `json:"data"`
}

func writeBody(hs int, d interface{}, err error) *BaseResponse {
	return &BaseResponse{
		Code:   hs,
		Status: http.StatusText(hs),
		Data:   d,
		Error:  err,
	}
}

func Write(w http.ResponseWriter, hs int, d interface{}, e error) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(hs)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(writeBody(hs, d, e))
	if err != nil {
		return err
	}
	return nil
}
