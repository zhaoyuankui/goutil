// httputil.go implements utilities for http.
// @Author zhaoyuankui@p1.com
// @Date 2019/10/09
package httputil

import (
	"encoding/json"
	"net/http"
)

const (
	respTypeJson = "application/json"
)

// RespondOK responds the obj in json and http.StatusOK to client.
func RespondOK(w http.ResponseWriter, obj interface{}) error {
	return RespondJson(w, obj, http.StatusOK)
}

// RespondJson responds the obj in json and the specified responseCode to client.
func RespondJson(w http.ResponseWriter, obj interface{}, responseCode int) error {
	data, err := json.Marshal(obj)
	if nil != err {
		return err
	}
	respond(w, respTypeJson, data, responseCode)
	return nil
}

func respond(w http.ResponseWriter, t string, data []byte, responseCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	w.Write(data)
}
