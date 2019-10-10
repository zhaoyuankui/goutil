// http.go implements test utilities for http.
// @Author zhaoyuankui@p1.com
// @Date 2019/10/09
package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// Get launch a http GET request.
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
	return doRequest(router, req)
}

// CustomizeGet launch a http GET request with a callback customizer to customize the request.
func CustomizeGet(router *gin.Engine, uri string, customizer func(*http.Request)) ([]byte, int) {
	req := httptest.NewRequest("GET", uri, nil)
	customizer(req)
	return doRequest(router, req)
}

func doRequest(router *gin.Engine, req *http.Request) ([]byte, int) {
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
	if nil != result {
		defer result.Body.Close()
	}
	body, _ := ioutil.ReadAll(result.Body)
	return body, result.StatusCode
}
