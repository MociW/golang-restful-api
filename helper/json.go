package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromReq(req *http.Request, result interface{}) {
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&result)
	PanicIfError(err)
}

func WriteToResponeBody(w http.ResponseWriter, respone interface{}) {
	w.Header().Add("Content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(respone)
	PanicIfError(err)
}
