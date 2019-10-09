// http.go implements test utilities for http.
// @Author zhaoyuankui@p1.com
// @Date 2019/10/09
package testutil

import (
	"io/ioutil"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func Get(router *gin.Engine, uri string, headers, params map[string]string) ([]byte, int) {
	req := httptest.NewRequest("GET", uri, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	query := req.URL.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	req.URL.RawQuery = query.Encode()
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
	if nil != result {
		defer result.Body.Close()
	}
	body, _ := ioutil.ReadAll(result.Body)
	return body, result.StatusCode
}
