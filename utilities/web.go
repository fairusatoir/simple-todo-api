package utilities

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
	PanicOnError(err)
}

func httpRes(httpstatus int, res interface{}, err interface{}) WebResponse {
	return WebResponse{
		Code:   httpstatus,
		Status: http.StatusText(httpstatus),
		Data:   res,
		Error:  err,
	}
}

func GenerateResponse(w http.ResponseWriter, httpstatus int, data interface{}, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	e := encoder.Encode(httpRes(httpstatus, data, err))
	PanicOnError(e)
}
