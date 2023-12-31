package todo

import (
	"encoding/json"
	"net/http"
)

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  interface{} `json:"error"`
	Data   interface{} `json:"data"`
}

func ReadRequest(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(result)
	PanicIfError(err)
}

func GenerateResponse(w http.ResponseWriter, res interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(res)
	PanicIfError(err)
}

func httpRes(httpstatus int, res interface{}, err interface{}) WebResponse {
	return WebResponse{
		Code:   httpstatus,
		Status: http.StatusText(httpstatus),
		Data:   res,
		Error:  err,
	}
}
