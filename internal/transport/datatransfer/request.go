package datatransfer

import (
	"encoding/json"
	"net/http"
)

func Bind(r *http.Request, result interface{}) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(result)
	if err != nil {
		return err
	}

	return nil
}
